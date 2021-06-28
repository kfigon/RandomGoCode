package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)


type config struct {
	name string
	num int
	trigger bool
}

func main() {
	cfg := parseConfig()
	start := time.Now()
	handleApp(cfg)
	fmt.Println("all done, bye, took", time.Since(start))
}

func parseConfig() config {
	cfg := config{}
	flag.BoolVar(&cfg.trigger, "trig", false, "switch to enable that big feature")
	flag.StringVar(&cfg.name, "name", "", "actual value of that thing")
	flag.IntVar(&cfg.num, "num", 15, "how many things to exec")

	flag.Parse()
	
	return cfg
}

func handleApp(cfg config) {
	if cfg.name == "" {
		fmt.Println("\"name\" param is required")
		return
	}

	if cfg.trigger {
		fmt.Println("running new algorithm")
		fmt.Printf("Print %v, %v times\n", cfg.name, cfg.num)
		newAlgo(os.Stdout, cfg)
		return
	}

	fmt.Println("running old algorithm")
	oldAlg(os.Stdout, cfg)
}

func newAlgo(w io.Writer, cfg config) {
	for i := 0; i < cfg.num; i++ {
		fmt.Fprintln(w, cfg.name)
	}
}

func oldAlg(w io.Writer, cfg config) {
	for i := 0; i < cfg.num; i++ {
		fmt.Fprint(w, cfg.name)
	}
	fmt.Fprintln(w)
}