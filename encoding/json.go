package encoding

import (
	"encoding/json"
	"errors"

	"github.com/acoderup/boost/ref"
)

var (
	ErrJSONWrongValueType = errors.New("encoding JSON converts on wrong type value")
)

type JSON struct{}

func init() {
	register(NewJSON())
}

func NewJSON() *JSON {
	return new(JSON)
}

func (json JSON) String() string {
	return ref.TypeName(json)
}

func (json JSON) Style() EncodingStyleType {
	return EncodingStyleStruct
}

func (JSON) Marshal(v interface{}) ([]byte, error) {
	switch v := v.(type) {
	case []byte:
		return v, nil
	case Bytes:
		return v.Data, nil
	case *Bytes:
		return v.Data, nil
	default:
		return json.Marshal(v)
	}
}

func (JSON) Unmarshal(data []byte, v interface{}) error {
	switch v := v.(type) {
	case *Bytes:
		v.Data = data
		return nil
	default:
		return json.Unmarshal(data, v)
	}
}

func (json JSON) Reverse() Encoding {
	return json
}
