package jsonserializer

import (
	"MinifyURL/shortener"
	"encoding/json"
	"errors"
)

type Redirect struct{}

func newJsonSerializer() *shortener.RedirectSerializer{
	//write code to create a new serializer
	// return nil
	return nil
}

func (r *Redirect) Decode(input []byte)(*shortener.Redirect,error){
	redirect := &shortener.Redirect{}
	err := json.Unmarshal(input, redirect)
	if err != nil{
		return nil, errors.New("Could not decode redirect object from bytes with error")
	}
	return redirect,nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error){
	encodedData, err := json.Marshal(input)
	if err != nil{
		return nil,err
	}
	return encodedData,nil
}