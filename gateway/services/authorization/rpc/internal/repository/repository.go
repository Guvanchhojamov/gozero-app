package repository

import (
	"context"
	"fmt"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/models"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/internal/response"
	"github.com/Guvanchhojamov/gozero-app/gateway/services/authorization/rpc/userauthservice"
	"github.com/go-errors/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"time"
)

type Repository struct {
	db *pgxpool.Pool
}

const (
	usersTable = "users"
)

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, userInput *userauthservice.SignUpRequest) (userId uint32, err error) {
	query := fmt.Sprintf(`INSERT INTO users (login,role_id, password) VALUES ($1, $2, $3, ...)`)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return 0, err
	}
	fmt.Println(rows)
	return 0, nil
}

func (r *Repository) GenerateToken(ctx context.Context, username, password string) (string, error) {
	// get user from DB
	user, err := r.getUserFromDB(ctx, username, generateHashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.CustomTokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(tokenTLL)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		user.Id,
	})
	return token.SignedString([]byte(signedKey))
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

//func (r *Repository) CheckRolePersmission(ctx context.Context, userId uint32) {
//	query := fmt.Sprintf(`SELECT roles.id AS role_id, roles.name FROM roles JOIN users ON roles.id = users.role_id WHERE users.id = 1`)
//}
