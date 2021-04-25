package storage

type Repository interface {
	AccountRepository
}


type AccountRepository interface {
	GetAccount(id string) (*Account, error)
}