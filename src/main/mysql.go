package main

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

// URLDetail contains the detail of the shortURL
type URLDetail struct {
	Id                  string    `json:"id"`
	URL                 string    `json:"url"`
	ShortCode           string    `json:"short_code"`
	CreatedBy           uint32    `json:"created_by"`
	CreatedAt           time.Time `json:"created_at"`
	ExpirationInMinutes uint32    `json:"expiration_in_minutes"`
}

func NewMySQL(driverName, dataSourceName string) *MySQL {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return &MySQL{DB: db}
}

func (m *MySQL) InsertURLDetail(d *URLDetail) (int64, error) {
	//插入数据
	stmt, err := m.DB.Prepare(
		"INSERT url_detail SET url=?, short_code=?, created_by=?, created_at=?, expiration_in_minutes=?")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(d.URL, d.ShortCode,
		d.CreatedBy, d.CreatedAt, d.ExpirationInMinutes)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}
func (m *MySQL) QueryUrlDetail(shortCode string) (*URLDetail, error) {
	rows, err := m.DB.Query("SELECT id, url, short_code, created_by, created_at, expiration_in_minutes FROM url_detail where short_code=?", shortCode)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		d := URLDetail{}
		err := rows.Scan(&d.Id, &d.URL, &d.ShortCode, &d.CreatedBy, &d.CreatedAt, &d.ExpirationInMinutes)
		if err != nil {
			return nil, err
		}
		if rows.Next() {
			return nil, errors.New("Should not have more then 1 URLDetail for shortCode: " + shortCode)
		}
		return &d, nil
	}
	return nil, nil
}
