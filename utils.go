package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getParentIdFromMessage(msg string) int64 {
	var parentId = int64(0)
	splitMsg := strings.Split(msg, " ")
	fmt.Printf("%#v", splitMsg)
	if len(splitMsg) > 0 {
		parentId, _ = strconv.ParseInt(splitMsg[1], 10, 0)
	}
	return parentId
}

func getDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
