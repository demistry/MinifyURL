package shortener

import (
	"errors"
	"time"
	"github.com/teris-io/shortid"
)

var(
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)

//implements RedirectService Interface
type redirectService struct{
	redirectRepo RedirectRepository
}
func NewRedirectService(redirectRepository RedirectRepository) RedirectService{
	return &redirectService{redirectRepo: redirectRepository}
}

func (r *redirectService) Find(code string) (*Redirect, error){
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error{
	//You can validate input url here before storing

	redirect.Code = shortid.MustGenerate() //generate a unique code for the link
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}