package fatal

import "encoding/json"

func UnlessMarshal(value any) []byte {
	data, err := json.Marshal(value)
	OnError(err, "marshalling: ")
	return data
}

func UnlessUnmarshal(data []byte, value any) {
	err := json.Unmarshal(data, value)
	OnError(err, "unmarshalling: ")
}
