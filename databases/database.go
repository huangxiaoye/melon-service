package databases

import (
	"log"
	"os"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"

	"github.com/spf13/viper"
)

type DB struct {
	sqlbuilder.Database
}

var db *DB

func Init() {
	var settings = postgresql.ConnectionURL{
		Host:     viper.GetString(`database.host`),
		Database: viper.GetString(`database.database`),
		User:     viper.GetString(`database.user`),
		Password: viper.GetString(`database.password`),
		//Options:  map[string]string{"charset": viper.GetString(`database.options.charset`)},
	}

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	if "1" == os.Getenv("UPPERIO_DB_DEBUG") {
		sess.SetLogging(true)
	}
	if err = sess.Ping(); err != nil {
		log.Fatal(err)
	}
	db = &DB{sess}
}

//GetDB ...
func GetDB() *DB {
	return db
}
