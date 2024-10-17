package services

import (
	models_repositories "api/models/repositories"
)

type PersonService struct {
	Repo models_repositories.PersonRepository
}

func NewPersonService(repo models_repositories.PersonRepository) *PersonService {
	return &PersonService{Repo: repo}
}

func (s *PersonService) GetCountry(name string) (string, error) {
	return s.Repo.GetCountryByName(name)
}
