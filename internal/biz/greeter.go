package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "helloworld/api/helloworld/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	gorm.Model
	Hello string `gorm:"hello"`
}

func (d *Greeter) ToProto() *v1.GreeterEntity {
	return &v1.GreeterEntity{
		Hello: d.Hello,
	}
}

// GreeterRepo is a Greater repo.
type IGreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo IGreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo IGreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

func (uc *GreeterUsecase) ListGreeter(ctx context.Context) ([]*Greeter, error) {
	uc.log.WithContext(ctx).Info("ListGreeter")
	return uc.repo.ListAll(ctx)
}
