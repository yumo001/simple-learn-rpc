package XAddress

import (
	"context"

	"github.com/yumo001/simple-learn-rpc/internal/svc"
	"github.com/yumo001/simple-learn-rpc/internal/utils/dberrorhandler"
	"github.com/yumo001/simple-learn-rpc/types/example"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateXAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateXAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateXAddressLogic {
	return &UpdateXAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateXAddressLogic) UpdateXAddress(in *example.XAddressInfo) (*example.BaseResp, error) {
	err:= l.svcCtx.DB.XAddress.UpdateOneID(*in.Id).
			SetNotNilUserID(in.UserId).
			SetNotNilDefault(in.Default).
			SetNotNilFirstName(in.FirstName).
			SetNotNilLastName(in.LastName).
			SetNotNilCountryID(in.CountryId).
			SetNotNilStreet(in.Street).
			SetNotNilProvince(in.Province).
			SetNotNilCity(in.City).
			SetNotNilPostalCode(in.PostalCode).
			SetNotNilPhone(in.Phone).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &example.BaseResp{Msg: errormsg.UpdateSuccess }, nil
}
