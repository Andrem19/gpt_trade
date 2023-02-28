package util

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Andrem19/gpt_trade/variables"
)

func GetFilesList(path string) []string {
	files, err := ioutil.ReadDir(path)
    if err != nil {
        log.Fatal(err)
    }

	var filesList []string

	for _, file := range files {
        filesList = append(filesList, file.Name())
    }
	// for _, file := range files {
    //     fmt.Println(file.Name(), file.IsDir())
    // }
	return filesList
   
}

func ReadAndConv(path string) {
    f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
    if err != nil {
        log.Fatalf("open file error: %v", err)
    }
    defer f.Close()

    sc := bufio.NewScanner(f)
	
    for sc.Scan() {
        line := sc.Text()  // GET the line string
		variables.ListToSave = append(variables.ListToSave, ConvLine(line))
    }
    if err := sc.Err(); err != nil {
        log.Fatalf("scan file error: %v", err)
    }
}

func ConvLine(line string) variables.OHLC {
	tmp := strings.Split(line, ",")
	var temp variables.OHLC

	temp.Time = tmp[0]
	temp.Open, _ = strconv.ParseFloat(tmp[1], 32)
	temp.High, _ = strconv.ParseFloat(tmp[2], 32)
	temp.Low, _ = strconv.ParseFloat(tmp[3], 32)
	temp.Close, _ = strconv.ParseFloat(tmp[4], 32)
	temp.Volume, _ = strconv.ParseFloat(tmp[5], 32)

	return temp
}

func SaveToFile(list []variables.OHLC, fileName string) error {
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

	for _, ohlc := range list {
		if _, err = f.WriteString(ohlc.ToString()); err != nil {
			panic(err)
		}
	}
	return err
}
