package packet

import (
	"bytes"
)

type Packet interface {
	Decode([]byte) error
	Encode() ([]byte, error)
}

type Submit struct {
	ID string
	Payload []byte
}


func (s *Submit) Decode(pktBody []byte) error {
	s.ID = string(pktBody[:8])
	s.Payload = pktBody[8:]
	return nil
}

func (s *Submit) Encode() ([]byte, error) {
	return bytes.Join([][]byte{
		[]byte(s.ID[:8]), s.Payload,
	}, nil), nil
}

type SubmitAck struct {
	ID string
	Result uint8
}

func (s *SubmitAck) Decode(pktBody []byte) error {
	s.ID = string(pktBody[:8])
	s.Result = uint8(pktBody[8])
	return nil
}

func (s *SubmitAck) Encode() ([]byte, error) {
	return bytes.Join([][]byte{
		[]byte(s.ID[:8]), []byte{s.Result},
	}, nil), nil
}
