package controller

import "net/http"

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithErr(w, 500, "Something went wrong")
}
