package xytsocket

type SpecType int

//Enum of spec types.
const (
	UINT SpecType = iota
	INT
	BYTE
	STRING
	STRUCT
	LIST
)

//String representation of spec types.
var specTypes = [...]string{
	"uint",
	"int",
	"byte",
	"string",
	"struct",
	"list",
}

//Get the string value from enum value.
func (specType SpecType) String() string {
	return specTypes[specType]
}

type ClassSpec struct {
	Name string 		`json:"name"`
	Pointer *uint 		`json:"pointer, omitempty"`
	Offset uint		`json:"offset"`
	Next uint		`json:"next"`
	Type SpecType		`json:"type"`

	parent *ClassSpec
}

type ClassStruct struct {
	ClassSpec
	attributes []*ClassSpec
}

type ClassList struct {
	ClassSpec
	Total uint
	Struct *ClassStruct
}

type Addressable interface {
	GetAddress(uint)(uint)
}

func (cs *ClassSpec) GetAddress(pos uint)(address uint){
	if cs.Pointer != nil {
		address = *cs.Pointer + cs.Offset + (cs.Next * pos)
	} else {
		address = cs.parent.GetAddress(pos) + cs.Offset
	}

	return
}

func (cs *ClassSpec) GetAddressMap(pos uint)(addressMap []uint){

	switch cs.Type {
		case LIST:
			csl := *cs.(&ClassList)
			for i:= 0 ; i < csl.Total; i++ {
				addressMap = append(addressMap,csl.Struct.GetAddressMap(i))
			}
		case STRUCT:
			css := *cs.(&ClassStruct)
			for _, element := range css.attributes {
				addressMap = append(addressMap, element.GetAddressMap(pos))
			}
		default:
			addressMap = append(addressMap, cs.GetAddress(pos))

	}
	return
}

func (cs *ClassStruct) addChild(child *ClassSpec){
	child.parent = cs;
	cs.attributes = append(cs.attributes, child)

	return
}
