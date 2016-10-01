
package main

import (

	"github.com/zombiecong/golang-stun/libs"

);

func main() {
	server := libs.Server{Port:3478}
	server.Serve()

}

//
///* A Simple function to verify error */
//func CheckError(err error) {
//	if err  != nil {
//		fmt.Println("Error: " , err)
//		os.Exit(0)
//	}
//}
//
//
//
//func main() {
//
//	var trans = make([]*big.Int,0)
//	/* Lets prepare a address at any address at port 10001*/
//	ServerAddr,err := net.ResolveUDPAddr("udp",":3478")
//	CheckError(err)
//
//	/* Now listen at selected port */
//	ServerConn, err := net.ListenUDP("udp", ServerAddr)
//	CheckError(err)
//	defer ServerConn.Close()
//
//	buf := make([]byte, 256)
//
//	for {
//		n,addr,err := ServerConn.ReadFromUDP(buf)
//
//		if err != nil {
//			fmt.Println("Error: ",err)
//		}else{
//
//			p ,err := libs.NewPacketFromBytes(buf[0:n])
//
//
//			flag := false
//			for _, tid := range trans{
//				if tid.Cmp(p.TransID) == 0  {
//					flag = true
//				}
//			}
//			if  !flag {
//				trans = append(trans,p.TransID)
//
//				if err == nil {
//					switch p.Types {
//					case libs.TypeBindingRequest:
//
//						fmt.Print(p)
//						fmt.Print("binding request \n")
//						fmt.Println("Received ",string(buf[0:n]), " from ",addr)
//
//
//					case libs.TypeAllocate:
//						//fmt.Print("allocate \n")
//					}
//
//					//z := new(big.Int)
//					//z.SetBytes(p.TransID)
//
//					//fmt.Print(p)
//					//
//					//fmt.Println("Received ",string(buf[0:n]), " from ",addr)
//				}else{
//					fmt.Printf("fucking packet , %s \n",err)
//				}
//			}else{
//				fmt.Print("received , ignore. \r\n")
//			}
//
//
//
//		}
//
//
//	}
//}