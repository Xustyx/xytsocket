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