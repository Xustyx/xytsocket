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
