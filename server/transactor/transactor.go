package transactor

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/ymohl-cl/gopkg/server/postgres"

	"github.com/ymohl-cl/gopkg/server/model"
)

// user transactions key to postgres database
const (
	postgresUserIDKey       = "id"
	postgresUserPseudoKey   = "pseudo"
	postgresUserPasswordKey = "password"
	postgresUserNameKey     = "name"
	postgresUserLastNameKey = "last_name"
	postgresUserAgeKey      = "age"
	postgresUserGenreKey    = "genre"
	postgresUserEmailKey    = "email"
)

// user transaction table name to postgres database
const (
	postgresUserTable = "users"
)

// Transactor service manage the request to postgres driver
type Transactor interface {
	ConnectUser(username, password string) (model.User, error)
	ReadUsers(filter model.User) ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
}

type transactor struct {
	postgresDriver postgres.Postgres
}

// New return a transactor service
func New(appName string) (Transactor, error) {
	var t transactor
	var err error

	if t.postgresDriver, err = postgres.New(appName); err != nil {
		return nil, err
	}
	return &t, nil
}

// ConnectUser with the authentication parameters
func (t transactor) ConnectUser(username, password string) (model.User, error) {
	var err error
	var users []model.User
	var user model.User

	if users, err = t.ReadUsers(model.User{Pseudo: username, Password: password}); err != nil {
		return model.User{}, err
	}
	count := len(users)
	if count > 1 {
		return model.User{}, errors.New("impossible multiple result found to connect user")
	} else if count == 1 {
		user = users[0]
	}
	return user, nil
}

// ReadUsers search in postgres sql users matches with the user parameter filters
func (t transactor) ReadUsers(user model.User) ([]model.User, error) {
	var err error

	// build query
	query := postgres.NewQuerySelect([]string{
		postgresUserIDKey,
		postgresUserPseudoKey,
		postgresUserNameKey,
		postgresUserLastNameKey,
		postgresUserAgeKey,
		postgresUserGenreKey,
		postgresUserEmailKey,
	}).From(
		postgresUserTable,
	)
	if user.Identifier != "" {
		query = query.Where(postgresUserIDKey, user.Identifier)
	}
	if user.Pseudo != "" {
		query = query.Where(postgresUserPseudoKey, user.Pseudo)
	}
	if user.Name != "" {
		query = query.Where(postgresUserNameKey, user.Name)
	}
	if user.LastName != "" {
		query = query.Where(postgresUserLastNameKey, user.LastName)
	}
	if user.Password != "" {
		query = query.Where(postgresUserPasswordKey, user.Password)
	}
	if user.Age != 0 {
		query = query.Where(postgresUserAgeKey, user.Age)
	}
	if user.Genre != "" {
		query = query.Where(postgresUserGenreKey, user.Genre)
	}
	if user.Email != "" {
		query = query.Where(postgresUserEmailKey, user.Email)
	}

	// read result
	var users []model.User
	if err = t.postgresDriver.Read(query, func(rows *sql.Rows) error {
		var u model.User
		var err error

		if err = rows.Scan(
			&u.Identifier,
			&u.Pseudo,
			&u.Name,
			&u.LastName,
			&u.Age,
			&u.Genre,
			&u.Email,
		); err != nil {
			return err
		}
		users = append(users, u)
		return nil
	}); err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser insert the specific user in postgres sql
// Return the user with the identifier set
func (t transactor) CreateUser(user model.User) (model.User, error) {
	var err error

	// build query
	query := postgres.NewQueryInsert([]string{
		postgresUserPseudoKey,
		postgresUserPasswordKey,
		postgresUserNameKey,
		postgresUserLastNameKey,
		postgresUserAgeKey,
		postgresUserGenreKey,
		postgresUserEmailKey,
	}).Into(postgresUserTable).WithReturn(postgresUserIDKey)
	query = query.WithValues([]interface{}{
		user.Pseudo,
		user.Password,
		user.Name,
		user.LastName,
		user.Age,
		user.Genre,
		user.Email,
	})
	// record user and get the identifier
	var untypedID interface{}
	if untypedID, err = t.postgresDriver.Create(query); err != nil {
		return model.User{}, err
	}
	var bufID []byte
	var ok bool
	if bufID, ok = untypedID.([]byte); !ok {
		return model.User{}, errors.New("can't extract identifier to the new item")
	}
	var id uuid.UUID
	if id, err = uuid.ParseBytes(bufID); err != nil {
		return model.User{}, err
	}
	user.Identifier = id.String()
	// to secure side effect, remove password here
	user.Password = ""
	return user, nil
}
