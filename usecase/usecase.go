package usecase

import (
	"github.com/teq-quocbang/course-register/repository"
	"github.com/teq-quocbang/course-register/usecase/example"
	"github.com/teq-quocbang/course-register/usecase/grpc"
	myS3 "github.com/teq-quocbang/course-register/util/s3"
	mySES "github.com/teq-quocbang/course-register/util/ses"
)

type UseCase struct {
	Example example.IUseCase
	GRPC    grpc.IUseCase

	SES mySES.ISES
	S3  myS3.IS3
}

func New(repo *repository.Repository) *UseCase {
	var (
		ses = mySES.NewSES()
		s3  = myS3.NewS3()
	)

	return &UseCase{
		Example: example.New(repo, ses),
		GRPC:    grpc.New(repo),
		SES:     ses,
		S3:      s3,
	}
}
