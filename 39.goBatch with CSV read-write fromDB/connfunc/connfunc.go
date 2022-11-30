package connfunc

import "database/sql"

func InitDbJob() *sql.DB {
	//set db for gobatch to store job&step execution context
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

// detup db for file
func InitDbStud() *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/student?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
