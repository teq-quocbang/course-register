package account

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	getDB func(context.Context) *gorm.DB
}

func NewAccountPG(getDB func(context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB: getDB}
}

func (p *pgRepository) CreateAccount(ctx context.Context, req *model.Account) (uint, error) {
	if err := p.getDB(ctx).Create(req).Error; err != nil {
		return 0, err
	}
	return req.ID, nil
}

func (p *pgRepository) GetAccountByID(ctx context.Context, studentID uint) (*model.Account, error) {
	return nil, nil
}

func (p *pgRepository) CreateVerifyAccount(ctx context.Context, req *model.AccountVerify) error {
	return nil
}

func (p *pgRepository) GetVerifyAccountByID(ctx context.Context, studentID uint) (*model.AccountVerify, error) {
	return nil, nil
}

func (p *pgRepository) GetByAccountConstraint(ctx context.Context, req *model.Account) (*model.Account, error) {
	var account *model.Account
	if err := p.getDB(ctx).Where(`username = ? or email = ?`, req.Username, req.Email).
		Take(&account).Error; err != nil {
		return nil, err
	}

	return account, nil
}
