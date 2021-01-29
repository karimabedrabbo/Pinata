package apputils

import (
	"bytes"
	"encoding/json"
)

func Marshal(i interface{}) ([]byte, error) {
	return json.Marshal(&i)
}

func Unmarshal(b []byte) (interface{}, error) {
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()

	var i interface{}
	err := d.Decode(&i)
	if err != nil {
		return nil, err
	}
	return i, nil
}


func Remarshal(obj interface{}) (interface{}, error) {
	var err error
	var objBytes []byte
	if objBytes, err = Marshal(obj); err != nil {
		return nil, err
	}

	var newObj interface{}
	if newObj, err = Unmarshal(objBytes); err != nil {
		return nil, err
	}

	return newObj, nil
}

