// tests/article_model_test.go

package tests

import (
	"GoBoard/models"
	"testing"
)

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := models.GetAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(models.ArticleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != models.ArticleList[i].Content ||
			v.ID != models.ArticleList[i].ID ||
			v.Title != models.ArticleList[i].Title {

			t.Fail()
			break
		}
	}
}

func TestCreateNewArticle(t *testing.T) {
	originalLength := len(models.GetAllArticles())

	a, err := models.CreateNewArticle("New test title", "New test content")

	allArticles := models.GetAllArticles()
	newLength := len(allArticles)

	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test title" || a.Content != "New test content" {

		t.Fail()
	}
}
