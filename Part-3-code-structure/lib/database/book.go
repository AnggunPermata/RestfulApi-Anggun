package database

import (
	"project/anggun/config"
	"project/anggun/models"

	"github.com/labstack/echo"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	/*
		tx := config.DB.Find(&users)
		err := tx.Error
	*/
	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBook(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func CreateBook(c echo.Context) (interface{}, error) {
	book := models.Book{}
	c.Bind(&book)
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id = ?", id).Delete(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(id int, book interface{}) (interface{}, error) {
	if err := config.DB.Find("id=?", id).Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

/*
func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email=? AND password=?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
*/
