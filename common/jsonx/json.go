package jsonx

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/modern-go/reflect2"
	"io"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	extra.RegisterFuzzyDecoders()
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Get(data []byte, path ...interface{}) jsoniter.Any {
	return json.Get(data, path)
}

func MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}

func NewEncoder(writer io.Writer) *jsoniter.Encoder {
	return json.NewEncoder(writer)
}

func NewDecoder(reader io.Reader) *jsoniter.Decoder {
	return json.NewDecoder(reader)
}

func Valid(data []byte) bool {
	return json.Valid(data)
}

func RegisterExtension(extension jsoniter.Extension) {
	json.RegisterExtension(extension)
}

func DecoderOf(typ reflect2.Type) jsoniter.ValDecoder {
	return json.DecoderOf(typ)
}

func EncoderOf(typ reflect2.Type) jsoniter.ValEncoder {
	return json.EncoderOf(typ)
}

func BorrowStream(writer io.Writer) *jsoniter.Stream {
	return json.BorrowStream(writer)
}

func ReturnStream(stream *jsoniter.Stream) {
	json.ReturnStream(stream)
}

func BorrowIterator(data []byte) *jsoniter.Iterator {
	return json.BorrowIterator(data)
}

func ReturnIterator(iter *jsoniter.Iterator) {
	json.ReturnIterator(iter)
}
