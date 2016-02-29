package main

import "github.com/peterh/liner"
import "fmt"
import "io"

const (
	promptDefault = "ws> "
)

type wsLiner struct {
	*liner.State
	buffer string
	depth  int
}

func newWsLiner() *wsLiner {
	l := liner.NewLiner()
	l.SetCtrlCAborts(true)
	return &wsLiner{State: l}
}

func (wl *wsLiner) Prompt() (string, error) {
	l, e := wl.State.Prompt(promptDefault)

	if e == io.EOF {
		if wl.buffer != "" {
			// cancel line continuation
			wl.Accepted()
			fmt.Println()
			e = nil
		}
	} else if e == liner.ErrPromptAborted {
		e = nil
		if wl.buffer != "" {
			wl.Accepted()
		} else {
			fmt.Println("(^D to quit)")
		}
	} else if e == nil {
		if wl.buffer != "" {
			wl.buffer = wl.buffer + "\n" + l
		} else {
			wl.buffer = l
		}
	}

	return wl.buffer, e
}

func (wl *wsLiner) Accepted() {
	wl.State.AppendHistory(wl.buffer)
	wl.buffer = ""
}
