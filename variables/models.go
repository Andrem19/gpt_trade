package variables

import "fmt"

type OHLC struct {
	Time string `json:"time"`
	Open float64`json:"open"`
	High float64 `json:"high"`
	Low float64 `json:"low"`
	Close float64 `json:"close"`
	Volume float64 `json:"volume"`
}

type Set_t struct {
	Prompt string `json:"prompt"`
	Completion string `json:"completion"`
}

func (ohlc *OHLC) ToString() string {
	return fmt.Sprintf("%s,%.2f,%.2f,%.2f,%.2f,%.2f\n", ohlc.Time, ohlc.Open, ohlc.High, ohlc.Low, ohlc.Close, ohlc.Volume)
}