package main

import (
	"fmt"
	"time"

	"github.com/Andrem19/gpt_trade/util"
	"github.com/Andrem19/gpt_trade/variables"
)

func main() {
	conf, err := util.LoadConfig(".")
	util.Check_Err(err)

	if variables.Mode == "PrepTrain" {
		util.ReadFileToTheModel()
		
		util.PrepareTrainingData()
		util.SaveTrainingData("train_set_ETHUSDT_1h.json")

	} else if variables.Mode == "PrepTrainBinary" {
		util.ReadFileToTheModel()
		fmt.Println("Lenth: ",len(variables.ListOHLC))

		util.PrepTrDataBin()
		util.SaveTrainingData("bin_train_set_ETHUSDT_1h.json")

		fmt.Println("File successfuly writen")
	} else if variables.Mode == "PrepNewData" {
		filesList := util.GetFilesList(variables.Path)
		RangeFilesWorker(filesList)
		err := util.SaveToFile(variables.ListOHLC, variables.OutputFileName)
		util.Check_Err(err)

		fmt.Println("File successfuly writen")
	} else if variables.Mode == "Check" {
		util.PrepCheckData()
		for i := 0; i < 10; i++ {
			resp, err := util.AskQuestion(variables.ListCheck[i].Prompt, conf.OPENAI_API_KEY)
			time.Sleep(8 * time.Second)
			util.Check_Err(err)
			body := variables.Check_data_to_save{
				Resp: resp,
				Check: variables.ListCheck[i].Check,
			}
			util.SaveFile("reults.txt", body.ToString())
		}
	}
}

func RangeFilesWorker(filesList []string) {
	for _, file := range filesList {
        path := fmt.Sprintf("%s/%s", variables.Path, file)
		util.ReadAndConv(path)
		fmt.Printf("%s readed and save to list. The lenth of list is %d\n\n", file, len(variables.ListOHLC))
    }
}

