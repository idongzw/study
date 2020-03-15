/*
 * @Author: dzw
 * @Date: 2020-03-13 14:37:54
 * @Last Modified by:   dzw
 * @Last Modified time: 2020-03-13 14:37:54
 */

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:qaz@tcp(127.0.0.1:3306)/chat"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open sql failed,", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("link db failed,", err)
		return
	}
	fmt.Println("open sql success")
}
