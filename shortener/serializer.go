package shortener

type RedirectSerializer interface{
	Decode(inputData []byte) (*Redirect, error)
	Encode(inputData *Redirect) ([]byte, error)
}