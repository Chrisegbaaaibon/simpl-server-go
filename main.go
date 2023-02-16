package main 

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() %v", err)
		return
	}

	fmt.Fprintf(w, "Request successful")
	name := r.FormValue("Name")
	address := r.FormValue("Address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello world")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/hello",  helloHandler)


	fmt.Printf("Starting server on port 4500\n")
	if err := http.ListenAndServe(" : 4500", nil ); err !=nil {
		log.Fatal(err)
	}
}