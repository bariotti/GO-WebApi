package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		nomes := r.URL.Query()["nome"]
		indice := r.URL.Query()["indice"]

		if len(nomes) > 0 {

			var indice_int int
			var err error

			if len(indice) > 0 {

				indice_int, err = strconv.Atoi(indice[0])

				if err != nil {
					indice_int = 0
				}

				if indice_int >= len(nomes) {
					indice_int = len(nomes) - 1
				}

				if indice_int < 0 {
					indice_int = 0
				}

			} else {
				indice_int = 0
			}

			nome := nomes[indice_int]

			w.Write([]byte("Hello " + nome))

		} else {
			w.Write([]byte("Hello World"))
		}
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("Erro ao subir servidor Web. " + err.Error())
	}
}
