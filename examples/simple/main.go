package main

// A simple program that counts down from 5 and then exits.

import (
	"fmt"
	"log"
	"tea"
	"time"
)

type model int

type tickMsg struct{}

func main() {
	err := tea.NewProgram(model(5), update, view, []tea.Sub{tick}).Start()
	if err != nil {
		log.Fatal(err)
	}
}

func update(msg tea.Msg, mdl tea.Model) (tea.Model, tea.Cmd) {
	m, _ := mdl.(model)

	switch msg.(type) {
	case tickMsg:
		m -= 1
		if m <= 0 {
			return m, tea.Quit
		}
	}
	return m, nil
}

func view(mdl tea.Model) string {
	m, _ := mdl.(model)
	return fmt.Sprintf("Hi. This program will exit in %d seconds...", m)
}

func tick(_ tea.Model) tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}
