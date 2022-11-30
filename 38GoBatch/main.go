package main

import (
	"batch/config"
	"batch/jobs"
	"context"
	"log"
	"time"

	"github.com/chararch/gobatch"
	"github.com/chararch/gobatch/util"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//connnecting to database diarectly

	// var db *sql.DB
	// var err error
	// db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	// if err != nil {
	// 	panic(err)
	// }

	// connecting gobatch database via config file

	err := config.NewConfig("./config/config.yaml")
	if err != nil {
		log.Printf("Failed to read config %s", err.Error())
	}
	sqlDb := config.InitDB()
	gobatch.SetDB(sqlDb)

	// //connecting studentdata databaase
	// sqlDbstud := config.InitDBstud()
	// gobatch.SetDB(sqlDbstud)

	param, _ := util.JsonString(map[string]interface{}{
		"rand": time.Now().Nanosecond(),
	})
	gobatch.Start(context.Background(), jobs.Job(), param)

	// //params for studentdata
	// params, _ := util.JsonString(map[string]interface{}{
	// 	"date": time.Now().Format("2006-01-02"),
	// 	"rand": time.Now().Nanosecond(),
	// })
	// gobatch.Start(context.Background(), jobs.RunStudJob(), params)

}
