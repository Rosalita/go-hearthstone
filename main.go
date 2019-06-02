package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hpcloud/tail"
	//enum "github.com/Rosalita/go-hearthstone/hs_enum"
)

type logCall int

const (
	unknown logCall = iota
	gameState
	powerProcessor
	powerTaskList
)

func (lc logCall) String() string {
	return [...]string{"Unknown", "GameState", "PowerProcessor", "PowerTaskList"}[lc]
}

var (
	showGameState      bool
	showPowerTaskList  bool
	showPowerProcessor bool
)

func init() {

	const (
		defaultGameState      = true
		defaultPowerTaskList  = false
		defaultPowerProcessor = false
		usageGS               = "show calls to GameState"
		usagePTL              = "show calls to PowerTaskList"
		usagePP               = "show calls to PowerProcessor"
	)

	flag.BoolVar(&showGameState, "game_state", defaultGameState, usageGS)
	flag.BoolVar(&showGameState, "gs", defaultGameState, usageGS+" (shorthand)")
	flag.BoolVar(&showPowerTaskList, "power_task_list", defaultPowerTaskList, usagePTL)
	flag.BoolVar(&showPowerTaskList, "ptl", defaultPowerTaskList, usagePTL+" (shorthand)")
	flag.BoolVar(&showPowerProcessor, "power_processor", defaultPowerProcessor, usagePP)
	flag.BoolVar(&showPowerProcessor, "pp", defaultPowerProcessor, usagePP+" (shorthand)")
}

func main() {

	flag.Parse()

	absPath, _ := filepath.Abs("../../../../../../Program Files (x86)/Hearthstone/Logs/Power.log")

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	t, err := tail.TailFile(absPath, tail.Config{
		Follow: true,
		ReOpen: true,
		Poll:   true,
	})

	for line := range t.Lines {
		if parsed := parse(line.Text); parsed != "" {
			fmt.Println(parsed)
		}
	}
}

func parse(logLine string) string {
	logCall := parseLogCall(logLine)

	if logCall == gameState && showGameState {
		return logLine
	}
	if logCall == powerTaskList && showPowerTaskList {
		return logLine
	}
	if logCall == powerProcessor && showPowerProcessor {
		return logLine
	}
	return ""
}

func parseLogCall(line string) logCall {
	lineSlice := strings.Split(line, " ")

	if len(lineSlice) < 2 {
		return unknown
	}

	logCall := strings.Split(lineSlice[2], ".")

	switch logCall[0] {
	case "GameState":
		return gameState
	case "PowerProcessor":
		return powerProcessor
	case "PowerTaskList":
		return powerTaskList
	default:
		return unknown
	}
}
