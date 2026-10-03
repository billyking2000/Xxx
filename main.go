package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/billyking2000/Xxx/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
