package helpers

import(
	//"net/http"
    "database/sql"
    _ "github.com/lib/pq"
    //_ "github.com/truongsinh/pq"
)

var Db = SetupDB()

func SetupDB() *sql.DB {
    db, err := sql.Open("postgres", "dbname=publish user=postgres password=hathat sslmode=disable")
    PanicIf(err)
    return db
}