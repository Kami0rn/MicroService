package authRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// AuthRepository interface
	AuthRepositoryService interface{}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db :db}
}

func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	_ = pctx // แสดงให้เห็นว่า pctx ไม่ถูกใช้งานโดยตั้งใจ
	return r.db.Database("auth_db")
}