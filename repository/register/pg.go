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

func (p *pgRepository) CreateRegister(ctx context.Context, s *model.Register) error {
	return p.getDB(ctx).Create(&s).Error
}

func (p *pgRepository) ListRegisterBySemester(ctx context.Context, semesterID string) ([]model.Register, error) {
	var register []model.Register
	err := p.getDB(ctx).Where(`semester_id = ?`, semesterID).Find(&register).Error
	return register, err
}

func (p *pgRepository) GetRegister(ctx context.Context, req *model.Register) (*model.Register, error) {
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
	return p.getDB(ctx).Where(`account_id = ? and class_id = ? and semester_id = ?`, req.AccountID, req.ClassID, req.SemesterID).
		Update("is_canceled", gorm.Expr("!is_canceled")).Error
}
