package httpMagic

type BackOffHandler interface {
	Execute(apiRequest func(req HttpRequest) (string, error), rec HttpRequest) (string, error)
}
