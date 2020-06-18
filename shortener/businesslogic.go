package shortener

import "errors"

var(
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)

//implements RedirectService Interfave
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
	return r.redirectRepo.Store(redirect)
}