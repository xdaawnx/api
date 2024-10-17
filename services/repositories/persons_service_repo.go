package services_repositories

type PersonService interface {
	GetCountry(name string) (string, error)
}
