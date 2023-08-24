package repository

import (
	"database/sql"
	"fmt"

	entity "github.com/brunobolting/go-rinha-backend/domain"

	_ "github.com/lib/pq"
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

	row := stmt.QueryRow(id)

	p := &entity.Person{}

	err = row.Scan(&p.ID, &p.Nickname, &p.Name, &p.Birthdate, &p.Stack)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PersonPosgreSql) Create(p *entity.Person) error {
	query := "INSERT INTO persons (id, nickname, name, birthdate, stack) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, p.ID, p.Nickname, p.Name, p.Birthdate, p.Stack)
	return err
}

func (r *PersonPosgreSql) Find(q string) ([]*entity.Person, error) {
	stmt, err := r.db.Prepare("SELECT id, nickname, name, birthdate, stack FROM persons WHERE search_values LIKE $1")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(fmt.Sprintf("%%%s%%", q))
	rows.Close()

	var persons []*entity.Person

	for rows.Next() {
		p := &entity.Person{}
		err = rows.Scan(&p.ID, &p.Nickname, &p.Name, &p.Birthdate, &p.Stack)
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
