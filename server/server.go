// Run for debugging puposes:
// > go run parse_log.go access.log
// Build and run:
// > go build server.go
// > ./parse_log access.log
package main

import (
	"fmt"
	"log"
	"net/http"
	"server/handlers"
	"strconv"
	// Use rand and  time for testing
	// "math/rand"
	// "time"
)




func main() {
	// Use rand and  time for testing
	// rand.Seed(time.Now().UnixNano())
	// var port int = rand.Intn(10) + 8000
	const port int = 8080
	http.HandleFunc("/person", handlers.HandlePerson)
	fmt.Println("Listening on  port ", port)

	// fmt.Println(colonPort)
	// http.ListenAndServe(colonPort, nil)
	// https://egel.github.io/development/2018/06/16/accept-incoming-network-connections.html
	log.Fatal(http.ListenAndServe(("localhost:" + strconv.Itoa(port)), nil))
}

