// sorter project main.go
package main

import "log"

//import "flag"
import "fmt"
import "bufio"
import "os"
import "database/sql"
import _ "github.com/wendal/go-oci8"

//var sqlscript *string = flag.String("sql", "", "The sqlscript that will be executed.")
var inputReader *bufio.Reader
var err error

var dbname string = ""
var user string = ""
var password string = ""
var dbcp string = ""
var sqlscript string = ""

func main() {
	inputReader = bufio.NewReader(os.Stdin)

	log.Println("Please input dbname.")

	dbname, err = inputReader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	dbname = dbname[:len(dbname)-2]

	log.Println("Please input user.")

	user, err = inputReader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	user = user[:len(user)-2]

	log.Println("Please input password.")

	password, err = inputReader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	password = password[:len(password)-2]

	dbcp = fmt.Sprintf("%s/%s@%s", user, password, dbname)

	log.Println("Oracle Driver Connecting....", dbcp)
	//用户名/密码@实例名 如system/123456@orcl、sys/123456@orcl
	db, err := sql.Open("oci8", dbcp)
	if err != nil {
		log.Fatal(err)
		panic("数据库连接失败")
	} else {
		defer db.Close()
	}

	log.Println("Please input sqlscript.")

	sqlscript, err = inputReader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	stmt, _ := db.Prepare(sqlscript)

	defer stmt.Close()

	result, err := stmt.Exec()

	if err != nil {
		log.Fatal(err)
	}

	count, _ := result.RowsAffected()
	log.Printf("result count:%d", count)
}
