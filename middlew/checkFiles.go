package middlew

import (
	"net/http"
	"os"
)

func CheckFiles(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, errBuy := os.Open("./solicitudes_compra.cvs")
		_, errSale := os.Open("./solicitudes_venta.cvs")
		switch {
		case errBuy != nil:
			http.Error(w, "El archivo de solicitudes compra no se encuentra en la ruta especificada", http.StatusInternalServerError)
			return
		case errSale != nil:
			http.Error(w, "El archivo de solicitudes venta no se encuentra en la ruta especificada", http.StatusInternalServerError)
			return
		default:
			next.ServeHTTP(w, r)
		}
	}
}
