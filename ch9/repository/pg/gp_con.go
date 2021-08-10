package pg_repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/amupxm/golang-backend-training/ch9/domain"
	_ "github.com/lib/pq"
)

type (
	pg_db struct {
		db *sql.DB
	}
	PG_DB interface {
		UsersGetAll() (users *[]domain.User, err error)
		CreateUser(user *domain.User) (*domain.User, error)
		UpdateUser(user *domain.User) (*domain.User, error)
		DeleteUser(userId int64) error
	}
)

func CreateConnection(username, password, host, database string) PG_DB {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, database))
	if err != nil {
		panic(err)
	}
	return &pg_db{db}
}

func (pg *pg_db) CloneConn() {
	defer pg.db.Close()
}

func (pg *pg_db) UsersGetAll() (users *[]domain.User, err error) {
	rs, err := pg.db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	res := []domain.User{}
	for rs.Next() {
		var user domain.User
		err := rs.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return &res, nil
			}
			return nil, err
		}
		res = append(res, user)
	}
	return &res, nil
}

func (pg *pg_db) CreateUser(user *domain.User) (*domain.User, error) {
	var id int64
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := pg.db.QueryRow(`INSERT INTO users (username, email, password, created_at, updated_at)
	 VALUES ($1, $2, $3, $4, $5) RETURNING id`, user.UserName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		return user, err
	}
	user.Id = id
	return user, nil
}

func (pg *pg_db) UpdateUser(user *domain.User) (*domain.User, error) {
	user.UpdatedAt = time.Now()
	err := pg.db.QueryRow(`UPDATE users SET username = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5`,
		user.UserName, user.Email, user.Password, user.UpdatedAt, user.Id).Err()
	return user, err
}

func (pg *pg_db) DeleteUser(userId int64) error {
	return pg.db.QueryRow(`DELETE FROM users WHERE id = $1`, userId).Err()
}
