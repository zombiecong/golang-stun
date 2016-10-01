package libs

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
)



type Packet struct {
	Types      uint16
	Length     uint16
	TransID    *big.Int
	Attributes []Attribute
}

func NewPacket(tid *big.Int,attrs []Attribute) (*Packet, error) {
	v := new(Packet)
	v.TransID = tid
	v.Attributes = attrs
	v.Length = 0
	return v, nil
}

func NewPacketFromBytes(packetBytes []byte) (*Packet, error) {
	if len(packetBytes) < 24 {
		return nil, errors.New("Received data length too short.")
	}

	pkgType := binary.BigEndian.Uint16(packetBytes[0:2])
	// check 00
	if pkgType >  (1 << 15 - 1 ) {
		return nil, errors.New("It is not stun message")
	}

	//check magic cookie
	magicCookieCheck := binary.BigEndian.Uint32(packetBytes[4:8]);
	if(magicCookie != magicCookieCheck){
		return nil, errors.New("no magic cookie,not supported")
	}

	pkt := new(Packet)
	pkt.Types = pkgType
	pkt.Length = binary.BigEndian.Uint16(packetBytes[2:4])
	pkt.TransID = new(big.Int).SetBytes(packetBytes[4:20])
	pkt.Attributes = make([]Attribute, 0, 10)
	for pos := uint16(20); pos < uint16(len(packetBytes)); {
		types := binary.BigEndian.Uint16(packetBytes[pos : pos+2])
		length := binary.BigEndian.Uint16(packetBytes[pos+2 : pos+4])
		if pos+4+length > uint16(len(packetBytes)) {
			return nil, errors.New("Received data format mismatch.")
		}
		value := packetBytes[pos+4 : pos+4+length]
		attribute := newAttribute(types, value)
		pkt.addAttribute(*attribute)
		pos += align(length) + 4
	}
	return pkt, nil
}

func (v *Packet) addAttribute(a Attribute) {
	v.Attributes = append(v.Attributes, a)
	v.Length += align(a.Length) + 4
}


func (v Packet) TypeToString() (typeString string)  {
	switch v.Types {
	case TypeBindingRequest:
		typeString = "BindRequest"
	case TypeAllocate:
		typeString = "Allocate"
	}
	return
}

func (v Packet) String() string {

	return fmt.Sprintf("packet : type -> %s , length -> %d , tid -> %d , length of the attr -> %d \n,  attributes -> %v",
		v.TypeToString(),v.Length,v.TransID,len(v.Attributes),v.Attributes)
}