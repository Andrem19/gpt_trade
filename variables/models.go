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

type Check_data struct {
	Prompt string `json:"prompt"`
	Check string `json:"check"`
}

type Check_data_to_save struct {
	Resp string `json:"resp"`
	Check string `json:"check"`
}

func (resp *Check_data_to_save) ToString() string {
	return fmt.Sprintf("%s=>\n%s;\n\n", resp.Resp, resp.Check)
}

func (ohlc *OHLC) ToString() string {
	return fmt.Sprintf("%.2f,%.2f,%.2f,%.2f,%.2f\n", ohlc.Open, ohlc.High, ohlc.Low, ohlc.Close, ohlc.Volume)
}