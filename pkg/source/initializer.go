package source

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SubInitializer interface {
	InitializerName() string // 不一定代表单独一个表，所以改成了更宽泛的语义
	MigrateTable(ctx context.Context, db *gorm.DB) error
	InitializeData(ctx context.Context, db *gorm.DB) error
	TableCreated(ctx context.Context, db *gorm.DB) bool
}

type InitHandler interface {
	InitTables(ctx context.Context, inits initSlice) error // 建表 handler
}

type initSlice []SubInitializer

var initializers initSlice

// RegisterInit 注册要执行的初始化过程，会在 InitDB() 时调用
func RegisterInit(i SubInitializer) {
	if initializers == nil {
		initializers = initSlice{}
	}
	initializers = append(initializers, i)
}

type initDBService struct {
	db *gorm.DB
}

func NewInitDBService(db *gorm.DB) *initDBService {
	return &initDBService{db: db}
}

func (i *initDBService) InitDB() error {
	if len(initializers) == 0 {
		return errors.New("无可用初始化过程，请检查初始化是否已执行完成")
	}
	// TODO support more database
	handler := newMysqlInitHandler(i.db)
	return handler.InitTables(context.TODO(), initializers)
}
