package database

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Log struct {
	UUID        string    `db:"uuid"`
	SessionUUID string    `db:"sessionUUID"`
	Doc         string    `db:"doc"`
	Created     time.Time `db:"created"`
}

type Session struct {
	UUID    string    `db:"uuid"`
	Created time.Time `db:"created"`
	Closed  time.Time `db:"closed"`
	User    string    `db:"user"`
	Remote  string    `db:"remote"`
}

func (l *Log) Insert(db *sqlx.DB) (err error) {
	sql := "INSERT INTO Log SELECT generateUUIDv4(), :sessionUUID, :doc, :created"

	_, err = db.NamedExec(sql, &l)
	return err
}

// get UUID via clickhouse
func GenUUID(db *sqlx.DB) (uuid string, err error) {
	err = db.Get(&uuid, "SELECT generateUUIDv4()")
	return uuid, err
}

func (s *Session) Insert(db *sqlx.DB) (err error) {
	sql := "INSERT INTO Session VALUES (toUUID(:uuid), :created, :closed, :user, :remote)"

	_, err = db.NamedExec(sql, &s)
	return err
}
