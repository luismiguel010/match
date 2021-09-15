package routers

import (
	"net/http"

	"github.com/luismiguel010/match/match"
)

func MakeMatch(w http.ResponseWriter, r *http.Request) {
	match.TotalMatch()
}
