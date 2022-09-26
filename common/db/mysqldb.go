package mysqldb

import (
	"database/sql"
	"fmt"
)

const dbuser = "root"
const dbpass = "Mobile@97701"
const dbname = "test"

func main() {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error

	}
	defer db.Close()
}
