package XAddress

import (
	"context"
	"time"

	"github.com/yumo001/simple-learn-rpc/ent/xaddress"
	"github.com/yumo001/simple-learn-rpc/ent/predicate"
	"github.com/yumo001/simple-learn-rpc/internal/svc"
	"github.com/yumo001/simple-learn-rpc/internal/utils/dberrorhandler"
	"github.com/yumo001/simple-learn-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetXAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetXAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetXAddressListLogic {
	return &GetXAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetXAddressListLogic) GetXAddressList(in *example.XAddressListReq) (*example.XAddressListResp, error) {
	var predicates []predicate.XAddress
	if in.CreatedAt != nil {
		predicates = append(predicates, xaddress.CreatedAtGTE(time.UnixMilli(*in.CreatedAt)))
	}
	if in.UpdatedAt != nil {
		predicates = append(predicates, xaddress.UpdatedAtGTE(time.UnixMilli(*in.UpdatedAt)))
	}
	if in.UserId != nil {
		predicates = append(predicates, xaddress.UserIDEQ(*in.UserId))
	}
	if in.Default != nil {
		predicates = append(predicates, xaddress.DefaultEQ(*in.Default))
	}
	if in.FirstName != nil {
		predicates = append(predicates, xaddress.FirstNameContains(*in.FirstName))
	}
	if in.LastName != nil {
		predicates = append(predicates, xaddress.LastNameContains(*in.LastName))
	}
	if in.CountryId != nil {
		predicates = append(predicates, xaddress.CountryIDEQ(*in.CountryId))
	}
	if in.Street != nil {
		predicates = append(predicates, xaddress.StreetContains(*in.Street))
	}
	if in.Province != nil {
		predicates = append(predicates, xaddress.ProvinceContains(*in.Province))
	}
	if in.City != nil {
		predicates = append(predicates, xaddress.CityContains(*in.City))
	}
	if in.PostalCode != nil {
		predicates = append(predicates, xaddress.PostalCodeContains(*in.PostalCode))
	}
	if in.Phone != nil {
		predicates = append(predicates, xaddress.PhoneContains(*in.Phone))
	}
	result, err := l.svcCtx.DB.XAddress.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &example.XAddressListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &example.XAddressInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			UserId:	&v.UserID,
			Default:	&v.Default,
			FirstName:	&v.FirstName,
			LastName:	&v.LastName,
			CountryId:	&v.CountryID,
			Street:	&v.Street,
			Province:	&v.Province,
			City:	&v.City,
			PostalCode:	&v.PostalCode,
			Phone:	&v.Phone,
		})
	}

	return resp, nil
}
