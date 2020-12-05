// database/database.go

package database

import (
	"GoBoard/models"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Taipei",
		viper.GetString("database.server"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
		viper.GetInt("database.ports"),
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal database error: \n", err))
	}

	db.AutoMigrate(&models.Article{}, &models.User{})
}

func GetDB() *gorm.DB {
	return db
}

// Return a list of all the articles
func GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	result := db.Find(&articles)
	return articles, result.Error
}

func GetArticleByID(id int) (models.Article, error) {
	var article models.Article
	result := db.Where("ID = ?", id).First(&article)
	return article, result.Error
}

func CreateNewArticle(title, content string) (models.Article, error) {
	article := models.Article{Title: title, Content: content}
	result := db.Create(&article)
	return article, result.Error
}

func RegisterNewUser(username, password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("The password can't be empty")
	} else if user_avial, err := IsUsernameAvailable(username); err == nil && !user_avial {
		return errors.New("The username isn't available")
	}
	user := models.User{Username: username, Password: password}
	result := db.Create(&user)
	return result.Error
}

func IsUsernameAvailable(username string) (bool, error) {
	var user []models.User
	result := db.Where("Username = ?", username).Find(&user)
	return result.RowsAffected == 0, result.Error
}

func IsUserValid(username, password string) (bool, error) {
	var user []models.User
	result := db.Where("Username = ? AND Password = ?", username, password).Find(&user)
	return result.RowsAffected == 1, result.Error
}
