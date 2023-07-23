package component

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FianGumilar/auth-api-token/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionDatabase(conf config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		conf.DB.User,
		conf.DB.Pass,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed Connect to Database %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed Connect to Database %s", err)
	}

	return db
}
