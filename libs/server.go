package libs

import (
	"net"
	"fmt"
	"strconv"
	"log"
)

type Server struct {
	Port int
	Registry *Registry
	connection *net.UDPConn
}

func (s *Server) serve() {
	for {
		var buf = make([]byte, 1024)
		size, remoteAddr, err := s.connection.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		go s.handleData(remoteAddr, buf[:size])
	}
}

//NewServer conveniently creates a new server from the given port
func NewServer(port int) *Server {
	ret := new(Server)
	ret.Port = port
	ret.Registry = new(Registry)
	ret.Registry.mappings = make(map[string]*Client)
	return ret
}


func (s *Server) Serve() {
	laddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(s.Port))
	if err != nil {
		log.Fatal(err)
	}
	s.connection, err = net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	s.serve()
}


func (s *Server) handleData(raddr *net.UDPAddr, data []byte){
	msg , err := UnMarshal(data)

	if err != nil{
		return
	}

	fmt.Print("message : " , msg)

	respMsg := new(Message)
	respMsg.MessageType = TypeBindingResponse
	respMsg.TransID = msg.TransID
	respMsg.Attributes = make(map[uint16][]byte)

	addXORMappedAddress(respMsg, raddr)

	response, err := Marshal(respMsg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("response : %s \n",response)

	//send response
	_, err = s.connection.WriteToUDP(response, raddr)
	if err != nil {
		fmt.Println(err)
	}

}