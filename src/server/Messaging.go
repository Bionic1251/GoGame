package server

import (
	"errors"
	"fmt"
	"log"
)

type Message struct {
	Id     int64
	TypeId int
	From   string
	To     string
}

type MessagingService struct {
	Agents     map[string]*Agent //local agents are stored here
	Ids        []string         //all ids are stired here
	Containers []Container
}

func (messagingService *MessagingService) GetAgents(agentMap map[string]*Agent) {
	messagingService.Agents = agentMap
	messagingService.Ids = messagingService.RetrieveLocalAgentIds()
}

func (messagingService *MessagingService) RetrieveLocalAgentIds() []string {
	ids := make([]string, len(messagingService.Agents))
	i := 0
	for key, _ := range messagingService.Agents {
		ids[i] = key
		i++
	}
	return ids
}

func (messagingService *MessagingService) AddIds(ids []string) {
	for _, v := range ids {
		messagingService.Ids = append(messagingService.Ids, v) //todo: correct it
	}
}

func (messagingService *MessagingService) SendMessage(message *Message) {
	if message.To == EVERYONE_NEARBY {
		for _, agentId := range messagingService.Ids {
			if agentId == message.From {
				continue
			}
			agent, error := messagingService.GetAgentById(agentId)
			if error != nil {
				messagingService.remoteCall(message, agentId)
				continue
			}
			agent.ReceiveMessage(message)
		}
	} else {
		agent, error := messagingService.GetAgentById(message.To)
		if error != nil {
			messagingService.remoteCall(message, message.To)
			return
		}
		agent.ReceiveMessage(message)
	}
}

func (messagingService *MessagingService) remoteCall(message *Message, agentId string) {
asd
	log.Printf("Remote call, destination %s ", agentId)
	newMessage := Message{}
	newMessage.From = message.From
	newMessage.To = agentId
	lastMessageId++
	newMessage.Id = lastMessageId
	newMessage.TypeId = message.TypeId
	RpcService.SendMessage(newMessage)
}

func (service *MessagingService) GetAgentById(id string) (agent *Agent, err error) {
	value, ok := service.Agents[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Agent with id %s haven't been found", id))
	}
	return value, nil
}

func (service *MessagingService) GetContainer(c Coordinate) Container {
	for _, container := range service.Containers {
		if container.Contains(c) {
			return container
		}
	}
	return nil
}
