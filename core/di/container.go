package di

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/authentication"
	authoptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/authentication/auth_options"
	dboptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/db_options"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/middleware"
	serveroptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/server/server_options"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/upload"
	_ "github.com/lib/pq"
)

type RouteHandler func(c *Container) echo.HandlerFunc

type HTTPVerb int

const (
	GET HTTPVerb = iota
	POST
	PUT
	PATCH
	DELETE
)

type ContainerBuilder struct {
	services map[identificators.Identificator]any
	// server   *echo.Echo
}

type Container struct {
	services map[identificators.Identificator]any
	server   *echo.Echo
}

func New() *ContainerBuilder {
	return &ContainerBuilder{
		services: make(map[identificators.Identificator]any),
	}
}

func (b *ContainerBuilder) RegisterService(identificator identificators.Identificator, service any) *ContainerBuilder {
	if _, ok := b.services[identificator]; ok {
		log.Fatalf("Service with identificator %s already exist!\n", identificator)
	}

	b.services[identificator] = service

	return b
}

func (b *ContainerBuilder) AddDatabase(opt dboptions.DatabaseOptions) *ContainerBuilder {
	db, err := sql.Open(opt.DBType, opt.GenerateConnectionString())
	if err != nil {
		log.Fatalf("Error when connecting to database: %e", err)
	}

	return b.RegisterService(identificators.Database, db)
}

func (b *ContainerBuilder) AddUploadService() *ContainerBuilder {
	uploadService := upload.NewUploadService()

	return b.RegisterService(identificators.UploadService, uploadService)
}

func (b *ContainerBuilder) Build() *Container {
	return &Container{
		services: b.services,
		server:   echo.New(),
	}
}

func GetService[T any](c *Container, identificator identificators.Identificator) (T, error) {

	var zeroValue T
	if _, ok := c.services[identificator]; !ok {
		return zeroValue, fmt.Errorf("service with identificator '%s' is not found", string(identificator))
	}

	service := c.services[identificator]
	switch res := service.(type) {
	case T:
		return res, nil
	default:
		return zeroValue, fmt.Errorf("service with identificator '%s' is with not type", string(identificator))
	}

}

func (c *Container) UseJWTHandlerMiddleware(opt authoptions.AuthOptions) *Container {
	rsaKey, err := authentication.LoadPublicKey(opt)

	if err != nil {
		log.Fatalf("Error: %e", err)
	}

	if err != nil {
		log.Fatalf("Error: rsaPublicKey - %s", err.Error())
	}

	c.server.Use(middleware.CreateJwtHandler(rsaKey))
	return c
}

func (c *Container) MapRoute(verb HTTPVerb, path string, handler RouteHandler) *Container {
	switch verb {
	case GET:
		c.server.GET(path, handler(c))
	case POST:
		c.server.POST(path, handler(c))
	case PUT:
		c.server.PUT(path, handler(c))
	case PATCH:
		c.server.PATCH(path, handler(c))
	case DELETE:
		c.server.DELETE(path, handler(c))
	}
	return c
}

func (c *Container) InitService(init func(c *Container)) *Container {
	init(c)
	return c
}

func (c *Container) RunServer(serverOptions serveroptions.ServerOtions) {
	c.server.Start(fmt.Sprintf(":%d", serverOptions.Port))
}
