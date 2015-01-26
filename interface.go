package db

// Database interfaces

type List interface {
	Add(value string) error
	GetAll() ([]string, error)
	GetLast() (string, error)
	GetLastN(n int) ([]string, error)
	Remove() error
	Clear() error
}

type Set interface {
	Add(value string) error
	Has(value string) (bool, error)
	GetAll() ([]string, error)
	Del(value string) error
	Remove() error
	Clear() error
}

type HashMap interface {
	Set(elementid, property, value string) error
	Get(elementid, property string) (string, error)
	Has(elementid, property string) (bool, error)
	Exists(elementid string) (bool, error)
	GetAll() ([]string, error)
	DelKey(elementid, property string) error
	Del(elementid string) error
	Remove() error
	Clear() error
}
