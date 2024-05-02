package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var VIA_CEP_API_URL = "https://viacep.com.br/ws/14092000/json/"

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// fmt.Fprint()
	req, err := http.NewRequestWithContext(context.Background(), "GET", VIA_CEP_API_URL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a resposta da requisição: %v\n", err)
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	checkErr(err)
	defer res.Body.Close()

	var address ViaCEP
	err = json.Unmarshal(body, &address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
	}

	f, err := os.Create("address.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo %v\n", err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf(
	"Cep: %s,\nLogradouro: %s,\nComplemento: %s,\nBairro: %s,\nLocalidade: %s,\nUF: %s,\nIbge: %s,\n",
	address.Cep, address.Logradouro, address.Complemento, address.Bairro, address.Localidade, address.Uf, address.Ibge))
	checkErr(err)
	fmt.Fprintf(os.Stdout, "Teste: \n", address)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}