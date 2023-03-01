package util

import (
	"strings"

	"github.com/Andrem19/gpt_trade/variables"
)

func PrepCheckBinData() {
	readFile()
	for i := 0; i < len(variables.ListOHLC)-25; i+=24 {
		var set variables.Check_data
		var sb strings.Builder
		
		for j := i; j < i+18; j++ {
			ohlc := variables.ListOHLC[j]
			sb.WriteString(ohlc.ToString())
		}
		set.Prompt = sb.String()
		
		set.Check = up_or_down(i)
		variables.ListCheck = append(variables.ListCheck, set)
	}
}