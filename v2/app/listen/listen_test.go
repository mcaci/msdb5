package listen

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestWithAINames(t *testing.T) { // add -race to go test for running this test
	ch := make(chan string)
	go WithAINames(ch)
	for c := range ch {
		t.Log(c)
	}
}

func TestWithRand(t *testing.T) { // add -race to go test for running this test
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	go WithRand(ctx, ch, func() int { return 0 })
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	for c := range ch {
		t.Log(c)
	}
}

func start(tr *TickerRand, f func()) {
	ctx, cancel := context.WithCancel(tr.Context)
	defer cancel()
	for i := 0; i < 5; i++ {
		select {
		case <-tr.tick.C:
			tr.fCh <- f
		case <-ctx.Done():
			close(tr.fCh)
		}
	}
}

func TestWithRand2(t *testing.T) { // add -race to go test for running this test
	tr := &TickerRand{Context: context.Background(), tick: time.NewTicker(10 * time.Millisecond), fCh: make(chan func())}
	go WithRand2(tr)
	start(tr, func() {
		log.Println("hello")
	})
}

func start3(cfc *ContextFChan, f func()) {
	// ctx, cancel := context.WithCancel(cfc.Context)
	// defer cancel()
	// for i := 0; i < 5; i++ {

	// 	select {
	// 	case <-tr.tick.C:
	// 		cfc.fCh <- f
	// 	case <-ctx.Done():
	// 		close(cfc.fCh)
	// 	}
	// }

	reader := bufio.NewReader(os.Stdin)

	log.Print("Enter your name: ")

	name, _ := reader.ReadString('\n')
	strings.NewReader("hello\n").WriteTo(log.Writer())
	log.Println(strings.NewReader("hello\n"))
	// reader.WriteTo(log.Writer())

	log.Printf("Hello %s\n", name)
}

func TestWithRand3(t *testing.T) { // add -race to go test for running this test
	cfc := &ContextFChan{Context: context.Background(), fCh: make(chan func())}
	// go WithRand2(cfc)
	start3(cfc, func() {
		log.Println("hello")
	})
	t.Fail()
}
