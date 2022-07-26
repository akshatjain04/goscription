package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/kecci/goscription/internal/controller"
	"github.com/kecci/goscription/mocks"
	"github.com/kecci/goscription/models"
	"github.com/kecci/goscription/utility"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bxcodec/faker"
)

func TestFetch(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)
	mockUCase := new(mocks.ArticleService)
	mockListArticle := make([]models.Article, 0)
	mockListArticle = append(mockListArticle, mockArticle)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", mock.Anything, cursor, int64(num)).Return(mockListArticle, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/articles?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	controller.InitArticleController(e, mockUCase)
	e.ServeHTTP(rec, req)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUCase := new(mocks.ArticleService)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", mock.Anything, cursor, int64(num)).Return(nil, "", utility.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/articles?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	controller.InitArticleController(e, mockUCase)
	e.ServeHTTP(rec, req)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "", responseCursor)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockUCase := new(mocks.ArticleService)
	num := int(mockArticle.ID)
	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(mockArticle, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/articles/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	controller.InitArticleController(e, mockUCase)
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockArticle := models.Article{
		Title:     "Title",
		Content:   "Content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempMockArticle := mockArticle
	tempMockArticle.ID = 0
	mockUCase := new(mocks.ArticleService)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)
	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("service.ArticleParam")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/articles", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	controller.InitArticleController(e, mockUCase)
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockUCase := new(mocks.ArticleService)
	num := int(mockArticle.ID)
	mockUCase.On("Delete", mock.Anything, int64(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/articles/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	controller.InitArticleController(e, mockUCase)
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)
}
