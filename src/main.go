package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Config struct {
	ShowRaw       bool
	ShowTemp      bool
	ShowHumid     bool
	ShowPrecipit  bool
	ShowPressure  bool
	ShowWindSpeed bool
	ShowWindDir   bool
	ShowWindGusts bool
}

func prepUrl(lat, lon float64) string {
	urlObj, err := url.Parse("https://api.open-meteo.com/v1/forecast")
	if err != nil {
		panic(err)
	}

	v := url.Values{}
	v.Add("latitude", strconv.FormatFloat(lat, 'g', 10, 64))
	v.Add("longitude", strconv.FormatFloat(lon, 'g', 10, 64))
	v.Add("current", "relative_humidity_2m")
	v.Add("current", "temperature_2m")
	v.Add("current", "precipitation")
	v.Add("current", "surface_pressure")
	v.Add("current", "wind_speed_10m")
	v.Add("current", "wind_direction_10m")
	v.Add("current", "wind_gusts_10m")
	v.Add("timezone", "auto")

	urlObj.RawQuery = v.Encode()

	return urlObj.String()
}

func GetWeatherRaw(lat, lon float64) (int, []byte) {
	res, err := http.Get(prepUrl(lat, lon))
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return res.StatusCode, body
}

type Units struct {
	Humidity      string `json:"relative_humidity_2m"`
	Temperature   string `json:"temperature_2m"`
	Precipitation string `json:"precipitation"`
	Pressure      string `json:"surface_pressure"`
	WindSpeed     string `json:"wind_speed_10m"`
	WindDirection string `json:"wind_direction_10m"`
	WindGusts     string `json:"wind_gusts_10m"`
}

type Values struct {
	Time          string  `json:"time"`
	Humidity      int     `json:"relative_humidity_2m"`
	Temperature   float64 `json:"temperature_2m"`
	Precipitation float64 `json:"precipitation"`
	Pressure      float64 `json:"surface_pressure"`
	WindSpeed     float64 `json:"wind_speed_10m"`
	WindDirection int     `json:"wind_direction_10m"`
	WindGusts     float64 `json:"wind_gusts_10m"`
}

type Response struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Elevation   float64 `json:"elevation"`
	Timezone    string  `json:"timezone_abbreviation"`
	ValuesUnits Units   `json:"current_units"`
	Values      Values  `json:"current"`
}

func Parse(c Config, raw []byte) {
	if c.ShowRaw {
		fmt.Printf("%s", raw)
		return
	}

	var data Response
	if err := json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	fmt.Printf("Time: %s\n", data.Values.Time)
	fmt.Printf("Timezone: %s\n", data.Timezone)
	fmt.Printf("Elevation: %.1f\n", data.Elevation)

	if c.ShowTemp {
		fmt.Printf("Temperature: %.1f %s\n",
			data.Values.Temperature, data.ValuesUnits.Temperature)
	}

	if c.ShowHumid {
		fmt.Printf("Humidity: %d %s\n",
			data.Values.Humidity, data.ValuesUnits.Humidity)
	}

	if c.ShowPrecipit {
		fmt.Printf("Precipitation: %.1f %s\n",
			data.Values.Precipitation, data.ValuesUnits.Precipitation)
	}

	if c.ShowPressure {
		fmt.Printf("Pressure: %.1f %s\n",
			data.Values.Pressure, data.ValuesUnits.Pressure)
	}

	if c.ShowWindSpeed {
		fmt.Printf("Wind speed: %.1f %s\n",
			data.Values.WindSpeed, data.ValuesUnits.WindSpeed)
	}

	if c.ShowWindDir {
		fmt.Printf("Wind direction: %d %s\n",
			data.Values.WindDirection, data.ValuesUnits.WindDirection)
	}

	if c.ShowWindGusts {
		fmt.Printf("Wind gusts: %.1f %s\n",
			data.Values.WindGusts, data.ValuesUnits.WindGusts)
	}
}

func main() {
	showRaw := flag.Bool("raw", false, "show http response body")
	showTemp := flag.Bool("T", false, "show temperature")
	showHumid := flag.Bool("H", false, "show humidity")
	showPrecipit := flag.Bool("P", false, "show precipitation")
	showPressure := flag.Bool("pressure", false, "show pressure")
	showWindSpeed := flag.Bool("wSpeed", false, "show wind speed")
	showWindDir := flag.Bool("wDir", false, "show wind direction")
	showWindGusts := flag.Bool("wGusts", false, "show wind gusts")

	flag.Usage = func() {
		w := flag.CommandLine.Output()

		fmt.Fprintf(w, "goweather usage: goweather <flag(s)> <lat> <lon>\n")

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s: %s\n", f.Name, f.Usage)
		})

		fmt.Fprintf(w, "\t<lat>: float value of latitude\n")
		fmt.Fprintf(w, "\t<lon>: float value of longitude\n")
	}

	flag.Parse()

	if narg := flag.NArg(); narg < 2 || narg > 2 {
		flag.Usage()
		return
	}

	config := Config{*showRaw,
		*showTemp,
		*showHumid,
		*showPrecipit,
		*showPressure,
		*showWindSpeed,
		*showWindDir,
		*showWindGusts,
	}

	lat, err := strconv.ParseFloat(flag.Arg(0), 64)
	if err != nil {
		flag.Usage()
		return
	}

	lon, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		flag.Usage()
		return
	}

	status, raw := GetWeatherRaw(lat, lon)
	fmt.Printf("Status code: %d\n", status)
	if status != 200 {
		panic("Status code returned is not 200.\n")
	}

	Parse(config, raw)
}
