package dboptions

import "fmt"

type DatabaseOptions struct {
	DBType   string
	DBName   string
	Username string
	Password string
	Host     string
	Port     int
	Sslmode  bool
}

func (opt DatabaseOptions) GenerateConnectionString() string {
	sslmode := "disable"
	if opt.Sslmode {
		sslmode = "enable"
	}

	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		opt.DBType,
		opt.Username, opt.Password,
		opt.Host, opt.Port,
		opt.DBName,
		sslmode,
	)
}
