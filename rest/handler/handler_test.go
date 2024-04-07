package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	h := handler{}
	c := e.NewContext(req, rec)

	// Act
	err := h.Greeting(c)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World! DevOps", rec.Body.String())
	}
}

func TestListNews(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/news", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	newsMockRow := sqlmock.NewRows([]string{"ID", "Title", "Content", "Author"}).
		AddRow(1, "Title 1", "Content 1", "Author 1")

	db, mock, err := sqlmock.New()
	mock.ExpectQuery("SELECT (.+) FROM news").WillReturnRows(newsMockRow)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub datase connection", err)
	}
	h := handler{db}
	c := e.NewContext(req, rec)
	expected := `[{"ID":1,"Title":"Title 1","Content":"Content 1","Author":"Author 1"}]`

	// Act
	err = h.ListNews(c)

	// Assertions
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}

}
