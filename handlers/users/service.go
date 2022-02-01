package users

import (
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Create(item *models.User) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.Users, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.User, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetByEmail(email string) (*models.User, error) {
	item, err := s.repository.getByEmail(email)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//
//func (s *Service) EmailExists(email string) (bool, error) {
//	item, err := s.repository.emailExists(email)
//	if err != nil {
//		return item, err
//	}
//	return item, nil
//}
//
//func (s *Service) AddToUser(values map[string]interface{}, table string) error {
//	err := s.repository.addToUser(values, table)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s *Service) RemoveFromUser(values map[string]interface{}, table string) error {
//	err := s.repository.removeFromUser(values, table)
//	if err != nil {
//		return err
//	}
//	return nil
//}
