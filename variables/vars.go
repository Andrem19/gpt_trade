package variables

var (
	Mode string = "Check_bin" // PrepTrain PrepNewData Check PrepTrainBinary Check_bin
	Path string = "/home/jupiter/Golang/GPT_Trade/ETHUSDT_Historical_data/training"
	OutputFileName string = "ETHUSDT_1h_check.csv"
	FineTuneModel string = "curie:ft-personal-2023-03-01-11-06-54"
	ListOHLC []OHLC
	ListTrain []Set_t
	ListCheck []Check_data
)