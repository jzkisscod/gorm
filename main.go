package main

import (
	"github.com/sirupsen/logrus"
)

func main() {

	// db connection
	db, err := Connect()

	if nil != err {
		logrus.Error("Error on connection")
	}

	db.Debug().DropTableIfExists(&Dependent{}, &User{})		// if logMode is not set as true. Can use Debug to trace specific statement

	TestOneToMany(db)
}
