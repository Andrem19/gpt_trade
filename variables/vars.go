package variables

var (
	Mode string = "PrepTrainBinary" // PrepTrain PrepNewData Check PrepTrainBinary
	Path string = "/home/jupiter/Golang/GPT_Trade/ETHUSDT_Historical_data/training"
	OutputFileName string = "ETHUSDT_1h_training.csv"
	FineTuneModel string = "text-davinci-003"
	ListOHLC []OHLC
	ListTrain []Set_t
	ListCheck []Check_data
)