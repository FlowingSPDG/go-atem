package atem

import (
	"fmt"
)

type AtemCommand struct {
	Name   string
	Body   []byte
	Header []byte
}

func NewCommand(Name string, Body []byte) *AtemCommand {
	return &AtemCommand{Name: Name, Body: Body}
}

func parseCommand(msg []byte) *AtemCommand {
	return &AtemCommand{Name: string(msg[4:8]), Body: msg[8:]}
}

func (ac *AtemCommand) length() uint16 {
	return uint16(len(ac.Body) + 8)
}

func (ac *AtemCommand) string() string {
	return fmt.Sprintf("Command:\t[%s]\t%d", ac.Name, ac.Body)
}

func (ac *AtemCommand) toBytes() []byte {
	var result []byte

	// Set length
	result = append(result, []byte{uint8(ac.length() >> 8), uint8(ac.length() & 0xFF)}...)

	// Set header
	result = append(result, []byte{0, 0}...)

	// Set cmd
	result = append(result, []byte(ac.Name)...)

	// Add body
	result = append(result, ac.Body...)

	return result
}
