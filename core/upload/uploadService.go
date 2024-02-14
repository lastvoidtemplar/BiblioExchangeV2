package upload

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/upload/upload_client"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/upload/upload_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type UploadFileTemplate struct {
	fileId             string
	allowedFileFormats []string
	maxSize            uint64
	userId             string
	urlExpirationTime  *time.Time
	onSuccess          func(fileId string, url string)
	onError            func(fileId string)
	onUrlExpiration    func(fileId string)
}

func NewUploadFileTemplate() UploadFileTemplate {
	return UploadFileTemplate{}
}

func (temp UploadFileTemplate) SetAllowedFileFormats(allowedFileFormats []string) UploadFileTemplate {
	temp.allowedFileFormats = allowedFileFormats
	return temp
}

func (temp UploadFileTemplate) SetMaxSize(maxSize uint64) UploadFileTemplate {
	temp.maxSize = maxSize
	return temp
}

func (temp UploadFileTemplate) OnSuccessCallback(callback func(fileId string, url string)) UploadFileTemplate {
	temp.onSuccess = callback
	return temp
}

func (temp UploadFileTemplate) OnErrorCallback(callback func(fileId string)) UploadFileTemplate {
	temp.onError = callback
	return temp
}

func (temp UploadFileTemplate) OnUrlExpirationCallback(callback func(fileId string)) UploadFileTemplate {
	temp.onUrlExpiration = callback
	return temp
}

func (temp UploadFileTemplate) SetPermission(userId string) UploadFileTemplate {
	temp.userId = userId
	return temp
}

func (temp UploadFileTemplate) SetPresignedUrlExpirationTime(expTime *time.Time) UploadFileTemplate {
	temp.urlExpirationTime = expTime
	return temp
}

func (temp UploadFileTemplate) SetIdentificator(fileId string) UploadFileTemplate {
	temp.fileId = fileId
	return temp
}

// Grpc Client server
type UploadClientServer struct {
	uploadingFiles map[string]UploadFileTemplate
	upload_client.UnimplementedUploadClientServer
}

func NewUploadClientServer() *UploadClientServer {
	return &UploadClientServer{
		uploadingFiles: make(map[string]UploadFileTemplate),
	}
}

func (clientServer *UploadClientServer) OnSuccessfulUpload(
	ctx context.Context, in *upload_client.CallbackSuccRequest) (
	*empty.Empty, error) {
	clientServer.uploadingFiles[in.FileId].onSuccess(in.FileId, in.Url)
	return &empty.Empty{}, nil
}
func (clientServer *UploadClientServer) OnErrorWhenUploading(
	ctx context.Context, in *upload_client.CallbackRequest) (
	*empty.Empty, error) {
	clientServer.uploadingFiles[in.FileId].onError(in.FileId)
	return &empty.Empty{}, nil
}
func (clientServer *UploadClientServer) OnUrlExpiration(
	ctx context.Context, in *upload_client.CallbackRequest) (
	*empty.Empty, error) {
	clientServer.uploadingFiles[in.FileId].onUrlExpiration(in.FileId)
	return &empty.Empty{}, nil
}

func (clientServer *UploadClientServer) AddFileUploadTemplate(temp UploadFileTemplate) {
	clientServer.uploadingFiles[temp.fileId] = temp
}

type UploadService struct {
	clientServer        *UploadClientServer
	grpcServer          *grpc.Server
	uploadClientService upload_server.UploadServerClient
	serviceAdrr         string
}

func NewUploadService() *UploadService {
	s := grpc.NewServer()
	reflection.Register(s)

	clientServer := NewUploadClientServer()

	upload_client.RegisterUploadClientServer(s, clientServer)

	return &UploadService{
		clientServer: clientServer,
		grpcServer:   s,
	}
}

func (s *UploadService) Init(options UploadServiceOptions) {
	go func() {
		s.serviceAdrr = options.ThisServiceAdrr
		listener, err := net.Listen("tcp", options.ThisServiceAdrr)
		if err != nil {
			log.Fatalf("Error while creating tcp listener for the upload service: %s\n", err.Error())
		}

		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}
		conn, err := grpc.Dial(options.UploadServiceAdrr, opts...)
		if err != nil {
			log.Fatalf("Error while creating grpc dial with the upload service: %s\n", err.Error())
		}

		s.uploadClientService = upload_server.NewUploadServerClient(conn)

		if err := s.grpcServer.Serve(listener); err != nil {
			log.Fatalf("Error while serving tcp listener for the upload service: %s\n", err.Error())
		}
	}()
}

func (s *UploadService) GetFilePresignedUrl(ctx context.Context, temp UploadFileTemplate) (string, error) {
	inp := &upload_server.PresighedURLRequest{
		FileId:              temp.fileId,
		AllowedFormats:      temp.allowedFileFormats,
		MaxSize:             temp.maxSize,
		IsoTimeOfExpiration: temp.urlExpirationTime.Format(time.RFC3339),
		UserId:              temp.userId,
		CallbackAdrr:        s.serviceAdrr,
	}
	log.Println(inp.IsoTimeOfExpiration)
	url, err := s.uploadClientService.GeneratePresignedURL(ctx, inp)

	if err != nil {
		return "", err
	}

	s.clientServer.AddFileUploadTemplate(temp)

	return url.Url, nil
}

type UploadServiceOptions struct {
	ThisServiceAdrr   string
	UploadServiceAdrr string
}
