package main

import (
	"github.com/jinzhu/gorm"
	"encoding/base64"
	"time"
	"math/rand"
	"github.com/spf13/viper"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type dbConnection struct {
	Host     string
	Database string
	Port     string
	User     string
	Pass     string
	SSL      string
	Type     string
}

func Connect() (dbConn *gorm.DB, err1 error) {
	db := &dbConnection{}
	err := viper.UnmarshalKey("db", db)

	switch db.Type {
	case "postgres":
		logrus.Info("Connecting to postgres...")
		dbPass, err := base64.StdEncoding.DecodeString(db.Pass)
		if err != nil {
			return nil, err
		}
		tryConnection(6, time.Second, func() error {
			dbConn, err = gorm.Open(
				"postgres",
				"host="+db.Host+" port="+db.Port+" password="+string(dbPass)+
					" user="+db.User+" dbname="+db.Database+" sslmode="+db.SSL,
			)
			if err != nil {
				return err
			}
			dbConn.LogMode(true)
			return nil
		})

		break
	case "sqlite":
	default:
		logrus.Info("Connecting to in-memory database...")
		err = tryConnection(6, time.Second, func() error {
			dbConn, err = gorm.Open(
				"sqlite3", ":memory:")
			if nil != err {
				logrus.Error("connect error ", err)
				return err
			}
			dbConn.LogMode(true)
			return nil

		})

		break
	}

	if err != nil {
		logrus.WithError(err)
	}

	return dbConn, err1
}

func tryConnection(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		logrus.Info("Attempting to connect")
		if s, ok := err.(stop); ok {
			// return the original error
			return s.error
		}
		if attempts--; attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2
			time.Sleep(sleep)
			return tryConnection(attempts, 2*sleep, fn)
		}
		return err
	}
	logrus.Info("Connection successful!")
	return nil
}

type stop struct {
	error
}