package libs

import (
	"errors"
	"encoding/binary"
	"math/big"
	"net"
)

var (
	ErrInvalidRequest = errors.New("Invalid STUN request")
)

type Message struct {
	MessageType uint16
	MessageLength uint16
	Magic uint32
	TransID *big.Int
	Attributes    map[uint16][]byte
}

func UnMarshal(data []byte) (* Message, error) {
	length := len(data)
	if length < 24 {
		return nil, ErrInvalidRequest
	}

	pkgType := binary.BigEndian.Uint16(data[0:2])
	// check 00
	if pkgType >  (1 << 15 - 1 ) {
		return nil, errors.New("It is not stun message")
	}

	//check magic cookie
	magicCookieCheck := binary.BigEndian.Uint32(data[4:8]);
	if(magicCookie != magicCookieCheck){
		return nil, errors.New("no magic cookie,not supported")
	}

	msg := new(Message)
	//parse the header
	msg.MessageType = binary.BigEndian.Uint16(data[0:2])
	msg.MessageLength = binary.BigEndian.Uint16(data[2:4])
	msg.TransID = new(big.Int).SetBytes(data[4:20])

	msg.Attributes = make(map[uint16][]byte)
	i := 20
	for i < length {
		attrType := binary.BigEndian.Uint16(data[i : i+2])
		attrLength := binary.BigEndian.Uint16(data[i+2 : i+4])
		i += 4 + int(attrLength)
		msg.Attributes[attrType] = data[i-int(attrLength) : i]
		if pad := int(attrLength) % 4; pad > 0 {
			i += 4 - pad
		}
	}
	//recover here to catch any index errors
	if recover() != nil {
		return nil, ErrInvalidRequest
	}

	return msg, nil
}

func Marshal(m *Message) ([]byte, error)  {
	result := make([]byte, 20)
	//first do the header
	binary.BigEndian.PutUint16(result[:2], m.MessageType)
	result = append(result[4:8], m.TransID.Bytes()...)

	//now we do the attributes
	if m.Attributes != nil {
		i := 20
		for t, v := range m.Attributes {
			length := len(v)
			binary.BigEndian.PutUint16(result[i:i+2], t)
			binary.BigEndian.PutUint16(result[i+2:i+4], uint16(len(v)))
			result = append(result[:i+4], v...)
			i += 4 + length
			//if we need to pad, do so
			if pad := length % 4; pad > 0 {
				result = append(result, make([]byte, 4-pad)...)
				i += 4 - pad
			}
		}
		binary.BigEndian.PutUint16(result[2:4], uint16(i-20))
	}
	return result, nil
}

func addXORMappedAddress(m *Message, raddr *net.UDPAddr) {
	addr := raddr.IP.To4()
	port := uint16(raddr.Port)
	xbytes := xorAddress(port, addr)
	m.Attributes[AttributeXorMappedAddress] = append([]byte{0,attributeFamilyIPv4},xbytes...)
}

func xorAddress(port uint16, addr []byte) []byte {
	xport := make([]byte ,2)
	xaddr := make([]byte ,4)
	binary.BigEndian.PutUint16(xport, port^uint16(magicCookie>>16))
	binary.BigEndian.PutUint32(xaddr, binary.BigEndian.Uint32(addr)^magicCookie)
	return append(xport, xaddr...)
}