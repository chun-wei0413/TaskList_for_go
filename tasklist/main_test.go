package main

import (
	"bufio"
	"fmt"
	"io"
	"sync"
	"testing"
)

//const prompt = "> " // 與 main.go 中的 prompt 保持一致

type scenarioTester struct {
	*testing.T
	inWriter   io.Writer
	outReader  io.Reader
	outScanner *bufio.Scanner
}

func TestRun(t *testing.T) {
	// 設置輸入輸出管道
	inPR, inPW := io.Pipe()
	defer inPR.Close()
	outPR, outPW := io.Pipe()
	defer outPR.Close()
	tester := &scenarioTester{
		T:          t,
		inWriter:   inPW,
		outReader:  outPR,
		outScanner: bufio.NewScanner(outPR),
	}

	// 在 goroutine 中運行 TaskManager
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		NewTaskManager(inPR, outPW).Run()
		outPW.Close()
		wg.Done()
	}()

	// 測試場景
	fmt.Println("(show empty)")
	tester.execute("show")
	tester.readLines([]string{""}) // 空列表時只輸出一個空行

	fmt.Println("(add first project)")
	tester.execute("add project ezKanban")
	fmt.Println("(add tasks to ezKanban)")
	tester.execute("add task ezKanban refactor")
	tester.execute("add task ezKanban feature addition")

	fmt.Println("(show tasks)")
	tester.execute("show")
	tester.readLines([]string{
		"ezKanban",
		"    [ ] 1: refactor",
		"    [ ] 2: feature addition",
		"",
	})

	fmt.Println("(add second project)")
	tester.execute("add project dcTrack")
	fmt.Println("(add tasks to dcTrack)")
	tester.execute("add task dcTrack testing")

	fmt.Println("(check tasks)")
	tester.execute("check 2") // feature addition
	tester.execute("check 3") // testing

	fmt.Println("(show completed tasks)")
	tester.execute("show")
	tester.readLines([]string{
		"ezKanban",
		"    [ ] 1: refactor",
		"    [X] 2: feature addition",
		"",
		"dcTrack",
		"    [X] 3: testing",
		"",
	})

	fmt.Println("(quit)")
	tester.execute("quit")

	// 確保程式已退出
	inPW.Close()
	wg.Wait()
}

// execute 執行一個命令，先讀取提示符，再發送命令
func (t *scenarioTester) execute(cmd string) {
	p := make([]byte, len(prompt))
	_, err := t.outReader.Read(p)
	if err != nil {
		t.Errorf("Prompt could not be read: %v", err)
		return
	}
	if string(p) != prompt {
		t.Errorf("Invalid prompt, expected \"%s\", got \"%s\"", prompt, string(p))
		return
	}
	fmt.Fprintln(t.inWriter, cmd)
}

// readLines 讀取輸出並與預期行比對
func (t *scenarioTester) readLines(lines []string) {
	for _, expected := range lines {
		if !t.outScanner.Scan() {
			t.Errorf("Expected \"%s\", no input found", expected)
			break
		}
		actual := t.outScanner.Text()
		if actual != expected {
			t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
		}
	}
	if err := t.outScanner.Err(); err != nil {
		t.Fatalf("Could not read input: %v", err)
	}
}

// discardLines 丟棄指定數量的行，用於清空緩衝區
func (t *scenarioTester) discardLines(n int) {
	for i := 0; i < n; i++ {
		if !t.outScanner.Scan() {
			t.Error("Expected a line, no input found")
			break
		}
	}
	if err := t.outScanner.Err(); err != nil {
		t.Fatalf("Could not read input: %v", err)
	}
}
