package base

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/yumo001/simple-learn-rpc/internal/svc"
	"github.com/yumo001/simple-learn-rpc/types/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *example.Empty) (*example.BaseResp, error) {
    if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
        logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
        return nil, errorx.NewInternalError(err.Error())
    }

	return &example.BaseResp{Msg: errormsg.Success}, nil
}
