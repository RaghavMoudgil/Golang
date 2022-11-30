package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type dataBaseConfig struct {
	Host     string `yaml:"HOST"`
	DbName   string `yaml:"DBNAME"`
	User     string `yaml:"USER"`
	Password string `yaml:"PASSWORD"`
	Port     int    `yaml:"PORT"`
}

var Credential *config

type config struct {
	DataBaseConfig dataBaseConfig `yaml:"DATABASE"`
}

func NewConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&Credential); err != nil {
		return err
	}

	log.Printf("DBNAME: %s", Credential.DataBaseConfig.DbName)
	return nil
}

func InitDB() *sql.DB {
	setup := Credential.DataBaseConfig
	DataBase, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", setup.User, setup.Password, setup.Host, setup.DbName))
	if err != nil {
		log.Print("Error while opening the DB", err)
	} else {
		log.Print("Connected Successfully")
	}

	return DataBase
}
func InitDBstud() *sql.DB {
	setup := Credential.DataBaseConfig
	DataBase, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", setup.User, setup.Password, setup.Host, setup.DbName))
	if err != nil {
		log.Print("Error while opening the studentdataDB", err)
	} else {
		log.Print("Connected Successfully")
	}

	return DataBase
}
