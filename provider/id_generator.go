package provider

import (
	"fmt"
	. "github.com/matoous/go-nanoid"
)

type IDGenerator interface {
	UserId() string
}

type NanoIDGenerator struct{
}

func (n NanoIDGenerator) UserId() string{
	return getId("user")
}

func getId(idPrefix string) string {
	id, _ := Nanoid()
	return fmt.Sprintf("%s_%s", idPrefix, id)
}