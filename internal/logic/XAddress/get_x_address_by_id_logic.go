package XAddress

import (
	"context"

	"github.com/yumo001/simple-learn-rpc/internal/svc"
	"github.com/yumo001/simple-learn-rpc/internal/utils/dberrorhandler"
	"github.com/yumo001/simple-learn-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetXAddressByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetXAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetXAddressByIdLogic {
	return &GetXAddressByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetXAddressByIdLogic) GetXAddressById(in *example.IDReq) (*example.XAddressInfo, error) {
	result, err := l.svcCtx.DB.XAddress.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.XAddressInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		UserId:	&result.UserID,
		Default:	&result.Default,
		FirstName:	&result.FirstName,
		LastName:	&result.LastName,
		CountryId:	&result.CountryID,
		Street:	&result.Street,
		Province:	&result.Province,
		City:	&result.City,
		PostalCode:	&result.PostalCode,
		Phone:	&result.Phone,
	}, nil
}

