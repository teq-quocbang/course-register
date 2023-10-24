package grpc

import (
	"github.com/teq-quocbang/course-register/proto"
	"github.com/teq-quocbang/course-register/usecase"
)

type TeqService struct {
	proto.UnimplementedTeqServiceServer
	UseCase *usecase.UseCase
}
