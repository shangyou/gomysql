package gomysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Db *sql.DB
}

func (this *Mysql) Init(dns string) error {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}

	this.Db = db
	return nil
}

func (this *Mysql) query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := this.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, err
}

func (this *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	res, err := this.Db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (this *Mysql) Insert(query string, args ...interface{}) (int64, error) {
	res, err := this.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	return id, err
}

func (this *Mysql) Update(query string, args ...interface{}) (int64, error) {
	res, err := this.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()

	return count, nil
}

func (this *Mysql) Delete(query string, args ...interface{}) (int64, error) {
	res, err := this.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()

	return count, err
}

func (this *Mysql) Fetchrow(query string, args ...interface{}) map[string]string {
	row, _ := this.query(query, args...)
	columns, _ := row.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]string, len(columns))

	for i, _ := range scanArgs {
		scanArgs[i] = &values[i]
	}

	for row.Next() {
		err := row.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
	}
	rs := make(map[string]string)
	for i, _ := range values {
		rs[columns[i]] = values[i]
	}

	return rs
}

func (this *Mysql) Fetchrows(query string, args ...interface{}) map[int]map[string]string {
	row, _ := this.query(query, args...)
	columns, _ := row.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]string, len(columns))

	for i, _ := range scanArgs {
		scanArgs[i] = &values[i]
	}

	rs := make(map[int]map[string]string)
	i := 0
	for row.Next() {
		err := row.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		r := make(map[string]string)
		for k, v := range values {
			r[columns[k]] = v
		}

		rs[i] = r
		i++
	}

	return rs
}
