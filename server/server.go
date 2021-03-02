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

	colonPort := fmt.Sprint(":", port)
	fmt.Println(colonPort)
	http.ListenAndServe(colonPort, nil)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

