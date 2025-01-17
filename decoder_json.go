package zerolog

import (
	"bytes"
	"encoding/json"
	"io"
)

func getDecoder(buf []byte) *json.Decoder {
	decoder := json.NewDecoder(bytes.NewReader(buf))
	decoder.UseNumber()
	return decoder
}

func decodeKeyValues(buf []byte, all bool, key ...string) (map[string][]interface{}, error) {
	if !all && len(key) == 0 {
		return nil, nil
	}
	dec := getDecoder(buf)
	kv := make(map[string][]interface{})
	for {
		token, err := dec.Token()
		if err == io.EOF {
			return kv, nil
		}
		t, ok := token.(string)
		if !ok {
			continue
		}
		if !all {
			var found bool
			for _, k := range key {
				if t != k {
					continue
				}
				found = true
			}
			if !found {
				continue
			}
		}
		var val interface{}
		if err := dec.Decode(&val); err != nil {
			return nil, err
		}
		if _, ok := kv[t]; !ok {
			kv[t] = make([]interface{}, 0)
		}
		kv[t] = append(kv[t], val)
	}
}
