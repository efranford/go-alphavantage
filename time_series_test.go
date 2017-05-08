package alphavantage

import "testing"

const (jsonResponseIntra = `
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

const (jsonResponseDaily = `
{
    "Meta Data": {
        "1. Information": "Daily Prices (open, high, low, close) and Volumes",
        "2. Symbol": "MSFT",
        "3. Last Refreshed": "2017-05-05",
        "4. Output Size": "Compact",
        "5. Time Zone": "US/Eastern"
    },
    "Time Series (Daily)": {
        "2017-05-05": {
            "1. open": "68.9000",
            "2. high": "69.0300",
            "3. low": "68.4900",
            "4. close": "69.0000",
            "5. volume": "18882800"
        },
        "2017-05-04": {
            "1. open": "69.0300",
            "2. high": "69.0800",
            "3. low": "68.6400",
            "4. close": "68.8100",
            "5. volume": "21502600"
        },
        "2017-05-03": {
            "1. open": "69.3800",
            "2. high": "69.3800",
            "3. low": "68.7100",
            "4. close": "69.0800",
            "5. volume": "28751500"
        }
    }
}
`)

const (jsonResponseWeekly = `
{
    "Meta Data": {
        "1. Information": "Weekly Prices (open, high, low, close) and Volumes",
        "2. Symbol": "MSFT",
        "3. Last Refreshed": "2017-05-05",
        "4. Time Zone": "US/Eastern"
    },
    "Weekly Time Series": {
        "2017-05-05": {
            "1. open": "68.6800",
            "2. high": "69.7100",
            "3. low": "68.4900",
            "4. close": "69.0000",
            "5. volume": "124445700"
        },
        "2017-04-28": {
            "1. open": "67.4800",
            "2. high": "69.1400",
            "3. low": "67.1000",
            "4. close": "68.4600",
            "5. volume": "158396000"
        },
        "2017-04-21": {
            "1. open": "65.0400",
            "2. high": "66.7000",
            "3. low": "64.8900",
            "4. close": "66.4000",
            "5. volume": "109098400"
        },
        "2017-04-13": {
            "1. open": "65.6100",
            "2. high": "65.8600",
            "3. low": "64.8500",
            "4. close": "64.9500",
            "5. volume": "70937000"
        },
        "2017-04-07": {
            "1. open": "65.8100",
            "2. high": "66.3500",
            "3. low": "65.1900",
            "4. close": "65.6800",
            "5. volume": "86558000"
        }
    }
}
`)

const (jsonResponseMonthly = `
{
    "Meta Data": {
        "1. Information": "Monthly Prices (open, high, low, close) and Volumes",
        "2. Symbol": "MSFT",
        "3. Last Refreshed": "2017-05-05",
        "4. Time Zone": "US/Eastern"
    },
    "Monthly Time Series": {
        "2017-05-05": {
            "1. open": "68.6800",
            "2. high": "69.7100",
            "3. low": "68.4900",
            "4. close": "69.0000",
            "5. volume": "124445700"
        },
        "2017-04-28": {
            "1. open": "65.8100",
            "2. high": "69.1400",
            "3. low": "64.8500",
            "4. close": "68.4600",
            "5. volume": "424989400"
        },
        "2017-03-31": {
            "1. open": "64.1300",
            "2. high": "66.1900",
            "3. low": "63.6200",
            "4. close": "65.8600",
            "5. volume": "478271400"
        }
    }
}
`)

func TestParseResponse(t *testing.T) {
	body := []byte(jsonResponseIntra)
	ts, err := parseResponse(body, "Time Series (1min)", "2006-01-02 15:04:05")
	if err != nil {
		t.Fatal(err)
	}
	if l := len(ts.TimeSeriesEntries); l != 5 {
		t.Fatalf("Should have 5 entries")
	}
	if ts.TimeSeriesEntries[3].High != 68.46 {
		t.Fatalf("Should equal 68.46")
	}

	body = []byte(jsonResponseDaily)
	ts, err = parseResponse(body, "Time Series (Daily)", "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}
	if l := len(ts.TimeSeriesEntries); l != 3 {
		t.Fatalf("Should have 3 entries")
	}
	if ts.TimeSeriesEntries[2].Close != 69.08 {
		t.Fatalf("Should equal 68.46")
	}

	body = []byte(jsonResponseWeekly)
	ts, err = parseResponse(body, "Weekly Time Series", "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}
	if l := len(ts.TimeSeriesEntries); l != 5 {
		t.Fatalf("Should have 5 entries")
	}
	if ts.TimeSeriesEntries[2].Close != 66.40 {
		t.Fatalf("Should equal 66.40")
	}

	body = []byte(jsonResponseMonthly)
	ts, err = parseResponse(body, "Monthly Time Series", "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}
	if l := len(ts.TimeSeriesEntries); l != 3 {
		t.Fatalf("Should have 3 entries")
	}
	if ts.TimeSeriesEntries[2].Close != 65.86 {
		t.Fatalf("Should equal 65.86")
	}
}
