package handlers

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

func Test(_ context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok!")
}
