package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PgDsn string = "host=localhost user=lwinmgmg password=frontiir dbname=chat port=5432 sslmode=disable TimeZone=Asia/Rangoon"
)

var (
	PgDb *gorm.DB = nil
)

func init() {
	var err error
	if PgDb == nil {
		PgDb, err = GetPgConn(PgDsn)
		if err != nil {
			panic(err)
		}
	}
}

func GetPgConn(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
}
