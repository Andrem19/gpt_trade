package variables

var (
	Mode string = "Check" // PrepTrain PrepNewData Check
	Path string = "/home/jupiter/Golang/GPT_Trade/ETHUSDT_Historical_data/check"
	OutputFileName string = "ETHUSDT_1h_check.csv"
	FineTuneModel string = "curie:ft-personal-2023-02-28-18-01-43"
	ListOHLC []OHLC
	ListTrain []Set_t
	ListCheck []Check_data
)