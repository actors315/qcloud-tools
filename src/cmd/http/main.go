package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",Welcome)
	http.ListenAndServe(":8888",nil)
}

func Welcome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<div>Welcome ~~</div>")
}

