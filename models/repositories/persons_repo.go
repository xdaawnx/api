package models_repositories

type PersonRepository interface {
	GetCountryByName(name string) (string, error)
}
