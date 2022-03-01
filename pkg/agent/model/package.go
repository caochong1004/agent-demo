package model

import "fmt"

type Packet struct {
	Key     string `json:"key,omitempty"`
	Type    string `json:"type,omitempty"`
	Payload string `json:"payload,omitempty"`
}

func (c *Packet) String() string {
	return fmt.Sprintf("{key: %s, type: %s}  %s", c.Key, c.Type, c.Payload)
}
