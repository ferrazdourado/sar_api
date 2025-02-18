package controllers_test

import (
    "context"
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/ferrazdourado/sar_api/internal/controllers"
    "github.com/ferrazdourado/sar_api/internal/models"
    "github.com/ferrazdourado/sar_api/internal/services"
)

// Interface do serviço de autenticação
type AuthServiceInterface interface {
    Login(ctx context.Context, credentials models.LoginCredentials) (string, error)
    Register(ctx context.Context, user *models.User) error
}

// Mock do serviço de autenticação
type MockAuthService struct {
    mock.Mock
}

func (m *MockAuthService) Login(ctx context.Context, credentials models.LoginCredentials) (string, error) {
    args := m.Called(ctx, credentials)
    return args.String(0), args.Error(1)
}

func (m *MockAuthService) Register(ctx context.Context, user *models.User) error {
    args := m.Called(ctx, user)
    return args.Error(0)
}

func TestLogin(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name           string
        credentials    models.LoginCredentials
        setupMock      func(*MockAuthService)
        expectedCode   int
        expectedBody   map[string]interface{}
    }{
        {
            name: "Login bem sucedido",
            credentials: models.LoginCredentials{
                Username: "testuser",
                Password: "123456",
            },
            setupMock: func(m *MockAuthService) {
                m.On("Login", mock.Anything, models.LoginCredentials{
                    Username: "testuser",
                    Password: "123456",
                }).Return("jwt.token.aqui", nil)
            },
            expectedCode: http.StatusOK,
            expectedBody: map[string]interface{}{
                "token": "jwt.token.aqui",
            },
        },
        {
            name: "Credenciais inválidas",
            credentials: models.LoginCredentials{
                Username: "wrong",
                Password: "wrong",
            },
            setupMock: func(m *MockAuthService) {
                m.On("Login", mock.Anything, models.LoginCredentials{
                    Username: "wrong",
                    Password: "wrong",
                }).Return("", services.ErrInvalidCredentials)
            },
            expectedCode: http.StatusUnauthorized,
            expectedBody: map[string]interface{}{
                "error": "credenciais inválidas",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockService := new(MockAuthService)
            tt.setupMock(mockService)

            controller := controllers.NewAuthController(mockService)
            router := gin.New()
            router.POST("/login", controller.Login)

            body, _ := json.Marshal(tt.credentials)
            req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")
            w := httptest.NewRecorder()

            router.ServeHTTP(w, req)

            assert.Equal(t, tt.expectedCode, w.Code)

            var response map[string]interface{}
            err := json.Unmarshal(w.Body.Bytes(), &response)
            assert.NoError(t, err)
            assert.Equal(t, tt.expectedBody, response)

            mockService.AssertExpectations(t)
        })
    }
}

func TestRegister(t *testing.T) {
    gin.SetMode(gin.TestMode)

    tests := []struct {
        name          string
        user          models.User
        mockError    error
        expectedCode int
        expectedBody map[string]interface{}
    }{
        {
            name: "Registro bem sucedido",
            user: models.User{
                Username: "newuser",
                Password: "123456",
                Email:    "new@user.com",
            },
            mockError:    nil,
            expectedCode: http.StatusCreated,
            expectedBody: map[string]interface{}{
                "message": "Usuário registrado com sucesso",
            },
        },
        {
            name: "Usuário já existe",
            user: models.User{
                Username: "existinguser",
                Password: "123456",
                Email:    "existing@user.com",
            },
            mockError:    services.ErrUserExists,
            expectedCode: http.StatusBadRequest,
            expectedBody: map[string]interface{}{
                "error": "Usuário já existe",
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockService := new(MockAuthService)
            mockService.On("Register", &tt.user).Return(tt.mockError)

            controller := controllers.NewAuthController(mockService)
            router := gin.New()
            router.POST("/register", controller.Register)

            body, _ := json.Marshal(tt.user)
            req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")
            w := httptest.NewRecorder()

            router.ServeHTTP(w, req)

            assert.Equal(t, tt.expectedCode, w.Code)

            var response map[string]interface{}
            err := json.Unmarshal(w.Body.Bytes(), &response)
            assert.NoError(t, err)
            assert.Equal(t, tt.expectedBody, response)

            mockService.AssertExpectations(t)
        })
    }
}