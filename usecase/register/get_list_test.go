package register

import (
	"context"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/course-register/codetype"
	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/repository/class"
	"github.com/teq-quocbang/course-register/repository/course"
	"github.com/teq-quocbang/course-register/repository/register"
	"github.com/teq-quocbang/course-register/repository/semester"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
	"github.com/teq-quocbang/course-register/util/token"
)

func (s *TestSuite) TestGetList() {
	assertion := s.Assertions

	testSemesterID := "S0001"
	testCourseID := "M0001"
	testClassID := "CL0001"
	userPrinciple := monkey.Patch(contexts.GetUserPrincipleByContext, func(context.Context) *token.JWTClaimCustom {
		return &token.JWTClaimCustom{
			SessionID: uuid.New(),
			User: token.UserInfo{
				ID:       1,
				Username: "test_username",
				Email:    "test@teqnological.asia",
			},
		}
	})
	defer monkey.Unpatch(userPrinciple)
	// good case
	{
		// Arrange
		mockSemesterRepo := semester.NewMockRepository(s.T())
		mockCourseRepo := course.NewMockRepository(s.T())
		mockClassRepo := class.NewMockRepository(s.T())
		mockRegisterRepo := register.NewMockRepository(s.T())
		mockRegisterRepo.EXPECT().GetListBySemesterID(s.ctx, uint(1), testSemesterID, []string{}, codetype.Paginator{
			Page:  1,
			Limit: 20,
		}).ReturnArguments = mock.Arguments{
			[]model.Register{
				{
					AccountID:  uint(1),
					SemesterID: testSemesterID,
					ClassID:    testClassID,
					CourseID:   testCourseID,
				},
			}, int64(1), nil,
		}

		mockSemesterRepo.EXPECT().GetByID(s.ctx, testSemesterID).ReturnArguments = mock.Arguments{
			model.Semester{
				ID: testSemesterID,
			}, nil,
		}
		mockClassRepo.EXPECT().GetByID(s.ctx, testClassID).ReturnArguments = mock.Arguments{
			model.Class{
				ID: testSemesterID,
			}, nil,
		}
		mockCourseRepo.EXPECT().GetByID(s.ctx, testCourseID).ReturnArguments = mock.Arguments{
			model.Course{
				ID: testCourseID,
			}, nil,
		}
		u := s.useCase(mockRegisterRepo, mockSemesterRepo, mockClassRepo, mockCourseRepo)
		req := &payload.ListRegisterInformationRequest{
			SemesterID: testSemesterID,
		}

		// Act
		reply, err := u.GetListBySemester(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.Equal(1, len(reply.Register))
	}

	// bad case
	{
		// Arrange
		req := &payload.ListRegisterInformationRequest{}
		u := s.useCase(register.NewMockRepository(s.T()), semester.NewMockRepository(s.T()), class.NewMockRepository(s.T()), course.NewMockRepository(s.T()))

		// Act
		_, err := u.GetListBySemester(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrRegisterInvalidParam("Key: 'ListRegisterInformationRequest.SemesterID' Error:Field validation for 'SemesterID' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
