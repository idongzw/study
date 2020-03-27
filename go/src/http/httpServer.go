/*
 * @Author: dzw
 * @Date: 2020-03-27 17:23:43
 * @Last Modified by: dzw
 * @Last Modified time: 2020-03-27 18:11:25
 */

package main

import (
	"flag"
	"net/http"
)

func main() {
	path := flag.String("filepath", "/", "file path")
	port := flag.String("port", "8080", "file server port")
	flag.Parse()

	// http server
	http.Handle("/", http.FileServer(http.Dir(*path)))
	http.ListenAndServe(":"+*port, nil)
}
