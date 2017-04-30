package alphavantage

import "testing"

const (jsonResponse = `
{
    "Meta Data": {
        "1. Information": "Intraday (1min) prices and volumes",
        "2. Symbol": "MSFT",
        "3. Last Refreshed": "2017-04-28 16:00:00",
        "4. Interval": "1min",
        "5. Output Size": "Compact",
        "6. Time Zone": "US/Eastern"
    },
    "Time Series (1min)": {
        "2017-04-28 16:00:00": {
            "1. open": "68.5000",
            "2. high": "68.5200",
            "3. low": "68.4500",
            "4. close": "68.4600",
            "5. volume": "3434306"
        },
        "2017-04-28 15:59:00": {
            "1. open": "68.4900",
            "2. high": "68.5000",
            "3. low": "68.4700",
            "4. close": "68.5000",
            "5. volume": "320871"
        },
        "2017-04-28 15:58:00": {
            "1. open": "68.4550",
            "2. high": "68.4900",
            "3. low": "68.4501",
            "4. close": "68.4900",
            "5. volume": "259084"
        },
        "2017-04-28 15:57:00": {
            "1. open": "68.4400",
            "2. high": "68.4600",
            "3. low": "68.4400",
            "4. close": "68.4550",
            "5. volume": "135126"
        },
        "2017-04-28 15:56:00": {
            "1. open": "68.4300",
            "2. high": "68.4800",
            "3. low": "68.4240",
            "4. close": "68.4400",
            "5. volume": "284161"
        }
    }
}
`)

func TestParseResponse(t *testing.T) {
	body := []byte(jsonResponse)
	ts, err := parseResponse(body)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(ts.TimeSeriesEntries); l != 5 {
		t.Fatalf("Should have 5 entries")
	}
	if ts.TimeSeriesEntries[3].High != 68.46 {
		t.Fatalf("Should equal 68.46")
	}
}
