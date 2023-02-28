package main

import (
	"fmt"
	"log"

	"github.com/Andrem19/gpt_trade/util"
	"github.com/Andrem19/gpt_trade/variables"
)

func main() {
	if variables.Mode == "PrepTrain" {
		util.ReadFileToTheModel()
		util.PrepareTrainingData()
		util.SaveTrainingData("train_set_ETHUSDT_1h.json")

	} else if variables.Mode == "PrepNewData" {
		filesList := util.GetFilesList(variables.Path)
		RangeFilesWorker(filesList)
		err := util.SaveToFile(variables.ListToSave, variables.OutputFileName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File successfuly writen")
	}
}

func RangeFilesWorker(filesList []string) {
	for _, file := range filesList {
        path := fmt.Sprintf("%s/%s", variables.Path, file)
		util.ReadAndConv(path)
		fmt.Printf("%s readed and save to list. The lenth of list is %d\n\n", file, len(variables.ListToSave))
    }
}