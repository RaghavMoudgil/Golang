package datareader

import (
	"batch/model"
	"database/sql"
	"fmt"

	"github.com/chararch/gobatch"
)

type StudRreader struct {
	db *sql.DB
}

func (r *StudRreader) Open(execution *gobatch.StepExecution) gobatch.BatchError {
	return nil
}
func (r *StudRreader) Close(exection *gobatch.StepExecution) gobatch.BatchError {
	return nil
}
func (r *StudRreader) ReadKeys() (interface{}, error) {
	rows, err := r.db.Query("select id from studentdata")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []interface{}
	var id int64
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		result = append(result, id)
	}
	return result, nil
}
func (r *StudRreader) ReadItem(key interface{}) (interface{}, error) {
	id := int64(0)
	switch s := key.(type) {
	case int64:
		id = s
	case float64:
		id = int64(s)
	default:
		return nil, fmt.Errorf("Key type error, type:%T, vaule:%v", key, key)
	}
	rows, err := r.db.Query("select studName,studRollno,studBranch from studentdata where id =?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	studData := &model.StudentData{}
	if rows.Next() {
		err = rows.Scan(&studData.StudName, &studData.StudRollno, &studData.StudBranch)
		if err != nil {
			return nil, err
		}
	}
	return studData, nil
}
