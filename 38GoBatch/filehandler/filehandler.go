package filehandler

import (
	"batch/model"
	"database/sql"
	"github.com/chararch/gobatch"
	"github.com/chararch/gobatch/file"
	"time"
)

// struct for file variabels
var ReadFile = file.FileObjectModel{
	FileStore:     &file.LocalFileSystem{},
	FileName:      "res/studentdata.csv",
	Type:          file.CSV,
	Encoding:      "utf-8",
	Header:        false,
	ItemPrototype: &model.StudentData{},
}
var CreateFile = file.FileObjectModel{
	FileStore:     &file.LocalFileSystem{},
	FileName:      "res/{date,yyyyMMdd}/Cstudentdata.csv",
	Type:          file.CSV,
	Encoding:      "utf-8",
	Checksum:      file.MD5,
	ItemPrototype: &model.StudentData{},
}
var FTP = &file.FTPFileSystem{
	Hort:        "127.0.0.1",
	Port:        3306,
	User:        "root",
	Password:    "root",
	ConnTimeout: time.Second,
}
var CopyFileToFtp = file.FileMove{
	FromFileName:  "res/{date,yyyyMMdd}/studentdata.csv",
	FromFileStore: &file.LocalFileSystem{},
	ToFileStore:   FTP,
	ToFileName:    "trade/{date,yyyyMMdd}/Cstudentdata.csv",
}
var CopyChecksumFileToFtp = file.FileMove{
	FromFileName:  "res/{date,yyyyMMdd}/studentdata.csv.md5",
	FromFileStore: &file.LocalFileSystem{},
	ToFileStore:   FTP,
	ToFileName:    "trade/{date,yyyyMMdd}/Cstudentdata.csv.md5",
}

// creating connection with the new DB
type SqlStudent struct {
	db *sql.DB
}

func (p *SqlStudent) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	for _, item := range items {
		studData := item.(model.StudentData)
		_, err := p.db.Exec("INSERT INTO studentdata(studName,studRollno,studBranch) values (?,?,?)",
			studData.StudName, studData.StudRollno, studData.StudBranch)
		if err != nil {
			return gobatch.NewBatchError(gobatch.ErrCodeDbFail, "Error while insert studentdata into db ", err)
		}
	}
	return nil
}
