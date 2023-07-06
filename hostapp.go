package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	myHost, err := os.Hostname() // Here we decided to explicitly handle all the return values
// 	if err != nil {
// 		log.Fatal(err) // Send error message to log output
// 	}
// 	fmt.Fprintln(w, "Greetings World WFH, 世界. You are connected to host", myHost)
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	myHost, err := os.Hostname() // Here we decided to explicitly handle all the return values
	if err != nil {
		log.Fatal(err) // Send error message to log output
	}
	fmt.Fprintln(w, "Greetings World WFH, 世界. You are connected to host", myHost)
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello endpoint")
    })

    http.HandleFunc("/yes", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Yes endpoint")
    })

    http.HandleFunc("/amen", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Amen endpoint")
    })

    http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "payment endpoint")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}

// This is to test git revert 
// This is to test git revert 2