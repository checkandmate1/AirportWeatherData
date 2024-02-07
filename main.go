package getweather

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var c *colly.Collector

type cloudData struct {
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
	RawMETAR    string      `json:"rawOb"`
	Clouds      []cloudData `json:"clouds"`
}



func GetWeather(icao string) ([]MetarData, []string) {
	data, ok := gatherData(icao)
	if len(data) <= 0 {
		ok = append(ok, "No METAR Availible")
		fmt.Println(len(data))
	}
	return data, ok
}

func gatherData(icao string) ([]MetarData, []string) {
	var erro []string
	var data []MetarData
	url := fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%v&format=json", icao)
	c = colly.NewCollector(
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 4,
		RandomDelay: 2 * time.Second,
	})

	c.OnResponse(func(r *colly.Response) {

		err := json.Unmarshal(r.Body, &data)
		if err != nil {
			erro = append(erro, "Error unmarshaling JSON")
		}
	})

	err := c.Visit(url)
	if err != nil {
		erro = append(erro, "Error visiting the URL")
	}
	c.Wait()
	return data, erro

}
