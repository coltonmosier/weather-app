package handlers

import (
	"fmt"
	"net/http"
)

func handle() {
	fmt.Println("Handler file")
}

func Home() *http.Handler {

}
func Weather() *http.Handler {

}
