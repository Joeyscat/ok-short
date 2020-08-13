package store

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/joeyscat/ok-short/common"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(driverName, dataSourceName string) *MySQL {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	return &MySQL{DB: db}
}

func (m *MySQL) InsertLink(d *Link) (int64, error) {
	stmt, err := m.DB.Prepare(
		"INSERT ok_link SET origin_url=?, short_code=?, created_by=?, created_at=?, expiration_in_minutes=?")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(d.OriginURL, d.ShortCode,
		d.CreatedBy, d.CreatedAt, d.Exp)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

func (m *MySQL) GetLinkInfo(shortCode string) (*Link, error) {
	rows, err := m.DB.Query("SELECT id, origin_url, short_code, created_by, created_at, expiration_in_minutes FROM ok_link where short_code=?", shortCode)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		d := Link{}
		err := rows.Scan(&d.Id, &d.URL, &d.ShortCode, &d.CreatedBy, &d.CreatedAt, &d.Exp)
		if err != nil {
			return nil, err
		}
		if rows.Next() {
			return nil, errors.New("Should not have more then 1 Link for shortCode: " + shortCode)
		}
		return &d, nil
	}
	return nil, nil
}

func (m *MySQL) GetLinkList(offset, size uint32) (*[]Link, uint32, uint32, error) {
	rows, err := m.DB.Query("SELECT id, origin_url, short_code, created_by, created_at, expiration_in_minutes FROM ok_link limit ?,?", offset, size)
	if err != nil {
		return nil, 0, 0, err
	}
	var links []Link
	for rows.Next() {
		l := Link{}
		err := rows.Scan(&l.Id, &l.URL, &l.ShortCode, &l.CreatedBy, &l.CreatedAt, &l.Exp)
		if err != nil {
			return nil, 0, 0, err
		}
		links = append(links, l)
	}
	// Total
	rows, err = m.DB.Query("SELECT count(id) FROM ok_link")
	if err != nil {
		return nil, 0, 0, err
	}
	var total uint32
	if rows.Next() {
		err = rows.Scan(&total)
		if err != nil {
			return nil, 0, 0, err
		}
	}

	return &links, uint32(len(links)), total, nil
}

func (m *MySQL) InsertLinkVisitedLog(log *LinkVisitedLog) (int64, error) {
	stmt, err := m.DB.Prepare(
		"INSERT ok_link_visited_log SET remote_addr=?, short_code=?, ua=?, cookie=?, visitor_id=?, visited_at=?")
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

func (m *MySQL) QueryLinkVisitedLog(shortCode string) ([]LinkVisitedLog, error) {
	rows, err := m.DB.Query("SELECT id, remote_addr, short_code, ua, cookie, visitor_id, visited_at FROM short_url_visited_log where short_code=?", shortCode)
	if err != nil {
		return nil, err
	}
	var logs []LinkVisitedLog

	for rows.Next() {
		l := &LinkVisitedLog{}
		err := rows.Scan(&l.Id, &l.RemoteAddr, &l.ShortCode, &l.UA, &l.Cookie, &l.VisitorId, &l.VisitedAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, *l)
	}
	return logs, nil
}
