package repository

import "github.com/abdabariassalam/majoo/internal/models"

func (r repository) FindByNameAndPassword(userName, password string) (*models.User, error) {
	user := models.User{}
	err := r.db.Where("user_name = ? AND password = MD5(?)", userName, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
