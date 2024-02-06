package getweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type CloudData struct {
	Cover string `json:"cover"`
	Base  int    `json:"base"`
}

type MetarData struct {
	MetarID     int         `json:"metar_id"`
	ICAOID      string      `json:"icaoId"`
	ReceiptTime string      `json:"receiptTime"`
	ObsTime     int         `json:"obsTime"`
	ReportTime  string      `json:"reportTime"`
	Temp        float64     `json:"temp"`
	Dewp        float64     `json:"dewp"`
	Wdir        int         `json:"wdir"`
	Wspd        int         `json:"wspd"`
	Wgst        int         `json:"wgst"`
	Altimiter   float32     `json:"altim"`
	Clouds      []CloudData `json:"clouds"`
}




func GetWeather(icao string) ([]MetarData, []string) {
	data, ok := gatherData(icao)

	return data, ok
}

func gatherData(icao string) ([]MetarData, []string) {
	url := fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%v&format=json", icao)
	var erro []string
	
	resp, err := http.Get(url)
	if err != nil {
		erro = append(erro, fmt.Sprintf("Unable to make the http.Get request for %v", icao))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		erro = append(erro, "Unable to read the JSON file for %v airport", icao)
	}
	var metarData []MetarData
	err = json.Unmarshal(body, &metarData)
	if err != nil {
		erro = append(erro, "Unable to unmarshall JSON for %v airport", icao)
	}

	if len(metarData) <= 0 {
		erro = append(erro, fmt.Sprintf("No METAR availble"))
	}
	return metarData, erro

}
