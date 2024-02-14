package upload

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"time"

	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/upload/upload_client"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/upload/upload_server"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const UploadServerIdentificator identificators.Identificator = "UploadServer"

func mapPresighedURLRequestToUploadFile(inp *upload_server.PresighedURLRequest) *dbmodels.Uploadfile {
	uploadFile := &dbmodels.Uploadfile{}
	uploadFile.FileID = null.StringFrom(inp.FileId)
	uploadFile.UserID = null.StringFrom(inp.UserId)
	uploadFile.Allowedfileformats = inp.AllowedFormats
	uploadFile.Maxsize = null.IntFrom(int(inp.MaxSize))
	uploadFile.Callbackadrr = null.StringFrom(inp.CallbackAdrr)
	return uploadFile
}

type UploadServerServer struct {
	minioClient *minio.Client
	db          *sql.DB
	uploadMap   map[string]chan bool
	upload_server.UnimplementedUploadServerServer
}

func (server *UploadServerServer) GeneratePresignedURL(
	ctx context.Context, in *upload_server.PresighedURLRequest) (*upload_server.PresighedURLResponse, error) {
	log.Println("hit")
	expatationDate, err := time.Parse(time.RFC3339, in.IsoTimeOfExpiration)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	expiration := time.Until(expatationDate)
	url, err := server.minioClient.PresignedPutObject(ctx, "bucket", in.FileId, expiration)
	if err != nil {
		return nil, err
	}
	uploadFile := mapPresighedURLRequestToUploadFile(in)
	uploadFile.DateOfExpration = null.TimeFrom(expatationDate)
	uploadFile.Presignedurl = null.StringFrom(url.String())

	err = uploadFile.Insert(ctx, server.db, boil.Blacklist(dbmodels.UploadfileColumns.ID))

	ch := make(chan bool)
	server.uploadMap[uploadFile.ID] = ch
	go func() {
		ctx := context.Background()
		defer close(ch)
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		conn, err := grpc.Dial(in.CallbackAdrr, opts...)
		if err != nil {
			log.Fatalf("Error while creating grpc dial: %s\n", err.Error())
		}
		defer conn.Close()
		client := upload_client.NewUploadClientClient(conn)
		select {
		case succ := <-ch:
			if !succ {
				client.OnErrorWhenUploading(ctx, &upload_client.CallbackRequest{
					FileId: in.FileId,
				})
			} else {
				client.OnSuccessfulUpload(ctx, &upload_client.CallbackSuccRequest{
					FileId: in.FileId,
					Url:    fmt.Sprintf("bucket/%s", in.FileId),
				})

			}
		case <-time.After(expiration):
			_, err := client.OnUrlExpiration(ctx, &upload_client.CallbackRequest{
				FileId: in.FileId,
			})
			log.Println(err)
		}
		delete(server.uploadMap, uploadFile.ID)
		uploadFile.Delete(ctx, server.db)
	}()

	if err != nil {
		return nil, err
	}

	return &upload_server.PresighedURLResponse{
		Url: uploadFile.ID,
	}, nil
}

type MinioOptions struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
}

func New(options MinioOptions) *UploadServerServer {
	minioClient, err := minio.New(options.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(options.AccessKeyId, options.SecretAccessKey, ""),
	})

	if err != nil {
		log.Fatalln(err)
	}

	return &UploadServerServer{
		minioClient: minioClient,
		uploadMap:   make(map[string]chan bool),
	}
}

type GrpcServerOptions struct{ Port string }

func (server *UploadServerServer) Init(db *sql.DB, options GrpcServerOptions) {
	server.db = db
	go func() {
		listener, err := net.Listen("tcp", options.Port)
		if err != nil {
			log.Fatalf("Error while creating tcp listener for the upload server: %s\n", err.Error())
		}
		s := grpc.NewServer()
		reflection.Register(s)

		upload_server.RegisterUploadServerServer(s, server)

		if err := s.Serve(listener); err != nil {
			log.Fatalf("Error while serving tcp listener for the upload server: %s\n", err.Error())
		}
	}()
}

func getTheMimeType(fileExt string) string {
	switch fileExt {
	case "jpg", "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "gif":
		return "image/gif"
	case "bmp":
		return "image/bmp"
	case "svg":
		return "image/svg+xml"
	case "webp":
		return "image/webp"
	case "ico":
		return "image/x-icon"
	default:
		return "unknown"
	}
}

func (server *UploadServerServer) Upload(id string, fileExt string, fileSize int64, src multipart.File) error {
	ctx := context.Background()
	record, err := dbmodels.FindUploadfile(ctx, server.db, id)
	if err != nil {
		return err
	}
	mime := getTheMimeType(fileExt)
	log.Println(fileExt)

	req, err := http.NewRequest(http.MethodPut, record.Presignedurl.String, src)
	if err != nil {
		log.Fatalln(err)
	}
	req.ContentLength = fileSize
	req.Header.Set("Content-Type", mime)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	server.uploadMap[id] <- resp.StatusCode == http.StatusOK
	return nil
}

func (server *UploadServerServer) RemoveUrl(id string) {
	server.uploadMap[id] <- false
}
