package register

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewClassPG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) Create(ctx context.Context, s *model.Register) error {
	return p.getDB(ctx).Create(&s).Error
}

func (p *pgRepository) GetListBySemesterID(ctx context.Context, semesterID string) ([]model.Register, error) {
	var register []model.Register
	err := p.getDB(ctx).Where(`semester_id = ?`, semesterID).Find(&register).Error
	return register, err
}

func (p *pgRepository) Get(ctx context.Context, req *model.Register) (*model.Register, error) {
	var register *model.Register
	err := p.getDB(ctx).Where(`account_id = ? and semester_id = ? and class_id = ? and course_id = ?`,
		req.AccountID,
		req.SemesterID,
		req.ClassID,
		req.CourseID).Take(&register).Error
	return register, err
}

// swap the state of the is_canceled field
// false -> true and true -> false
func (p *pgRepository) BatchUpdateSwapIsCanCeledStatus(ctx context.Context, req *model.Register) error {
	return p.getDB(ctx).Model(&model.Register{}).Where(`account_id = ? and class_id = ? and semester_id = ?`, req.AccountID, req.ClassID, req.SemesterID).
		Update("is_canceled", gorm.Expr(" !is_canceled")).Error
}

// GetListByFirstCourseChar is get list all the course that student registered
// use the first character of course_id
// ex:
//
//	student registered S0001, T0001, M0001
//
// with S in param so the result is:
//
//	[S0001]
func (p *pgRepository) GetListByFirstCourseChar(ctx context.Context, firstChar string, accountID uint) ([]model.Register, error) {
	var registers []model.Register
	err := p.getDB(ctx).Where(`account_id = ? and substring(course_id, 1, 1) = ?`, accountID, firstChar).Find(&registers).Error
	return registers, err
}

func (p *pgRepository) GetListRegistered(ctx context.Context, accountID uint, semesterID string) ([]model.Register, error) {
	var registers []model.Register
	err := p.getDB(ctx).Where(`account_id = ? and semester_id = ?`, accountID, semesterID).Find(&registers).Error
	return registers, err
}
