package di

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	dboptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/db_options"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	serveroptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/server/server_options"
)

// type RouteHandler func

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
	server   *echo.Echo
}

type Container struct {
	services map[identificators.Identificator]any
	server   *echo.Echo
}

func New() *ContainerBuilder {
	return &ContainerBuilder{
		services: make(map[identificators.Identificator]any),
		server:   echo.New(),
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

func (b *ContainerBuilder) AddRoute(verb HTTPVerb, path string, handler echo.HandlerFunc) *ContainerBuilder {
	switch verb {
	case GET:
		b.server.GET(path, handler)
	case POST:
		b.server.POST(path, handler)
	case PUT:
		b.server.PUT(path, handler)
	case PATCH:
		b.server.PATCH(path, handler)
	case DELETE:
		b.server.DELETE(path, handler)
	}
	return b
}

func (b *ContainerBuilder) Build() *Container {
	return &Container{
		services: b.services,
		server:   b.server,
	}
}

func GetService[T any](c *Container, identificator identificators.Identificator) (*T, error) {
	if _, ok := c.services[identificator]; !ok {
		return nil, fmt.Errorf("service with identificator '%s' is not found", string(identificator))
	}

	service := c.services[identificator]
	switch service.(type) {
	case T:
		res := service.(T)
		return &res, nil
	default:
		return nil, fmt.Errorf("service with identificator '%s' is with not type", string(identificator))
	}

}

func (c *Container) InitService(init func(c *Container)) {
	init(c)
}

func (c *Container) RunServer(serverOptions serveroptions.ServerOtions) {
	c.server.Start(fmt.Sprintf(":%d", serverOptions.Port))
}
