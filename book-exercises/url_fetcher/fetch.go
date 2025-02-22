package urlfetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// os.Args[1:]
func FetchAll(args []string) {
	for _, v := range args {
		if 	err := Fetch(v); err != nil {
			log.Fatal(err)
		}	
	}
}

func FetchAllConcurrently(args []string) {
	fetch := func(url string) (io.ReadCloser, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("http error code: %v", resp.StatusCode)
		}
		return resp.Body, nil
	}

	c := make(chan io.ReadCloser)

	for _, u := range args {
		go func(){
			data, err := fetch(u)
			if err != nil {
				log.Println(err)
			}
			c <- data
		}()
	}

	for i := 0; i < len(args); i++ {
		if d := <- c; d != nil {
			io.Copy(os.Stdout, d)
			d.Close()
		}
	}
}

func Fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("http error code: %v", resp.StatusCode)
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		return err
	}
	return nil
}
