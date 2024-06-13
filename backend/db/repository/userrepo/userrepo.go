package userrepo

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
	"wt-planning/db/models"
)

type UserRepository interface {
	FindAll(filter map[string]any) ([]models.User, error)
	Update(user *models.User) error

	// Redis methods
	SetMessageResponse(ctx context.Context, key string, value string) error
	GetMessageResponse(ctx context.Context, key string) (string, error)
}

//go:generate mockgen -destination=../../../mocks/repository/userrepo/userrepo.go -package=userrepo -source=userrepo.go
type dataAccess struct {
	db      *gorm.DB
	redisDb *redis.Client
}

func New(db *gorm.DB, redisDb *redis.Client) UserRepository {
	return &dataAccess{
		db:      db,
		redisDb: redisDb,
	}
}

func (d *dataAccess) FindAll(filter map[string]any) (user []models.User, err error) {
	// Only status is pending data will be fetched
	result := d.db.Where("status = ?", filter["status"]).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (d *dataAccess) Update(user *models.User) error {
	result := d.db.Table("public.users").Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *dataAccess) SetMessageResponse(ctx context.Context, key string, value string) error {

	// Cache to redis 1 hour
	result := *d.redisDb.Set(ctx, key, value, time.Hour)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (d *dataAccess) GetMessageResponse(ctx context.Context, key string) (string, error) {
	val, err := d.redisDb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
