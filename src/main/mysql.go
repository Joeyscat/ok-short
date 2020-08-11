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

// GenURLDetail contains the detail of the shortURL
type GenURLDetail struct {
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

func (m *MySQL) InsertGenURLDetail(d *GenURLDetail) (int64, error) {
	stmt, err := m.DB.Prepare(
		"INSERT short_url_gen_detail SET url=?, short_code=?, created_by=?, created_at=?, expiration_in_minutes=?")
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
func (m *MySQL) QueryGenURLDetail(shortCode string) (*GenURLDetail, error) {
	rows, err := m.DB.Query("SELECT id, url, short_code, created_by, created_at, expiration_in_minutes FROM short_url_gen_detail where short_code=?", shortCode)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		d := GenURLDetail{}
		err := rows.Scan(&d.Id, &d.URL, &d.ShortCode, &d.CreatedBy, &d.CreatedAt, &d.ExpirationInMinutes)
		if err != nil {
			return nil, err
		}
		if rows.Next() {
			return nil, errors.New("Should not have more then 1 GenURLDetail for shortCode: " + shortCode)
		}
		return &d, nil
	}
	return nil, nil
}

func (m *MySQL) InsertShortURLVisitedLog(log *ShortURLVisitedLog) (int64, error) {
	stmt, err := m.DB.Prepare(
		"INSERT short_url_visited_log SET remote_addr=?, short_code=?, ua=?, cookie=?, visitor_id=?, visited_at=?")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(log.RemoteAddr, log.ShortCode, log.UA, log.Cookie, log.VisitorId, log.VisitedAt)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

func (m *MySQL) QueryShortURLVisitedLog(shortCode string) ([]ShortURLVisitedLog, error) {
	rows, err := m.DB.Query("SELECT id, remote_addr, short_code, ua, cookie, visitor_id, visited_at FROM short_url_visited_log where short_code=?", shortCode)
	if err != nil {
		return nil, err
	}
	var logs []ShortURLVisitedLog

	for rows.Next() {
		l := &ShortURLVisitedLog{}
		err := rows.Scan(&l.Id, &l.RemoteAddr, &l.ShortCode, &l.UA, &l.Cookie, &l.VisitorId, &l.VisitedAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, *l)
	}
	return logs, nil
}
