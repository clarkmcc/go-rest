package marshaller

import "encoding/json"

type jsonMarshaller struct{}

func NewJsonMarshaller() Marshaller {
	return &jsonMarshaller{}
}

func (m *jsonMarshaller) Marshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (m *jsonMarshaller) Unmarshal(b []byte, obj interface{}) error {
	return json.Unmarshal(b, obj)
}
