package getweather

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"net"
	"net/http"
	"time"

	"github.com/caffix/cloudflare-roundtripper/cfrt"
)
var direction, speed, gust2 int
type CloudData struct {
	Cover string `json:"cover"`
	Base  int    `json:"base"`
}

type MetarData struct {
	MetarID     int        `json:"metar_id"`
	ICAOID      string     `json:"icaoId"`
	ReceiptTime string     `json:"receiptTime"`
	ObsTime     int        `json:"obsTime"`
	ReportTime  string     `json:"reportTime"`
	Temp        float64    `json:"temp"`
	Dewp        float64    `json:"dewp"`
	Wdir        int        `json:"wdir"`
	Wspd        int        `json:"wspd"`
	Wgst        *int       `json:"wgst"`
	Clouds      []CloudData `json:"clouds"`
}


func GetWeather(icao string) (int, int, int) {
	initColly(icao)
	return direction, speed, gust2
}


func initColly(icao string) (){
	url := fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%v&format=json", icao)
	var err error

	// Setup your client however you need it. This is simply an example
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   15 * time.Second,
				KeepAlive: 15 * time.Second,
				DualStack: true,
			}).DialContext,
		},
	}
	
	client.Transport, err = cfrt.New(client.Transport)
	if err != nil {
		fmt.Println("Error creating RoundTripper:", err)
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", resp.StatusCode)
		return
	}

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	c.OnResponse(func(r *colly.Response) {
		
		var metarData []MetarData
		err := json.Unmarshal(r.Body, &metarData)
		if err != nil {
			fmt.Printf("Error Unmarshaling Weather JSON: %v\n", err)
			return
		}

		printWeatherData(metarData)
		
		
	})

	err = c.Visit(url)
	if err != nil {
		fmt.Println("Error visiting the URL:", err)
	}
}

func printWeatherData(data []MetarData) {
	for _, metar := range data {
		

		direction = metar.Wdir
		speed = metar.Wspd
		gust := metar.Wgst
		
		if gust == nil {
			gust2 = 0 
		}
		
		
	}
	
	
}