package proto

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/nsyszr/admincenter/pkg/util/typeconv"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (dec *Decoder) Decode() (Envelope, error) {
	var msg []interface{}
	if err := json.NewDecoder(dec.r).Decode(&msg); err != nil {
		return nil, err
	}

	if len(msg) == 0 {
		return nil, fmt.Errorf("empty message")
	}

	// Resolve the message type
	v, err := typeconv.AnyToFloat64(msg[0])
	if err != nil {
		return nil, err
	}
	msgType := int64ToMessageType(int64(v))

	raw := RawMessage{
		messageType: msgType,
		message:     msg,
	}

	return raw, nil
}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}
