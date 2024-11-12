# golog
A logger wrapper built on top of charmbracelet/lipgloss and charmbracelet/log.

## Disclaimer
I did not write any of this code, all credit goes to [charmbracelet](https://github.com/charmbracelet) for their excellent modules. I simply got tired of writing, copying, the same wrappers between projects for my use, so I decided to publish a module I could pull directly.

## Usage
```go
package main

import "github.com/aN0mad/golog/log"

var (
    Debug = true
)

func main() {

	// Setup logger
	if DEBUG {
		// Enable debug logging
		log.EnableDebug()
		log.Debug("Debug logging enabled")
	}

	// Do logging things
	log.Debug("Debug message")
	log.Info("Info message")
	log.Error("Error message")
	log.Fatal("Fatal message")
}
```

## Resources
- [https://github.com/charmbracelet/log](https://github.com/charmbracelet/log)
- [https://github.com/charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)