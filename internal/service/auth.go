package service

import (
	"strings"
	"time"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretJWT = []byte("detalles_personalizados_api")

const duracionToken = 24 * time.Hour

type Claims struct {
	Email     string `json:"email"`
	UsuarioID int    `json:"usuario_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo storage.UserRepository
}

func NewAuthService(repo storage.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

// =========================================================
// REGISTRAR
// =========================================================

func (s *AuthService) Registrar(email, password string) (models.Usuario, error) {

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return models.Usuario{}, ErrCredencialesInvalidas
	}

	if _, existe := s.repo.BuscarUsuarioPorEmail(email); existe {
		return models.Usuario{}, ErrEmailEnUso
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return models.Usuario{}, err
	}

	return s.repo.CrearUsuario(models.Usuario{
		Email:        email,
		PasswordHash: string(hash),
	})
}

// =========================================================
// LOGIN
// =========================================================

func (s *AuthService) Login(email, password string) (string, error) {

	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return "", ErrCredencialesInvalidas
	}

	usuario, existe := s.repo.BuscarUsuarioPorEmail(email)

	if !existe {
		return "", ErrCredencialesInvalidas
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(usuario.PasswordHash),
		[]byte(password),
	); err != nil {

		return "", ErrCredencialesInvalidas
	}

	return s.GenerarToken(usuario)
}

// =========================================================
// GENERAR TOKEN
// =========================================================

func (s *AuthService) GenerarToken(u models.Usuario) (string, error) {

	claims := &Claims{
		Email:     u.Email,
		UsuarioID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duracionToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretJWT)
}

// =========================================================
// VALIDAR TOKEN
// =========================================================

func (s *AuthService) ValidarToken(token string) (int, error) {

	parsedToken, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrCredencialesInvalidas
			}

			return secretJWT, nil
		},
	)

	if err != nil || !parsedToken.Valid {
		return 0, ErrCredencialesInvalidas
	}

	claims, ok := parsedToken.Claims.(*Claims)

	if !ok {
		return 0, ErrCredencialesInvalidas
	}

	return claims.UsuarioID, nil
}
