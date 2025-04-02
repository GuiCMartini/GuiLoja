package main

import (
	"net/http"

	"github.com/GuiCMartini/GuiLoja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
