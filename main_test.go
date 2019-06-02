package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestLineType(t *testing.T) {

	tests := []struct {
		line     string
		logCall logCall
	}{
		{"D 21:34:16.8743346 GameState.DebugPrintPower() - BLOCK_END", gameState},
		{"D 21:34:18.2823406 PowerProcessor.PrepareHistoryForCurrentTaskList() - m_currentTaskList=9", powerProcessor},
		{"D 21:34:18.2818457 PowerTaskList.DebugDump() - ID=9 ParentID=0 PreviousID=0 TaskCount=1", powerTaskList},
		{"foo bar baz", unknown},
	}

	for _, test := range tests {
		result := parseLogCall(test.line)
		assert.Equal(t, test.logCall.String(), result.String())
	}
}
