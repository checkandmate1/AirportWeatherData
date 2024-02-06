Package to obtain wind speeds from any airport in the Aviation Weather Center database.
The function (GetWeather) needs an ICAO code for the airport. The function returns a slice of METAR structs, which can be used to extract the necessary weather data. The function also returns an array of strings, for example it will return [ "Unable to read the JSON file for %v airport" "Unable to Unmarshal the JSON file for %v airport"].
data, ok := getweather.GetWeather("KEWR")
if !ok {
    fmt.Println("Unable to get the METAR for the airport.")
}
```
Enjoy!
