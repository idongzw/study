/*
 * @Author: dzw
 * @Date: 2020-03-27 17:23:43
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-27 17:33:30
 */

package main

import (
	"flag"
	"net/http"
)

func main() {
	path := flag.String("filePath", "/", "file path")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*path)))
	http.ListenAndServe(":8080", nil)
}
