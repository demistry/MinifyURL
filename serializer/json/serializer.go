package json

import (
	"MinifyURL/shortener"
	"encoding/json"
	"errors"
)

type Redirect struct{}

func newSerializer(input *shor) *RedirectSerializer{
	//write code to create a new serializer
	return nil
}

func (r *Redirect) Decode(input []byte)(*shortener.Redirect,error){
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect), err != nil{
		return nil, errors.New("Could not decode redirect object from bytes")
	}
	return redirect,nil
}

func (r *Redirect) Encode(input *Redirect) ([]byte, error){
	encodedData, err := json.Marshal(input)
	if err != nil{
		return nil,err.New("Could not encode data to bytes...")
	}
	return encodedData,nil
}