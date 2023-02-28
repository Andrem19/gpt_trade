package util

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/Andrem19/gpt_trade/variables"
)

func PrepCheckData() {
	readFile()
	for i := 0; i < len(variables.ListOHLC); i+=24 {
		var set variables.Check_data
		var sb strings.Builder
		for j := i; j < i+18; j++ {
			ohlc := variables.ListOHLC[j]
			sb.WriteString(ohlc.ToString())
		}
		set.Prompt = sb.String()
		sb.Reset()
		for j := i+18; j < i+24; j++ {
			ohlc := variables.ListOHLC[j]
			sb.WriteString(ohlc.ToString())
		}
		set.Check = sb.String()
		variables.ListCheck = append(variables.ListCheck, set)
	}
}

func readFile() {
	f, err := os.OpenFile(variables.OutputFileName, os.O_RDONLY, os.ModePerm)
    if err != nil {
        log.Fatalf("open file error: %v", err)
    }
    defer f.Close()

    sc := bufio.NewScanner(f)
	
    for sc.Scan() {
        line := sc.Text()  // GET the line string
		variables.ListOHLC = append(variables.ListOHLC, ConvLine(line))
    }
    if err := sc.Err(); err != nil {
        log.Fatalf("scan file error: %v", err)
    }
}

func SaveFile(fileName string, body string) error {

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(body); err != nil {
		panic(err)
	}
	return err
}