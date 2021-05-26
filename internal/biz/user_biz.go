package biz

import (
	"github.com/blackbinbinbinbin/Bingo-gin/internal/ent"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
	"context"
)


// 这里因为直接应用了 orm , 所以直接包入 UserEntity，如果没有使用则
type UserEntity struct {
	ent.User
}

// 这里声明下层 repo 持久层的接口形式，并且声明对 entity 对象的操作方法，一般都有：
// Create()
// Update()
// Select()
// Delete()
// 这些需要持久层 data 自行实现
type UserRepo interface {
	CreateUser(context.Context, *UserEntity) error
	UpdateUser(context.Context, *UserEntity) error
}




// 这里主要是业务逻辑层：
// 负责 DTO 和 DO 数据的相互转化，往上屏蔽了下层持久层代码细节，往下隔离了业务层的逻辑细节，让业务代码聚合，而不是散落在各处
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper("usecase/user", logger)}
}

func (uc *UserUsecase) Create(ctx context.Context,u *UserEntity) error {
	return uc.repo.
		CreateUser(ctx,u)
}

func (uc *UserUsecase) Update(ctx context.Context,u *UserEntity) error {
	return uc.repo.
		UpdateUser(ctx,u)
}