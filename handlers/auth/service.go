package auth

import (
	"errors"
	"mdgkb/mdgkb-server/handlers/users"
	"mdgkb/mdgkb-server/models"
)

func (s *Service) Register(item *models.User) (*models.TokensWithUser, error) {
	//err := item.GenerateHashPassword()
	//if err != nil {
	//	return nil, err
	//}
	err := users.CreateService(s.repository.getDB(), s.helper).Create(item)
	if err != nil {
		return nil, err
	}
	ts, err := s.helper.Token.CreateToken(item.ID.String())
	if err != nil {
		return nil, err
	}
	return &models.TokensWithUser{Tokens: ts, User: *item}, nil
}

func (s *Service) Login(item *models.User) (*models.TokensWithUser, error) {
	//err := item.GenerateHashPassword()
	//if err != nil {
	//	return nil, err
	//}
	findedUser, err := users.CreateService(s.repository.getDB(), s.helper).GetByEmail(item.Email)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item.Password)
	//fmt.Println(findedUser.CompareWithHashPassword(&item.Password))
	if findedUser.Password != item.Password {
		//err =
		return nil, errors.New("wrong password")
	}
	//if !findedUser.CompareWithHashPassword(&item.Password) {
	//	return nil, err
	//}

	ts, err := s.helper.Token.CreateToken(findedUser.ID.String())
	if err != nil {
		return nil, err
	}

	//saveErr := s.helper.CreateAuth(item.ID.String(), ts, s.redis)
	//if saveErr != nil {
	//	return nil, err
	//}
	//fmt.Println(ts)
	//fmt.Println(findedUser)
	return &models.TokensWithUser{Tokens: ts, User: *findedUser}, nil
}
