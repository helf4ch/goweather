package main

import (
	"flag"
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

	lat := flag.Arg(0)
	lon := flag.Arg(1)

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
