package server

import (
	"log"
	"net"
	"net/rpc"
)

var RpcService RPCService

type RPCService struct {
	Address1  string
	Address2  string
	client    *rpc.Client
	Messaging *MessagingService //todo: messaging should listen to RPC
}

func (service *RPCService) Connect() {
	go service.startServer()
	go service.startClient()
}

func (service *RPCService) startClient() {
	client, err := rpc.Dial("tcp", service.Address2)
	if err != nil {
		log.Println(err)
		return
	}
	service.client = client
	err = client.Call("CommunicationStruct.InitServer", service.Messaging.RetrieveLocalAgentIds(), &struct{}{})
	FatalError(err)
	log.Println("Connected to the server")
}

func (service *RPCService) SendMessage(message Message) {
	if service.client == nil {
		log.Println("Client is nil")
		return
	}
	err := service.client.Call("CommunicationStruct.SendMessage", message, &struct{}{})
	FatalError(err)
}

type CommunicationStruct struct{}

func (service *RPCService) startServer() {
	communication := new(CommunicationStruct)
	rpc.Register(communication)
	tcpAddr, err := net.ResolveTCPAddr("tcp", service.Address1)
	FatalError(err)
	log.Println(tcpAddr)

	listener, err1 := net.ListenTCP("tcp", tcpAddr)
	FatalError(err1)

	for {
		conn, _ := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (service *CommunicationStruct) InitServer(ids []string, reply *struct{}) error {
	log.Println("Got ids")
	log.Println(ids)
	RpcService.Messaging.AddIds(ids)
	log.Println(RpcService.Messaging.Ids)
	if RpcService.client == nil {
		RpcService.startClient()
	}
	return nil
}

func (service *CommunicationStruct) SendMessage(message Message, reply *struct{}) error {
	log.Printf("Remote message from %s, to %s", message.From, message.To)
	RpcService.Messaging.SendMessage(&message)
	return nil
}
