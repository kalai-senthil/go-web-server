package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, `<html lang="en">  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Go Server</title>
  </head>
  <body>
    <h1>A simple Web Server</h1>
    <p>Developed using Go</p>
  </body>
</html>
`)
}
func main() {
	http.HandleFunc("/", handle)
	if len(os.Args) == 2 {
		http.ListenAndServe(":"+os.Args[1], nil)
	} else {
		log.Fatal("Provide a port number")
	}
}
