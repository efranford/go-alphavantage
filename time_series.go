package alphavantage

import (
	"time"
	"encoding/json"
	"log"
	"sort"
	"fmt"
)

type TimeSeriesService service

type MetaData struct {
	Information	string	`json:"1. Information"`
	Symbol		string	`json:"2. Symbol"`
	LastRefreshed	string	`json:"3. Last Refreshed"`
	Interval	string	`json:"4. Interval"`
	OutputSize	string	`json:"5. Output Size"`
	TimeZone	string	`json:"6. Time Zone"`
}

type TimeSeriesEntry struct {
	Time 	time.Time
	Open 	float32		`json:"1. open,string"`
	High 	float32		`json:"2. high,string"`
	Low 	float32		`json:"3. low,string"`
	Close 	float32		`json:"4. close,string"`
	Volume	int		`json:"5. volume,string"`
}

type TimeSeries struct {
	MetaData 		*MetaData
	TimeSeriesEntries	[]TimeSeriesEntry
}

type jsonMap map[string]TimeSeriesEntry

type response struct {
	MetaData	*MetaData		`json:"Meta Data"`
	TimeSeries 	*json.RawMessage
}

func parseResponse(rawRes []byte) (*TimeSeries, error) {

	jsonRaw := map[string]json.RawMessage{}
	json.Unmarshal(rawRes, &jsonRaw)

	res := response{}

	if err := json.Unmarshal(rawRes, &res); err != nil {
		log.Fatal(err)
		return nil, err
	}

	var jsonMap jsonMap

	title := fmt.Sprintf("Time Series (%s)", res.MetaData.Interval)

	if err := json.Unmarshal(jsonRaw[title], &jsonMap); err != nil {
		log.Fatal(err)
		return nil, err
	}

	result := TimeSeries{
		MetaData: res.MetaData,
		TimeSeriesEntries: []TimeSeriesEntry{},
	}

	for key, _ := range jsonMap {
		loc, _ := time.LoadLocation(res.MetaData.TimeZone)
		timeP, _ := time.ParseInLocation("2006-01-02 15:04:05", key, loc)
		tmp := jsonMap[key]
		tmp.Time = timeP
		jsonMap[key] = tmp
		result.TimeSeriesEntries = append(result.TimeSeriesEntries, jsonMap[key])
	}

	sort.Slice(result.TimeSeriesEntries, func(i, j int) bool {
		return result.TimeSeriesEntries[i].Time.After(result.TimeSeriesEntries[j].Time)
	})

	return &result, nil
}

func (s *TimeSeriesService) IntraDay(symbol string, interval string) (*TimeSeries, error) {

	req, _ := s.client.NewGetRequest("query")

	q := req.URL.Query()
	q.Add("function", "TIME_SERIES_INTRADAY")
	q.Add("symbol", symbol)
	q.Add("interval", interval)
	req.URL.RawQuery = q.Encode()

	body, err := s.client.Do(req)

	if err != nil {
		return nil, err
	}

	return parseResponse(body)

}