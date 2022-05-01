package db

import (
	"database/sql"
	"origin-auth/getconf"

	"github.com/udonetsm/help/helper"
	"github.com/udonetsm/help/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func sqlDb() *sql.DB {
	db, err := sql.Open("pgx", getconf.Storeconf)
	helper.Errors(err, "sqlopen")
	db.SetConnMaxIdleTime(5)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(6)
	db.SetConnMaxLifetime(6)
	return db
}

func gormDb(sqlDb *sql.DB) *gorm.DB {
	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDb}), &gorm.Config{})
	helper.Errors(err, "gormopen")
	return gormDb
}

func Authentificate(auth models.Auth) (ok bool, user models.User) {
	sdb := sqlDb()
	defer sdb.Close()
	gdb := gormDb(sdb)
	var pass string
	gdb.Table("auth").Select("password").Where("email = ?", auth.Email).Scan(&pass)
	if pass != "" {
		passok := bcrypt.CompareHashAndPassword([]byte(pass), []byte(auth.Password)) == nil
		if passok {
			gdb.Table("users").Select("*").Where("email=?", auth.Email).Scan(&user)
			ok = true
			return
		}
	}
	ok = false
	return
}
