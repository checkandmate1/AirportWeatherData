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

func GetWeather(icao string) ([]MetarData, bool) {
	data, ok := gatherData(icao)

	return data, ok
}

func gatherData(icao string) ([]MetarData, bool) {
	url := fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%v&format=json", icao)
	var err error
	var sucsess bool
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting URL ", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading JSON")
	}
	var metarData []MetarData
	err = json.Unmarshal(body, &metarData)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
	}

	if len(metarData) > 0 {
		sucsess = true
	} else {
		sucsess = false
	}
	return metarData, sucsess

}
