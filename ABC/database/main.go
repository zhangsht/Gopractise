package main

import (
	"os"
	"io"
	"strings"
	"bufio"
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// open database and query
	db, err := sql.Open("mysql", "root:GZHHSH2017android@tcp(localhost:3306)/test?charset=utf8")
	checkError(err)
	age := 0
	rows, err := db.Query("SELECT name, age FROM gotest WHERE age>?", age)
	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age)
		checkError(err)
		fmt.Printf("%s is %d\n", name, age)
	}
}

func run(db *sql.DB, name string) {
	f, err := os.Open(name)
	checkError(err)
	defer f.Close()
	tx, err := db.Begin()
	buf := bufio.NewReader(f)
	LineNumber := 0
	for {
		LineNumber++
		line, err := buf.ReadString('\n')
		line = strings.Replace(line, "set", "update", -1)
		println(line)
		if err == io.EOF {
			tx.Exec(line)
			tx.Commit()
			return
		} else {
			tx.Exec(line)
		}
	}

	/* // prepare stmt
	age := 27
	stmt, err := db.Prepare("SELECT name FROM users WHERE age = ?")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(age) */
}

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

/* func Loop(db, *sql.DB, dirs [string])  {
	for _, v := range dirs {
		if strings.HasSuffix(v, ".sql") {
			run(db, v)
		} else (IsDir(v)) {
			dirs := ListDir(v)
			Loop(db, dirs)
		}
	}
	ch <- 0
} */

/* func GetRst(db *sql.DB, sql string) (eles []map[string]interface{}) {
    rst, err := db.Query(sql)
    Check(err)
    columns, err := rst.Columns()
    count := len(columns)

    values := make([]string, count)
    ptr := make([]interface{}, count)
    for rst.Next() {
        for i := 0; i < count; i++ {
            ptr[i] = &values[i]
        }
        rst.Scan(ptr...)
        entry := make(map[string]interface{}, 15)
        for i, col := range columns {
            val := values[i]
            //println(val)
            entry[col] = val
        }
        eles = append(eles, entry)

    }
    return eles
} */