package db

import (
	"strings"
	"web-service/pkg/server"
)

// Initialize Function in DB
func init() {
	// Database Configuration Value
	switch strings.ToLower(server.Config.GetString("DB_DRIVER")) {
	case "postgres":
		server.Config.SetDefault("DB_PORT", "5432")

		psqlCfg.Host = server.Config.GetString("DB_HOST")
		psqlCfg.Port = server.Config.GetString("DB_PORT")
		psqlCfg.User = server.Config.GetString("DB_USER")
		psqlCfg.Password = server.Config.GetString("DB_PASSWORD")
		psqlCfg.Name = server.Config.GetString("DB_NAME")

		if len(psqlCfg.Host) != 0 && len(psqlCfg.Port) != 0 &&
			len(psqlCfg.User) != 0 && len(psqlCfg.Password) != 0 &&
			len(psqlCfg.Name) != 0 {

			// Do MySQL Database Connection
			PSQL = psqlConnect()
		}
	case "mysql":
		server.Config.SetDefault("DB_PORT", "3306")

		mysqlCfg.Host = server.Config.GetString("DB_HOST")
		mysqlCfg.Port = server.Config.GetString("DB_PORT")
		mysqlCfg.User = server.Config.GetString("DB_USER")
		mysqlCfg.Password = server.Config.GetString("DB_PASSWORD")
		mysqlCfg.Name = server.Config.GetString("DB_NAME")

		if len(mysqlCfg.Host) != 0 && len(mysqlCfg.Port) != 0 &&
			len(mysqlCfg.User) != 0 && len(mysqlCfg.Password) != 0 &&
			len(mysqlCfg.Name) != 0 {

			// Do MySQL Database Connection
			MySQL = mysqlConnect()
		}
	}
}
