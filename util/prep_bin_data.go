package util

import (
	"strings"

	"github.com/Andrem19/gpt_trade/variables"
)

func PrepTrDataBin() {
	for i := 0; i < len(variables.ListOHLC) - 25; i += 24 {
		var set variables.Set_t
		var sb strings.Builder
		for j := i; j < i+18; j++ {
			ohlc := variables.ListOHLC[j]
			sb.WriteString(ohlc.ToString())
		}
		set.Prompt = sb.String()

		set.Completion = up_or_down(i)
		variables.ListTrain = append(variables.ListTrain, set)
	}
}

func up_or_down(iter int) string {
	point_in := variables.ListOHLC[iter+17].Close

	one_proc := point_in / 100

	for i := iter+18; i < iter+24; i++ {
		if variables.ListOHLC[i].High > (point_in + one_proc) {
			return "up"
		}
		if variables.ListOHLC[i].High < (point_in - one_proc) {
			return "down"
		}
	}
	return "none"
}