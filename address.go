package xytsocket

import "errors"

type Address struct {
	Name    string 		`json:"name"`
	Pointer *uint 		`json:"pointer, omitempty"`
	Offset  uint		`json:"offset"`
	Next    uint		`json:"next"`
	Type    addressType     `json:"type"`

	parent  *Address
}


type Addressable interface {
	GetAddress(uint)(uint, error)
	GetAddressMap(uint)([]uint)
}

func (adr *Address) GetAddress(pos uint)(address uint, err error){
	if adr.Pointer != nil {
		address = *adr.Pointer + adr.Offset + (adr.Next * pos)
		return
	}

	if adr.parent == nil {
		err = errors.New("Invalid address tree.")
		return
	}

	_addres, err := adr.parent.GetAddress(pos)
	if err != nil{
		return
	}

	address = _addres + adr.Offset

	return
}

func (adr *Address) GetAddressMap(pos uint)(addressMap []uint) {
	_addres, err := adr.GetAddress(pos)
	if err != nil {
		return
	}

	addressMap = append(addressMap, _addres)

	return
}
