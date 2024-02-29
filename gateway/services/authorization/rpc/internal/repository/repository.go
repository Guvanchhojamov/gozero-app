package repository

import (
	"context"
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/models"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/response"
	"github.com/go-errors/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"strings"
	"time"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, userInput *models.SignUp, salt string) (userId uint32, err error) {
	query := fmt.Sprintf(`INSERT INTO users (login,role_id, password) VALUES ($1, $2, $3) RETURNING id`)
	if err = r.db.QueryRow(ctx, query, userInput.Login, userInput.RoleId, generateHashPassword(userInput.Password, salt)).Scan(&userId); err != nil {
		if strings.Contains(err.Error(), errUserExistsPG) {
			return 0, errors.New("user already exists")
		}
	}

	return userId, nil
}

func (r *Repository) GenerateToken(ctx context.Context, username, password, jwtSignedKey, salt string) (string, error) {
	// get user from DB
	user, err := r.getUserFromDB(ctx, username, generateHashPassword(password, salt))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.CustomTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(tokenTLL)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		UserId: user.Id,
	})
	return token.SignedString([]byte(jwtSignedKey))
}

func (r *Repository) getUserFromDB(ctx context.Context, login string, passHash string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password=$2", usersTable)
	err := r.db.QueryRow(ctx, query, login, passHash).Scan(&user.Id)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, response.NewRpcErrResponse(codes.NotFound, errors.New(err))
	}
	return user, nil
}
