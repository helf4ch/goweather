# Goweather

It's a simple go app for fetching weather. The data comes from open-meteo.com. No API key required.

# Build

For build, clone the repo. Then:
```
$ cd goweather
$ make
```

You can run executable in buid directory.

For install, run:
```
$ make install
```

This will copy executable into `~/.local/bin/`.

# Usage

For use, select all needed flags and specify latitude and longitude.

Help:
```
goweather usage: goweather <flag(s)> <lat> <lon>
	-H: show humidity
	-P: show precipitation
	-T: show temperature
	-pressure: show pressure
	-raw: show http response body
	-wDir: show wind direction
	-wGusts: show wind gusts
	-wSpeed: show wind speed
	<lat>: float value of latitude
	<lon>: float value of longitude
```

# Example

Parsed weather in Moscow:
```
$ goweather -H -P -T -pressure -wSpeed -wDir -wGusts 55.7512 37.6184
Status code: 200
Time: 2025-12-25T06:45
Timezone: GMT+3
Elevation: 151.0
Temperature: -4.1 °C
Humidity: 93 %
Precipitation: 0.0 mm
Pressure: 1006.0 hPa
Wind speed: 7.6 km/h
Wind direction: 275 °
Wind gusts: 20.5 km/h
```

Or, in raw response body (json):
```
$ goweather -raw 55.7512 37.6184
{"latitude":55.75,"longitude":37.625,"generationtime_ms":0.07319450378417969,"utc_offset_seconds":10800,"timezone":"Europe/Moscow","timezone_abbreviation":"GMT+3","elevation":151.0,"current_units":{"time":"iso8601","interval":"seconds","relative_humidity_2m":"%","temperature_2m":"°C","precipitation":"mm","surface_pressure":"hPa","wind_speed_10m":"km/h","wind_direction_10m":"°","wind_gusts_10m":"km/h"},"current":{"time":"2025-12-25T06:45","interval":900,"relative_humidity_2m":93,"temperature_2m":-4.1,"precipitation":0.00,"surface_pressure":1006.0,"wind_speed_10m":7.6,"wind_direction_10m":275,"wind_gusts_10m":20.5}}
```
