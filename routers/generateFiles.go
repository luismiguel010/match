package routers

import (
	"net/http"
	"strconv"

	generatorfiles "github.com/luismiguel010/match/generatorFiles"
)

func GeneratorFiles(w http.ResponseWriter, r *http.Request) {
	AMOUNTREAD := r.URL.Query().Get("amount")
	amount, err := strconv.Atoi(AMOUNTREAD)
	if err != nil {
		http.Error(w, "Ingrese un número válido por favor", http.StatusBadRequest)
		return
	}
	if amount < 1 {
		http.Error(w, "Debe enviar una cantidad de solicitudes mayor a 0", http.StatusBadRequest)
		return
	}
	generatorfiles.GeneratorFiles(amount)
}
