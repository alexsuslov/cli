package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexsuslov/cli"
	"github.com/mmcdole/gofeed"
	"io"
	"net/http"
	"os"
)

func main() {
	feedUrl := "https://www.ecdc.europa.eu/en/taxonomy/term/2942/feed"
	C := cli.
		New("covid-cli", "ECDC - RSS - COVID-19").
		AddAction("all", All(feedUrl), helpTemplate).
		AddAction("json", Json(feedUrl), helpJsonTemplate)

	if err := C.Action(os.Args); err != nil {
		panic(err)
	}
}

func All(url string) func([]string) error {
	return func([]string) error {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			return fmt.Errorf("status code %v", res.StatusCode)
		}
		defer res.Body.Close()
		_, err = io.Copy(os.Stdout, res.Body)
		return nil
	}
}

func Json(url string) func([]string) error {
	return func([]string) error {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(url)
		if err != nil {
			return err
		}
		encoder := json.NewEncoder(os.Stdout)
		return encoder.Encode(feed)
	}
}

var helpTemplate = `
> {{.name}} all
ECDC - RSS - COVID-19
`

var helpJsonTemplate = `
> {{.name}} json
JSON ECDC - RSS - COVID-19
`
