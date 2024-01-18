package cloudfunc

import (
	"log"
	"bytes"
	"encoding/ascii85"
	"encoding/json"

	usecbor "github.com/fxamacker/cbor/v2"
)

var cbor usecbor.EncMode

func init() {
		// setup CBOR encoer
		cb, err := usecbor.CanonicalEncOptions().EncMode()
		if err != nil {
			log.Fatalln(err)
		}
		cbor = cb
}

func MarshalJSON(x interface{}) ([]byte, error) {
	return json.Marshal(x)
}

func UnmarshalJSON(b []byte, dst interface{}) error {
	return json.Unmarshal(b, dst)
}

func MarshalCBOR(x interface{}) ([]byte, error) {
	return cbor.Marshal(x)
}

func UnmarshalCBOR(b []byte, dst interface{}) error {
	return usecbor.Unmarshal(b, dst)
}

func CompactSerial(x interface{}) (string, error) {
	b, err := MarshalCBOR(x)
	if err != nil {
		return "", nil
	}
	buf := bytes.NewBuffer(nil)
	enc := ascii85.NewEncoder(buf)
	enc.Write(b)
	enc.Close()
	return string(buf.Bytes()), nil
}

func ExpandSerial(serial []byte, dst interface{}) error {
	out := make([]byte, len(serial)*2)
	d, _, err := ascii85.Decode(out, serial, true)
	if err != nil {
		return err
	}
	if err := UnmarshalCBOR(out[:d], dst); err != nil {
		return err
	}
	return nil
}
