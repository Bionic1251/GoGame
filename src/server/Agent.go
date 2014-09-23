package server

import (
	"fmt"
	"log"
	"time"
	"strconv"
)

const maxAge = 100000

var AgentNumber int
var lastMessageId = int64(-1)

type Agent struct {
	Id string
	Person
	drinkingCoffee bool
}

func createAgent(person Person, gate string) (agent Agent) {
	AgentNumber++
	id := gate + "." + strconv.Itoa(AgentNumber)
	return Agent{id, person, false}
}

func (agent *Agent) String() string {
	return fmt.Sprintf("Agent: %s, person: %v ", agent.Id, agent.Person)
}

func (a *Agent) Run(channel chan struct{}) {
	for ; a.Person.Age < maxAge; a.Person.Age++ {
		a.live()
		time.Sleep(1000000000)
	}
	channel <- struct{}{}
}

func (agent *Agent) live() {
	if !agent.drinkingCoffee {
		log.Printf("Agent %s wants to drink coffee", agent.Id)
		message := agent.createMessage(SUGGEST_COFFEE, EVERYONE_NEARBY)
		messaging.SendMessage(&message)
	} else {
		log.Printf("Agent %s is busy", agent.Id)
		agent.drinkingCoffee = false
	}
}

func (agent *Agent) createMessage(messageType int, to string) Message {
	message := Message{}
	message.From = agent.Id
	message.To = to
	lastMessageId++
	message.Id = lastMessageId
	message.TypeId = messageType
	return message
}

func (agent *Agent) DrinkCoffee() {
	log.Printf("Agent %s is drinking coffee", agent.Id)
	agent.drinkingCoffee = true
}

func (agent *Agent) WantsCoffee() bool {
	return !agent.drinkingCoffee
}

func (agent *Agent) ReceiveMessage(message *Message) {
	log.Printf("Agent %s received message from %s", agent.Id, message.From)
	messageDescriptions[message.TypeId](agent, message.From)
}
