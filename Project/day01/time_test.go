package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	timer := time.Now().UnixNano()
	t.Log(timer)
}
