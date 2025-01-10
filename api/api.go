package api

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// t := Response{
	// 	Message: "yo",
	// }
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "TEST HANDLER IS READY")
}