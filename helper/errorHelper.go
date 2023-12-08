package helper

import "log"

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
