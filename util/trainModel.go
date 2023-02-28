package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Andrem19/gpt_trade/variables"
)

func ReadFileToTheModel() {
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

func PrepareTrainingData() {

	for i := 0; i < len(variables.ListOHLC); i+=24 {
		var set variables.Set_t
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
		set.Completion = sb.String()
		variables.ListTrain = append(variables.ListTrain, set)
	}

}

func SaveTrainingData(fileName string) error {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		fmt.Println("File not exist")
	} else {
		e := os.Remove(fileName)
    	if e != nil {
        	log.Fatal(e)
    	}
		fmt.Println("File removed")
	}

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()



	for _, set := range variables.ListTrain {
		bt, _ := json.Marshal(set)
		str := fmt.Sprintf("%s\n", bt)
		if _, err = f.WriteString(str); err != nil {
			panic(err)
		}
	}
	return err
}