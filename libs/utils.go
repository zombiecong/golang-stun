package libs

func padding(bytes []byte) []byte  {
	length := uint16(len(bytes))
	return append(bytes,make([]byte,align(length)-length)...)

}

func align(n uint16) uint16  {
	return (n + 3) & 0xfffc
}