package sevice

import (
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/biz"
	"context"
)

type MyappService struct {
	// if need to work for grpc server. Need grpc protoc gen server.
	// grc.Server

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewMyappService new a service.
func NewMyappService(uc *biz.UserUsecase, logger log.Logger) *MyappService {
	return &MyappService{uc: uc, log: log.NewHelper("service/myapp", logger)}
}


// 这里的 request 和 response 可以自己声明，也可以依靠 grpc 生成
// 但是建议 grpc 生成，因为 IDL 声明式的接口，可以对服务的入参出参进行强约束
// 如果是针对内部对象的返回，可能需要写两份，一个是内部的 Entity，一个是 grpc 内声明的 message 结构体
type InRequest struct {
	Name string `json:"name,omitempty"`
}

type OutRequest struct {
	Message string `json:"message,omitempty"`
}

// MyappService 返回的是 DTO 对象，也就是领域对象，业务的最小单元
func (s *MyappService) GetUser(ctx context.Context, in *InRequest) (*OutRequest, error) {
	return &OutRequest{Message: "Hello " + in.Name}, nil
}


