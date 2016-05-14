//The MIT License (MIT)
//
//Copyright (c) 2016 Xustyx
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

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
