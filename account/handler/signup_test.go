package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/masnann/memo/account/model"
	"github.com/masnann/memo/account/model/apperrors"
	"github.com/masnann/memo/account/model/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Email and Password Required", func(t *testing.T) {
		// We just want this to show that it's not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email": "",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Invialid Email", func(t *testing.T) {
		// We just want this to show that it's not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@bo",
			"password": "avalidpassword123",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Password too short", func(t *testing.T) {
		// We just want this to show that it's not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@bob.com",
			"password": "inval",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Password too long", func(t *testing.T) {
		// We just want this to show that it's not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password

		reqBody, err := json.Marshal(gin.H{
			"email":    "bob@bob.com",
			"password": "invalkadsfjasdfkj;askldfj;askldfj;asdfiuerueuuuuuudfjasdfasdkfjj",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Error calling UserService", func(t *testing.T) {
		u := &model.User{
			Email:    "bob@bob.com",
			Password: "avalidpassword",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), u).Return(apperrors.NewConflict("User Already Exists", u.Email))

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 409, rr.Code)
		mockUserService.AssertExpectations(t)
	})

}