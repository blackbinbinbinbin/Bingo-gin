package data

import (
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
	"context"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper("data/user", logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.UserEntity) error {
	// 这里使用 r.data 类去操作持久化
	// 这里并不关心吃操作 mysql 还是缓存 redis
	// 只是声明了需要存储 biz.UserEntity
	return nil
}

func (r *userRepo) UpdateUser(ctx context.Context, u *biz.UserEntity) error {
	return nil
}