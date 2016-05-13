package main

import (
	"github.com/Xustyx/xytsocket"
	"fmt"
)

func main() {
	address := uint(0x0059B3D8)
	var csPlayer = new(xytsocket.ClassStruct)
	csPlayer.Name = "Player"
	csPlayer.Pointer = &address
	csPlayer.Offset = 0x00
	csPlayer.Type = xytsocket.STRUCT
	csPlayer.Next = 0x48

	var csName = new(xytsocket.ClassSpec)
	csName.Name = "Name"
	csName.Offset = 0x00
	csName.Type = xytsocket.STRING
	csName.Parent = csPlayer


	var csId = new(xytsocket.ClassSpec)
	csId.Name = "Id"
	csId.Offset = 0x38
	csId.Type = xytsocket.UINT
	csId.Parent = csPlayer

	//Player 1
	fmt.Printf("Player1 pos: %x\n", csPlayer.GetAddress(0))
	fmt.Printf("Player1 pos: %x\n", csName.GetAddress(0))
	fmt.Printf("Player1 pos: %x\n", csId.GetAddress(0))

	fmt.Println()

	//Player 2
	fmt.Printf("Player2 pos: %x\n", csPlayer.GetAddress(1))
	fmt.Printf("Player2 pos: %x\n", csName.GetAddress(1))
	fmt.Printf("Player2 pos: %x\n", csId.GetAddress(1))

}
