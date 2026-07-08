package service

import (
	"strings"
	"time"

	"proyecto-detallesPersonalizados-api/internal/models"
	"proyecto-detallesPersonalizados-api/internal/storage"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretJWTPorDefecto = []byte("detalles_personalizados_api")

const duracionTokenPorDefecto = 24 * time.Hour

type Claims struct {
	Email     string `json:"email"`
	UsuarioID int    `json:"usuario_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo          storage.UserRepository
	secretJWT     []byte
	duracionToken time.Duration
}

type AuthOption func(*AuthService)

func WithJWTSecreto(secret []byte) AuthOption {
	return func(s *AuthService) {
		if len(secret) > 0 {
			s.secretJWT = append([]byte(nil), secret...)
		}
	}
}

func WithJWTDuracion(duracion time.Duration) AuthOption {
	return func(s *AuthService) {
		if duracion > 0 {
			s.duracionToken = duracion
		}
	}
}

func NewAuthService(repo storage.UserRepository, opts ...AuthOption) *AuthService {
	s := &AuthService{
		repo:          repo,
		secretJWT:     append([]byte(nil), secretJWTPorDefecto...),
		duracionToken: duracionTokenPorDefecto,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.duracionToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secretJWT)
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

			return s.secretJWT, nil
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
