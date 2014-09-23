package server

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"util"
)

/**
* To run the program on two servers it's necessary to run two instances
* with different ports
* Port conflict of port1 == port2 send log.fatal
const port1 = 4001//port of the current server
const port2 = 4000//port of the server to connect to
*/

var messaging MessagingService

type InitService struct {
	Agents      map[string]*Agent
	Containers  []Container
	AgentCount  int
	AddressBook util.IPConns
}

func (initService *InitService) Run() {
	initService.AddressBook = util.IpConnections
	log.Println("address book: \n", initService.AddressBook)
	AgentNumber = 0
	messaging = MessagingService{}
	initService.AgentCount = 100
	initService.CreateLocations()
	initService.CreateAgents()
	initService.initRPC()
	go initService.RunAgents()
}

//Gate of another ip could have been used here
//Ip + Port is used for reliability reason
func (initService *InitService) getLocalGate() string { 
	return initService.AddressBook.Ips[0].Ip + ":" + initService.AddressBook.Ips[0].Port
}

func (initService *InitService) initRPC() {
	address1 := initService.AddressBook.Ips[1].Gate
	address2 := initService.AddressBook.Ips[0].Gate
	RpcService = RPCService{Address1: address1, Address2: address2}
	RpcService.Messaging = &messaging
	RpcService.Connect()
}

func (service *InitService) CreateLocations() {
	log.Println("Create locations")
	cooridnate := Coordinate{0, 0}
	location := Location{cooridnate, 5}
	cafeteria := Cafeteria{location}
	service.Containers = []Container{cafeteria}
}

func (service *InitService) CreateAgents() {
	log.Println("Create agents")
	gate := service.getLocalGate()
	agents := make(map[string]*Agent)
	for i := 0; i < service.AgentCount; i++ {
		person, error := CreatePerson(0, 0)
		if error != nil {
			log.Fatal(error)
		}
		agent := createAgent(person, gate)
		agents[agent.Id] = &agent
	}
	service.Agents = agents
	messaging.GetAgents(agents)
	messaging.Containers = service.Containers
}

func (service *InitService) RunAgents() {
	var channel = make(chan struct{})
	for _, agent := range service.Agents {
		log.Println(agent)
		go agent.Run(channel)
	}
	for i := 0; i < service.AgentCount; i++ {
		<-channel
	}
}

func CreatePerson(x float64, y float64) (p Person, err error) {
	p.FirstName = getRandomString()
	p.LastName = getRandomString()
	p.Coordinate = Coordinate{x, y}
	p.Age = 0
	return p, nil
}

const STRLEN = 10

func getRandomString() string {
	b := make([]byte, STRLEN)
	rand.Read(b)
	en := base64.StdEncoding
	d := make([]byte, en.EncodedLen(len(b)))
	en.Encode(d, b)
	return string(d)
}

func CreateLocation(radius float64, coord Coordinate) (loc Location, err error) {
	loc.Radius = radius
	loc.Coordinate = coord
	return loc, nil
}

func SetPerson(firstName string, lastName string, loc Coordinate) (p Person, err error) {
	p.FirstName = firstName
	p.LastName = lastName
	p.Coordinate = loc
	return p, nil
}

/**
* Test function to test
 */
func SetLocation(radius float64, coord Coordinate) (loc Location, err error) {
	if (coord.X < 0) || (coord.Y < 0) {
		return Location{}, errors.New("coordinate cannot be negative")
	}
	loc.Radius = radius
	loc.Coordinate = coord
	return loc, nil
}
