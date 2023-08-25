package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	entity "github.com/brunobolting/go-rinha-backend/domain"

	"github.com/lib/pq"
)

type PersonPosgreSql struct {
	db *sql.DB
}

func NewPersonPostgreSql(db *sql.DB) *PersonPosgreSql {
	return &PersonPosgreSql{
		db,
	}
}

func (r *PersonPosgreSql) Get(id string) (*entity.Person, error) {
	stmt, err := r.db.Prepare("SELECT id, nickname, name, birthdate, stack FROM persons WHERE id = $1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	row := stmt.QueryRowContext(ctx, id)

	p := &entity.Person{}

	err = row.Scan(&p.ID, &p.Nickname, &p.Name, &p.Birthdate, pq.Array(&p.Stack))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entity.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PersonPosgreSql) Create(p *entity.Person) error {
	query := "INSERT INTO persons (id, nickname, name, birthdate, stack) VALUES ($1, $2, $3, $4, $5)"
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	_, err := r.db.ExecContext(ctx, query, p.ID, p.Nickname, p.Name, p.Birthdate, pq.Array(p.Stack))
	return err
}

func (r *PersonPosgreSql) Find(q string) ([]*entity.Person, error) {
	query := "SELECT id, nickname, name, birthdate, stack FROM persons WHERE search ~ $1 LIMIT 50"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := r.db.QueryContext(ctx, query, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []*entity.Person

	for rows.Next() {
		p := &entity.Person{}
		err = rows.Scan(&p.ID, &p.Nickname, &p.Name, &p.Birthdate, pq.Array(&p.Stack))
		if err != nil {
			return nil, err
		}

		persons = append(persons, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}

func (r *PersonPosgreSql) Count() (int32, error) {
	stmt, err := r.db.Prepare("SELECT COUNT(id) FROM persons")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var c int32

	err = stmt.QueryRow().Scan(&c)
	if err != nil {
		return 0, err
	}

	return c, nil
}
