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

package xytsocket_test

import (
	"github.com/Xustyx/xytsocket"
	"testing"
)

//Test struct to make the test.
type testAddress struct {
	adr      *xytsocket.Address
	expected []uint
}

var adrLstPlayerList = xytsocket.AddressList()
var adrStPlayer = xytsocket.AddressStruct()
var adrName = new(xytsocket.Address)
var adrId = new(xytsocket.Address)

var testGetAddress = []testAddress{
	{&adrLstPlayerList.Address,	[]uint{0x0059B3D8,0x0059B420,0x0059B468}},	//Next
	{&adrStPlayer.Address,		[]uint{0x00,0x00,0x00}},			//None
	{adrName,			[]uint{0x10,0x20,0x30}},			//Offset + Next
	{adrId,				[]uint{0x00,0x00,0x00}},			//Offset
}

var testGetAddressSelf = []testAddress{
	{&adrLstPlayerList.Address,	[]uint{0x11223344,0x1122338C,0x112233D4}},	//Next
	{&adrStPlayer.Address,		[]uint{0x11223345,0x11223345,0x11223345}},	//None
	{adrName,			[]uint{0x11223346,0x11223366,0x11223376}},	//Offset + Next
	{adrId,				[]uint{0x11223347,0x1122337F,0x1122337F}}, 	//Offset
}


//Test: Initializes the objects.
func TestInit(t *testing.T){
	var address = uint(0x0059B3D8)
	var zero = uint(0x00)

	adrLstPlayerList.Name = "Player-list"
	adrLstPlayerList.Pointer = &address
	adrLstPlayerList.Offset = 0x00
	adrLstPlayerList.Next = 0x48
	adrLstPlayerList.Total = 2

	adrStPlayer.Name = "Player"
	adrStPlayer.Offset = 0x00

	adrName.Name = "Name"
	adrName.Pointer = &zero
	adrName.Offset = 0x10
	adrName.Next = 0x10
	adrName.Type = xytsocket.STRING

	adrId.Name = "Id"
	adrId.Offset = 0x38
	adrId.Type = xytsocket.UINT
}

//Test: For each value in expected test if getAddress of expected pos match.
func TestGetAddress(t *testing.T){
	t.Logf("\nChecking TestGetAddress")
	for _,elem := range testGetAddress {
		for i, expected := range elem.expected {
			oValue, _ := elem.adr.GetAddress(uint(i))
			eValue := expected

			t.Logf("Checking %s spec(%s) --> (%x == %x)", elem.adr.Name, elem.adr.Type, oValue, eValue)

			if oValue != eValue {
				t.Errorf("Object value %x is not expected %x", oValue, eValue)
			}
		}
	}
}

//Test: Set the pointer with the expected value [0] and test the address of the next 2.
func TestGetAddressSelf(t *testing.T){
	t.Logf("\nChecking TestGetAddressSelf")

	for _,elem := range testGetAddressSelf {
		//Set pointer with expected value [0]
		epValue := elem.expected[0]
		elem.adr.Pointer = &epValue

		for i, expected := range elem.expected[1:] {
			eValue := expected
			oValue, _ := elem.adr.GetAddress(uint(i+1))

			t.Logf("Checking %s spec(%s) --> (%x == %x)", elem.adr.Name, elem.adr.Type, oValue, eValue)

			if oValue != eValue {
				t.Errorf("Object value %x is not expected %x", oValue, eValue)
			}
		}
	}
}

