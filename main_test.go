package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printTheString(t *testing.T) {

	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)
	go printTheString("Saini", &wg)
	wg.Wait()

	_ = w.Close()

	resultStore, _  := io.ReadAll(r)

	output := string(resultStore)

	os.Stdout = stdOut

	if !strings.Contains(output, "Saini") {
		t.Errorf("Expected to find Saini but it is not there")
	}
}
