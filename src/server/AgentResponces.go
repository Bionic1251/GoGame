package server

import (
	"log"
)

const EVERYONE_NEARBY = "ALL"

const SUGGEST_COFFEE = 0
const WANTS_COFFEE_ANSWER = 1
const DRINKING_COFFEE_ANSWER = 2

var messageDescriptions []func(a *Agent, from string)

func init() {
	messageDescriptions = []func(a *Agent, from string){respondCoffee, processCoffeeResponce, processCoffeeAction}
}

func respondCoffee(agent *Agent, from string) {
	if agent.WantsCoffee() {
		log.Printf("%s agreed to drink coffee with %s", agent.Id, from)
		message := agent.createMessage(WANTS_COFFEE_ANSWER, from)
		messaging.SendMessage(&message)
	} else {
		log.Printf("%s ignored %s", agent.Id, from)
	}
}

func processCoffeeResponce(agent *Agent, from string) {
	if agent.WantsCoffee() {
		log.Printf("%s is drinking coffee with %s", agent.Id, from)
		agent.DrinkCoffee()
		message := agent.createMessage(DRINKING_COFFEE_ANSWER, from)
		messaging.SendMessage(&message)
	} else {
		log.Printf("%s refused to drink coffee with %s", agent.Id, from)
	}
}

func processCoffeeAction(agent *Agent, from string) {
	if agent.WantsCoffee() {
		log.Printf("%s is drinking coffee with %s as well", agent.Id, from)
		agent.DrinkCoffee()
	} else {
		log.Printf("%s refused to drink coffee with %s, because they don't want to", agent.Id, from)
	}
}
