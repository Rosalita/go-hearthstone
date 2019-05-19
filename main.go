package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hpcloud/tail"
)

func main() {

	absPath, _ := filepath.Abs("../../../../../../Program Files (x86)/Hearthstone/Logs/Power.log")

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t, err := tail.TailFile(absPath, tail.Config{
		Follow: true,
		ReOpen: true})

	for line := range t.Lines {
		fmt.Println(line.Text)
	}

}
