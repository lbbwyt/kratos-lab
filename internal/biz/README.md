# Biz

定义业务model

```azure
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
	    
```

定义Repo层的接口

```azure

// GreeterRepo is a Greater repo.
type IGreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}
```

业务逻辑实现
```azure
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

```

