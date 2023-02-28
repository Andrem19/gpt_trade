package util

import "log"

func Check_Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}