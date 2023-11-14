package services

import (
	"fmt"

	"github.com/lwinmgmg/chat/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Env          = env.GetEnv()
	PgDsn string = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=UTC",
		Env.Settings.Postgres.Host,
		Env.Settings.Postgres.Port,
		Env.Settings.Postgres.Login,
		Env.Settings.Postgres.Password,
		Env.Settings.Postgres.DB,
	)
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
