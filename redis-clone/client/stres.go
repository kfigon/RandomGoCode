package client

import (
	"fmt"
	"math/rand"
	"redis-clone/command"
	"redis-clone/util"
	"sync"
	"time"
)

func RunStres(port int, threads int) {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			SendData(port, []byte(randomCommand()))
		}()	
	}

	wg.Wait()
	took := time.Since(start)
	average := float64(threads)/float64(took.Milliseconds())
	fmt.Println("all done, took", took, "average [ms]", average)
}

func randomCommand() string {
	switch rand.Intn(5) {
	case 0: return command.BuildPingCommand()
	case 1: return command.BuildEchoCommand(util.RandStringRunes(500))
	case 2:	return command.BuildSetCommand(util.RandStringRunes(200))
	case 3: return command.BuildGetCommand(util.RandStringRunes(200))
	case 4: return command.BuildDeleteCommand(util.RandStringRunes(200))
	default: return command.BuildPingCommand()
	}
}