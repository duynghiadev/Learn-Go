// package db

// // https://github.com/rubenv/sql-migrate
// import (
// 	"fmt"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/labstack/gommon/log"
// 	_ "github.com/lib/pq"
// )

// type Sql struct {
// 	Db       *sqlx.DB
// 	Host     string
// 	Port     int
// 	UserName string
// 	Password string
// 	DbName   string
// }

// func (s *Sql) Connect() {
// 	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		s.Host, s.Port, s.UserName, s.Password, s.DbName)

// 	s.Db = sqlx.MustConnect("postgres", dataSource)

// 	if err := s.Db.Ping(); err != nil {
// 		log.Error(err.Error())
// 		return
// 	}

// 	fmt.Println("Connect database ok")
// }

// func (s *Sql) Close() {
// 	s.Db.Close()
// }

package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     int
	UserName string
	Password string
	DbName   string
}

func (s *Sql) Connect() error {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.UserName, s.Password, s.DbName)

	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		log.Errorf("Failed to connect to database: %v", err)
		return err
	}

	if err := db.Ping(); err != nil {
		log.Errorf("Database ping failed: %v", err)
		return err
	}

	s.Db = db
	fmt.Println("Connect database ok")
	return nil
}

func (s *Sql) Close() {
	if s.Db != nil {
		if err := s.Db.Close(); err != nil {
			log.Errorf("Error closing database connection: %v", err)
		} else {
			fmt.Println("Database connection closed")
		}
	}
}
