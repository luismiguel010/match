package routers

import (
	"fmt"
	"net/http"

	"github.com/luismiguel010/match/match"
)

func MakeMatch(w http.ResponseWriter, r *http.Request) {
	match.TotalMatch()
	respuesta := "Match exitoso"
	fmt.Fprintf(w, "Respuesta: %s", respuesta)
}
