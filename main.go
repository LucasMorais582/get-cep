package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

const URL = "http://viacep.com.br/ws/"

func main() {
	for _, cep := range os.Args[1:] {
		request, errorRequest := http.Get(URL + cep + "/json/")
		if errorRequest != nil {
			fmt.Fprintf(os.Stderr, "Request error: %v\n\n", errorRequest)
			return
		}
		defer request.Body.Close()

		response, errorResponse := io.ReadAll(request.Body)
		if errorResponse != nil {
			fmt.Fprintf(os.Stderr, "Response error: %v\n\n", errorResponse)
			return
		}

		var data CEP
		errorData := json.Unmarshal(response, &data)
		if errorData != nil {
			fmt.Fprintf(os.Stderr, "Data error: %v\n\n", errorData)
			return
		}

		fmt.Println(data)

		// creating file
		file, errorFile := os.Create("city.txt")
		if errorFile != nil {
			fmt.Fprintf(os.Stderr, "File creation error: %v\n", errorFile)
			return
		}
		defer file.Close()

		_, errorFile = file.WriteString(fmt.Sprintf(" CEP: %s\n Location: %s\n UF: %s\n", data.Cep, data.Localidade, data.Uf))

		fmt.Printf("File created successuly!\n\n")

	}
}
