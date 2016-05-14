package xytsocket

type addressType int

//Enum of address types.
const (
	UINT addressType = iota
	INT
	BYTE
	STRING
	STRUCT
	LIST
)

//String representation of address types.
var addressTypes = [...]string{
	"uint",
	"int",
	"byte",
	"string",
	"struct",
	"list",
}

//Get the string value from enum value.
func (addressType addressType) String() string {
	return addressTypes[addressType]
}
