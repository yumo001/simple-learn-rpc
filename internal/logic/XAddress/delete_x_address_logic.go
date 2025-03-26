package XAddress

import (
	"context"

    "github.com/yumo001/simple-learn-rpc/ent/xaddress"
    "github.com/yumo001/simple-learn-rpc/internal/svc"
    "github.com/yumo001/simple-learn-rpc/internal/utils/dberrorhandler"
    "github.com/yumo001/simple-learn-rpc/types/example"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteXAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteXAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteXAddressLogic {
	return &DeleteXAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteXAddressLogic) DeleteXAddress(in *example.IDsReq) (*example.BaseResp, error) {
	_, err := l.svcCtx.DB.XAddress.Delete().Where(xaddress.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &example.BaseResp{Msg: errormsg.DeleteSuccess }, nil
}
