package codec

import (
	"bytes"
	"encoding/gob"
)

func GobEncode(obj interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	// 编码后的结果输出到buf里面
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(obj); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func GobDecode(data []byte, obj interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	return decoder.Decode(obj)
}
