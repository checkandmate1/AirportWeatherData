Package to obtain wind speeds from any airport in the Aviation Weather Center database.
The function (GetWeather) needs an ICAO code for the airport. The function returns a slice of METAR structs, which can be used to extract the necessary weather data. The function also returns a boolean, which will return true if the program returns a METAR, or it will return false if nothing is returned (If "KAAC" is entered for example)
data, ok := getweather.GetWeather("KEWR")
if !ok {
    fmt.Println("Unable to get the METAR for the airport.")
}
```
Enjoy!
