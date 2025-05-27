package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)


type Storage struct {
	Db *pgxpool.Pool
}

// type Task struct {
//   ID  int
//   Title string
//   Description string
//   Status int
//   CreatedAt int64
//   ClosedAt int64
//   UpdatedAt int64
//   AuthorId int
//   AssignedId int
// }


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


// создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
  var context = context.Background()
  var defaultStatus = 1 //статус задачи - создана

	err := s.db.QueryRow(context, `
		INSERT INTO tasks (title, description, status, author_id, assigned_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;
		`,
		t.Title,
		t.Description,
    defaultStatus,
    t.AuthorId,
    t.AssignedId,
	).Scan(&id)
	return id, err
}