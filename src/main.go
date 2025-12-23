package main

import (
	"fmt"
	"io"
	"net/http"
)

var (
	url = "https://api.open-meteo.com/v1/forecast?" +
		"latitude=50.2796&" +
		"longitude=127.5405&" +
		"current=relative_humidity_2m," +
		"temperature_2m," +
		"precipitation," +
		"wind_speed_10m," +
		"wind_direction_10m," +
		"wind_gusts_10m," +
		"surface_pressure&" +
		"timezone=auto"
)
func main() {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.StatusCode)
	fmt.Printf("%s", body)
}
