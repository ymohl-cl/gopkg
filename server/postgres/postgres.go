package postgres

import (
	"database/sql"

	"github.com/ymohl-cl/gopkg/server/postgres/config"

	// import blank to get the postgesql sriver in sql management
	_ "github.com/lib/pq"
)

// Postgres driver
type Postgres interface {
	Close() error
	Create(query QueryInsert) (interface{}, error)
	Update(query string, values []interface{}) ([]interface{}, error)
	Read(query QuerySelect, scanner func(*sql.Rows) error) error
	Delete(query string, values []interface{}) ([]interface{}, error)
}

type postgres struct {
	driver *sql.DB
}

// New Provide a simple new sql client
func New(appName string) (Postgres, error) {
	var c config.Config
	var err error
	var p postgres

	if c, err = config.New(appName); err != nil {
		return nil, err
	}
	str := "user=" + c.User + " "
	str += "password=" + c.Password + " "
	str += "dbname=" + c.DBName + " "
	str += "host=" + c.Host + " "
	str += "port=" + c.Port + " "
	str += "sslmode=disable"
	if p.driver, err = sql.Open(config.DriverName, str); err != nil {
		return nil, err
	}
	if err = p.driver.Ping(); err != nil {
		return nil, err
	}
	return &p, nil
}

// Close the postgres driver
func (p postgres) Close() error {
	return p.driver.Close()
}

// Create to crud method
func (p postgres) Create(query QueryInsert) (interface{}, error) {
	var item interface{}
	var err error

	if err = p.driver.QueryRow(query.String(), query.Values()...).Scan(&item); err != nil {
		return nil, err
	}
	return item, nil
}

// Update to crud method
func (p postgres) Update(query string, values []interface{}) ([]interface{}, error) {
	var err error

	if _, err = p.driver.Exec(query, values...); err != nil {
		return nil, err
	}
	return nil, nil
}

// Read to crud method
func (p postgres) Read(query QuerySelect, scanner func(*sql.Rows) error) error {
	var rows *sql.Rows
	var err error

	if rows, err = p.driver.Query(query.String(), query.Values()...); err != nil {
		return err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return err
	}

	for rows.Next() {
		if err = scanner(rows); err != nil {
			return err
		}
	}
	return nil
}

// Delete to crud method
func (p postgres) Delete(query string, values []interface{}) ([]interface{}, error) {
	var err error

	if _, err = p.driver.Exec(query, values...); err != nil {
		return nil, err
	}
	return nil, nil
}
