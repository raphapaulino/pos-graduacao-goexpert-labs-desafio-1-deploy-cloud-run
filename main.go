package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

var VIA_CEP_API_BASE_URL = "https://viacep.com.br/ws"

var WEATHER_API_BASE_URL = "https://api.weatherapi.com/v1"

type ViaCEP struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Localidade  string `json:"localidade,omitempty"`
	Uf          string `json:"uf,omitempty"`
	Ibge        string `json:"ibge,omitempty"`
	Gia         string `json:"gia,omitempty"`
	Ddd         string `json:"ddd,omitempty"`
	Siafi       string `json:"siafi,omitempty"`
	Erro        bool   `json:"erro,omitempty"`
}

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %s", err)
	// }

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/{cep}", func(r chi.Router) {
		r.Use(checkCepMiddleware)
		r.Get("/", handleGetTemperatureByCEP)
	})

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", r)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCepMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cep := chi.URLParam(r, "cep")

		if cep == "" || len(cep) == 0 {
			http.Error(w, "CEP is required", http.StatusBadRequest)
			return
		}

		if !isValidZipcode(cep) {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleGetTemperatureByCEP(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	address, err := getAddressFromViaCEP(cep)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weather, err := getWeather(address.Localidade)
	if err != nil {
		http.Error(w, "can not find weather", http.StatusNotFound)
		return
	}

	temperature := TemperatureResponse{
		TempC: weather.Current.TempC,
		TempF: celsiusToFahrenheit(weather.Current.TempC),
		TempK: celsiusToKelvin(weather.Current.TempC),
	}
	json.NewEncoder(w).Encode(temperature)
}

func getAddressFromViaCEP(cep string) (*ViaCEP, error) {
	url := fmt.Sprintf(VIA_CEP_API_BASE_URL+"/%s/json/", cep)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var address ViaCEP
	err = json.NewDecoder(resp.Body).Decode(&address)
	if address.Erro {
		fmt.Println(err)
		return nil, fmt.Errorf("zipcode not found")
	}
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func getWeather(city string) (*WeatherResponse, error) {
	fmt.Println("Cidade:", city)
	cleanedCity := replaceSpecialCharacters(city) // to deal with strings, for example, like "Ribeirão Preto" or "São Simão"
	cityEncoded := url.QueryEscape(cleanedCity)
	// weatherApiKey := os.Getenv("WEATHER_API_KEY")
	weatherApiKey := "cbca91bf0fb24a7c97835630240205"

	url := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no", WEATHER_API_BASE_URL, weatherApiKey, cityEncoded)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func replaceSpecialCharacters(str string) string {
	replacements := map[string]string{
		`[áàâãªä]`: "a",
		`[ÁÀÂÃÄ]`:  "A",
		`[ÍÌÎÏ]`:   "I",
		`[íìîï]`:   "i",
		`[éèêë]`:   "e",
		`[ÉÈÊË]`:   "E",
		`[óòôõºö]`: "o",
		`[ÓÒÔÕÖ]`:  "O",
		`[úùûü]`:   "u",
		`[ÚÙÛÜ]`:   "U",
		`ç`:        "c",
		`Ç`:        "C",
		`ñ`:        "n",
		`Ñ`:        "N",
		`–`:        "-",
		`[’‘‹›‚]`:  " ",
		`[“”«»„]`:  " ",
	}

	for pattern, replacement := range replacements {
		re, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println(err)
			continue
		}
		str = re.ReplaceAllString(str, replacement)
	}
	return str
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}

func isValidZipcode(zipcode string) bool {
	if len(zipcode) != 8 {
		return false
	}

	for _, char := range zipcode {
		if _, err := strconv.Atoi(string(char)); err != nil {
			return false
		}
	}

	return true
}
