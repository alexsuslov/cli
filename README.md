# Cli
Simple console cli

## cmd/example
```go
package main

import (
	"github.com/alexsuslov/cli"
	"os"
)

func main() {
	if err := cli.
		New("cli", "Example Cli").
		Action(os.Args); err != nil {
		panic(err)
	}
}
```

```.sh
> go run cmd/example/example.go
Example Cli
---
> cli help
return this help

```
### Add Action
AddAction

    - command   // string
    - action    // function func(args []string)error
    - help      // template help
    
    
```go
    New("covid-cli", "ECDC - RSS - COVID-19").
		AddAction("all", All(feedUrl), helpTemplate).
		AddAction("json", Json(feedUrl), helpJsonTemplate)
```


## cmd/example1
```bash
 go run cmd/example1/example.go json | jq

{
  "title": "ECDC - RSS - COVID-19",
  "link": "https://www.ecdc.europa.eu/en",
  "language": "en",
  "items": [
    {
      "title": "Rapid Risk Assessment: Resurgence of reported cases of COVID 19 in the EU/EEA, the UK and EU candidate and potential candidate countries",
      "description": "While decreasing trends in disease incidence are being observed in Europe overall (12% decrease in 14-day incidence of reported cases between 16 and 30 June), there is still community transmission reported in most EU/EEA countries, the UK and EU candidate and potential candidate countries. Additionally, some countries are reporting a resurgence of observed cases or large localised outbreaks.",
      "link": "https://www.ecdc.europa.eu/en/publications-data/rapid-risk-assessment-resurgence-reported-cases-covid-19",
      "published": "Thu, 02 Jul 2020 14:00:00 +0200",
      "publishedParsed": "2020-07-02T12:00:00Z",

```