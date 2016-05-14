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

type addressStruct struct {
	Address
	attributes []*Address
}

func AddressStruct()(css *addressStruct){
	css = new(addressStruct)
	css.Type = STRUCT
	return
}

func (adrSt *addressStruct) GetAddressMap(pos uint)(addressMap []uint){
	for _, element := range adrSt.attributes {
		for _,subelement := range element.GetAddressMap(pos){
			addressMap = append(addressMap, subelement)
		}
	}

	return
}

func (adrSt *addressStruct) AppendAddress(adr *Address){
	adr.parent = &adrSt.Address
	adrSt.attributes = append(adrSt.attributes, adr)

	return
}
