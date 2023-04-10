package db

import (
	"platform/cmd/order/config"
	"platform/pkg"
	"platform/pkg/mysqld"
	"platform/pkg/postgres"
)

func GetDb(ds config.DataSource) (pkg.DB, error) {
	switch ds.Type {
	case "mysql":
		pg, err := mysqld.NewMysqlDb(ds.Mysql)
		return pg, err
	case "postgres":
		pg, err := postgres.NewPostgresDB(ds.PG.DsnURL)
		return pg, err
	}
	return nil, nil
}
