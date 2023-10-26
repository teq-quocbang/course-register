package semester

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewSemesterPG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) CreateSemester(ctx context.Context, s *model.Semester) error {
	return p.getDB(ctx).Create(&s).Error
}

func (p *pgRepository) ListSemesterByYear(ctx context.Context, year string) ([]model.Semester, error) {
	var semesters []model.Semester
	err := p.getDB(ctx).Where(`year(start_time) = ?`, year).Order("start_time DESC").Find(&semesters).Error
	return semesters, err
}

func (p *pgRepository) GetSemester(ctx context.Context, semesterID string) (model.Semester, error) {
	var semester model.Semester
	err := p.getDB(ctx).Where(`id = ?`, semesterID).Take(&semester).Error
	return semester, err
}
