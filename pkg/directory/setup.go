package directory

import (
	"log"
)

func Setup() {
	err := EnsureDir("./logs")
	if err != nil {
		log.Fatal(err)
	}
	err = EnsureDir("./data/storage")
	if err != nil {
		log.Fatal(err)
	}
}
