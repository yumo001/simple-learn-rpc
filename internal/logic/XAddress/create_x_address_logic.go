package XAddress

import (
	"context"

	"github.com/yumo001/simple-learn-rpc/internal/svc"
	"github.com/yumo001/simple-learn-rpc/internal/utils/dberrorhandler"
	"github.com/yumo001/simple-learn-rpc/types/example"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateXAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateXAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateXAddressLogic {
	return &CreateXAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateXAddressLogic) CreateXAddress(in *example.XAddressInfo) (*example.BaseIDResp, error) {
    result, err := l.svcCtx.DB.XAddress.Create().
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
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &example.BaseIDResp{Id: result.ID, Msg: errormsg.CreateSuccess }, nil
}
