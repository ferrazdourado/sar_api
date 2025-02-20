package services

import (
    "context"
    "errors"
    "github.com/ferrazdourado/sar_api/internal/models"
    "github.com/ferrazdourado/sar_api/internal/repository/interfaces"
    "github.com/ferrazdourado/sar_api/pkg/utils"
    "github.com/ferrazdourado/sar_api/pkg/config"
    "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userRepo interfaces.UserRepository
    config   *config.Config
}
var (
    // Erros de autenticação
    ErrInvalidCredentials = errors.New("credenciais inválidas")
    ErrUserExists        = errors.New("usuário já existe")
    ErrUserNotFound      = errors.New("usuário não encontrado")
    
    // Erros de VPN
    ErrInvalidConfig     = errors.New("configuração VPN inválida")
    ErrConfigNotFound    = errors.New("configuração VPN não encontrada")
    ErrVPNConnection     = errors.New("erro na conexão VPN")
)

func NewAuthService(userRepo interfaces.UserRepository, cfg *config.Config) *AuthService {
    return &AuthService{
        userRepo: userRepo,
        config:   cfg,
    }
}

func (s *AuthService) Register(ctx context.Context, user *models.User) error {
    // Verificar se usuário já existe
    existingUser, err := s.userRepo.FindByUsername(ctx, user.Username)
    if err != nil {
        return err
    }
    if existingUser != nil {
		return ErrUserExists
    }

    // Hash da senha
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    // Definir role padrão se não especificada
    if user.Role == "" {
        user.Role = "user"
    }

    return s.userRepo.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, credentials models.LoginCredentials) (string, error) {
    user, err := s.userRepo.FindByUsername(ctx, credentials.Username)
    if err != nil {
        return "", err
    }
    if user == nil {
        return "", ErrUserNotFound
    }

    // Verificar senha
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
    if err != nil {
        return "", ErrInvalidCredentials
    }

    // Criar claims com os dados do usuário
    claims := utils.Claims{
        UserID: user.ID.Hex(),
        Role:   user.Role,
    }

    // Gerar token JWT passando as claims e a configuração
    token, err := utils.GenerateToken(claims, &s.config.JWT)
    if err != nil {
        return "", err
    }

    return token, nil
}