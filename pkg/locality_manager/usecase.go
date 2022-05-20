package locality_manager

import (
	"strings"
)

//Service  interface
type service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) UseCase {
	return &service{
		repo: r,
	}
}

//CreateLocality Create locality
func (s *service) CreateLocality(code LocalityCode, statisticalCode LocalityStatisticalCode, name LocalityName, status LocalityStatus) (LocalityCode, error) {
	l, err := NewLocality(code, statisticalCode, name, status)
	if err != nil {
		return l.Code, err
	}
	return s.repo.Create(l)
}

//GetLocality Get locality
func (s *service) GetLocality(code LocalityCode) (*Locality, error) {
	return s.repo.Get(code)
}

//SearchLocalities Search localities
func (s *service) SearchLocalities(query string) ([]*Locality, error) {
	return s.repo.Search(strings.ToLower(query))
}

//ListLocalities List localities
func (s *service) ListLocalities() ([]*Locality, error) {
	return s.repo.List()
}

//DeleteLocality Delete locality
func (s *service) DeleteLocality(code LocalityCode) error {
	l, err := s.GetLocality(code)
	if l == nil {
		return ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(code)
}

//UpdateLocality Update locality
func (s *service) UpdateLocality(l *Locality) error {
	err := l.Validate()
	if err != nil {
		return ErrInvalidEntity
	}
	return s.repo.Update(l)
}
