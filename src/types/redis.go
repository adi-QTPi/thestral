package types

type Storage interface {
	Add(*ProxyRoute) error
	GetAll() (map[string]*ProxyRoute, error)
}
