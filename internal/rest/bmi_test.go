package rest_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/rest"
	"github.com/bxcodec/go-clean-arch/internal/rest/mocks"
	faker "github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	var mockArticle domain.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)
	mockUCase := new(mocks.ArticleService)
	mockListArticle := make([]domain.Article, 0)
	mockListArticle = append(mockListArticle, mockArticle)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", mock.Anything, cursor, int64(num)).Return(mockListArticle, "10", nil)

	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(),
		echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := rest.ArticleHandler{
		Service: mockUCase,
	}
	err = handler.FetchArticle(c)
	require.NoError(t, err)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestBMICalculate(t *testing.T) {
	// Set up Echo and the request/response objects
	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(),
		echo.POST, "/bmi_calculate?weight=50&height=185", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the actual handler function directly
	handler := rest.BMIHandler{}

	// Assuming CalculateBMI is the method that handles the request
	err = handler.CalculateBMI(c)
	require.NoError(t, err)

	// Validate the response
	assert.Equal(t, http.StatusOK, rec.Code)

	// Expected JSON response
	expectedResponse := `{"message": "You are underweight."}`
	assert.JSONEq(t, expectedResponse, rec.Body.String())
}
