package main

import (
	"github.com/Xustyx/xytsocket"
	"fmt"
)

func main() {
	var address = uint(0x0059B3D8)

	var adrLstPlayerList = xytsocket.AddressList()
	adrLstPlayerList.Name = "Player-list"
	adrLstPlayerList.Pointer = &address
	adrLstPlayerList.Offset = 0x00
	adrLstPlayerList.Next = 0x48
	adrLstPlayerList.Total = 2

	var adrStPlayer = xytsocket.AddressStruct()
	adrStPlayer.Name = "Player"
	adrStPlayer.Offset = 0x00

	var adrName = new(xytsocket.Address)
	adrName.Name = "Name"
	adrName.Offset = 0x00
	adrName.Type = xytsocket.STRING

	var adrId = new(xytsocket.Address)
	adrId.Name = "Id"
	adrId.Offset = 0x38
	adrId.Type = xytsocket.UINT

	adrLstPlayerList.SetStruct(adrStPlayer)
	adrStPlayer.AppendAddress(adrName)
	adrStPlayer.AppendAddress(adrId)

	addresses := adrLstPlayerList.GetAddressMap(0)

	for i, e := range addresses {
		fmt.Printf("Address %d: %x\n", i ,e)
	}
}
