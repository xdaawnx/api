package models

import (
	models_repositories "api/models/repositories"
	"errors"

	"gorm.io/gorm"
)

type personQueries struct {
	DB *gorm.DB
}

func NewPersonQueries(db *gorm.DB) models_repositories.PersonRepository {
	return &personQueries{DB: db}
}

func (pq *personQueries) GetCountryByName(name string) (string, error) {
	var country string
	query := "SELECT country FROM persons WHERE name = ?"
	pq.DB.Raw(query, name).Scan(&country)
	if country == "" {
		return "", errors.New("not found")
	}
	return country, nil
}
