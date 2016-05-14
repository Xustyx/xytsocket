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

package xytsocket

type addressList struct {
	Address
	Total uint
	Struct *addressStruct
}

func AddressList()(css *addressList){
	css = new(addressList)
	css.Type = LIST
	return
}

func (adrLst *addressList) GetAddressMap(pos uint)(addressMap []uint){
	for i:= uint(0) ; i < adrLst.Total; i++ {
		for _,subelement := range adrLst.Struct.GetAddressMap(i){
			addressMap = append(addressMap, subelement)
		}
	}

	return
}

func (adrLst *addressList) SetStruct(adrSt *addressStruct){
	adrSt.parent = &adrLst.Address
	adrLst.Struct = adrSt

	return
}