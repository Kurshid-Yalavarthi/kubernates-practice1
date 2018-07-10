package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type WeatherByLatLongAndDateResponse struct {
	Currently struct {
		Time                int     `json:"time"`
		Summary             string  `json:"summary"`
		Icon                string  `json:"icon"`
		PrecipIntensity     float64 `json:"precipIntensity"`
		PrecipProbability   float64 `json:"precipProbability"`
		Temperature         float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		DewPoint            float64 `json:"dewPoint"`
		Humidity            float64 `json:"humidity"`
		WindSpeed           float64 `json:"windSpeed"`
		WindGust            float64 `json:"windGust"`
		WindBearing         float64 `json:"windBearing"`
		Visibility          float64 `json:"visibility"`
		CloudCover          float64 `json:"cloudCover"`
		Pressure            float64 `json:"pressure"`
		Ozone               float64 `json:"ozone"`
		UvIndex             int     `json:"uvIndex"`
	} `json:"currently"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int     `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     float64 `json:"precipIntensity"`
			PrecipProbability   float64 `json:"precipProbability"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			WindSpeed           float64 `json:"windSpeed"`
			WindGust            float64 `json:"windGust"`
			WindBearing         float64 `json:"windBearing"`
			Visibility          float64 `json:"visibility"`
			CloudCover          float64 `json:"cloudCover"`
			Pressure            float64 `json:"pressure"`
			Ozone               float64 `json:"ozone"`
			UvIndex             int     `json:"uvIndex"`
		} `json:"data"`
	} `json:"hourly"`
}

type HTMLError struct {
	Err  string
	Code int
}

func WeatherByLatLongAndDate(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path[len("/weatherByLatLongAndDate/"):]
	splitPath := strings.Split(path, "/")
	lat := splitPath[0]
	long := splitPath[1]
	startDate := splitPath[2]

	i, err := strconv.ParseInt(startDate, 10, 64)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		errResult := &HTMLError{
			Err:  err.Error(),
			Code: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(errResult)
		log.Println("WARN: fatal error: %s", err)
	}
	tm := time.Unix(i, 0)
	var forecast []WeatherByLatLongAndDateResponse
	w.Header().Set("Content-Type", "application/json")

	for i := 0; i < 8; i++ {

		temp := tm.AddDate(0, 0, -i)

		resp, err := http.Get("https://api.darksky.net/forecast/095d9044727f5bf447237bb93505f765/" + lat + "," + long + "," + strconv.Itoa(int(temp.Unix())))
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			errResult := &HTMLError{
				Err:  err.Error(),
				Code: http.StatusBadRequest,
			}
			json.NewEncoder(w).Encode(errResult)
			log.Println("WARN: fatal error: %s", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			errResult := &HTMLError{
				Err:  err.Error(),
				Code: http.StatusBadRequest,
			}
			json.NewEncoder(w).Encode(errResult)
			log.Println("WARN: fatal error: %s", err)
		}

		var day WeatherByLatLongAndDateResponse
		if err := json.Unmarshal(body, &day); err != nil {

			w.WriteHeader(http.StatusBadRequest)
			errResult := &HTMLError{
				Err:  err.Error(),
				Code: http.StatusBadRequest,
			}
			json.NewEncoder(w).Encode(errResult)
			log.Println("WARN: fatal error: %s", err)
		}

		forecast = append(forecast, day)
	}

	json.NewEncoder(w).Encode(forecast)
}
