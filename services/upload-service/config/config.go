package config

import (
	"upload-service/upload"

	authoptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/authentication/auth_options"
	dboptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/db_options"
	serveroptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/server/server_options"
)

type ServiceConfig struct {
	ServerOptions     serveroptions.ServerOtions
	DatabaseOptions   dboptions.DatabaseOptions
	AuthOptions       authoptions.AuthOptions
	MinioOptions      upload.MinioOptions
	GrpcServerOptions upload.GrpcServerOptions
}

//in the future will be loadEnv function in the core lib

var Config ServiceConfig = ServiceConfig{
	ServerOptions: serveroptions.ServerOtions{
		Port: 5002,
	},
	DatabaseOptions: dboptions.DatabaseOptions{
		DBType:   "postgres",
		DBName:   "library",
		Username: "postgres",
		Password: "admin",
		Host:     "localhost",
		Port:     5432,
		Sslmode:  false,
	},
	AuthOptions: authoptions.AuthOptions{
		RealmInfoUrl: "http://localhost:8080/realms/master",
	},
	MinioOptions: upload.MinioOptions{
		Endpoint:        "127.0.0.1:9000",
		AccessKeyId:     "JEuIi3A7rEz1Nj9X0jMx",
		SecretAccessKey: "JiJyHV2F7gJFoN9Etwau6YQNvp4843uipjNReJJq",
	},
	GrpcServerOptions: upload.GrpcServerOptions{
		Port: ":5001",
	},
}
