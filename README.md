Package to obtain wind speeds from any airport in the Aviation Weather Center database.
The function (GetWeather) needs an ICAO code for the airport. It also returns the wind direction, speed, and gust (if any.) It will look something like this:
```
direction, speed, gust := getweather.GetWeather("KEWR")
```
Enjoy!
