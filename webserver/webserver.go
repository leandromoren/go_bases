package webserver

import (
	"fmt"
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/", home)
	fmt.Println("Corriendo servidor...")
	http.ListenAndServe(":3000", nil)
}

// Sintaxis basica de webserver para procesar una peticion HTTP
func home(writer http.ResponseWriter, req *http.Request) {
	// Sirve a la web un archivo
	http.ServeFile(writer, req, "./webserver/index.html")
}
