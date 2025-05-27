package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)


type Storage struct {
	Db *pgxpool.Pool
}


func New(connString string) (*Storage, error) {
  	var context = context.Background()
	db, err := pgxpool.Connect(context, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = db.Ping(context)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	s := Storage{
		Db: db,
	}

	return &s, nil
}

func (db *Storage) Close() {
	db.Db.Close()
}