package dao

import (
	"context"
	"fpga-controller/internal/app/dao/model"
	"fpga-controller/internal/app/dao/util"
	"fpga-controller/internal/app/schema"
	"fpga-controller/pkg/logger"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var httpRepoSet = wire.NewSet(wire.Struct(new(httpRepo), "*"))

type httpRepo struct {
	DB *gorm.DB
}

func (a *httpRepo) Create(ctx context.Context, one *schema.http) error {
	err := util.GetDBWithModel(ctx, a.DB, &model.http{}).Create(model.httpFrom(one)).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return err
	}
	return nil
}

func (a *httpRepo) Update(ctx context.Context, one *schema.http) error {
	err := util.GetDBWithModel(ctx, a.DB, &model.http{}).Where("id=?", one.ID).Updates(model.httpFrom(one)).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return err
	}
	return nil
}

func (a *httpRepo) One(ctx context.Context, id uint64) (*schema.http, error) {
	one := &model.http{}
	err := util.GetDBWithModel(ctx, a.DB, &model.http{}).Where("id=?", id).First(one).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return model.httpTo(one), nil
}

func (a *httpRepo) Has(ctx context.Context, id uint64) (bool, error) {
	var count int64
	err := util.GetDBWithModel(ctx, a.DB, &model.http{}).Where("id=?", id).Count(&count).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return false, err
	}
	return count > 0, nil
}

func (a *httpRepo) Delete(ctx context.Context, id uint64) error {
	err := util.GetDBWithModel(ctx, a.DB, new(model.http)).Delete(&model.http{}, id).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
	}
	return err
}

type httpListReq struct {
	PageSize  int64
	PageIndex int64
	//Name      string
}

func (a *httpRepo) List(ctx context.Context, req *httpListReq) ([]*schema.http, error) {
	records := make([]*model.http, 0)
	err := a.listReq(ctx, req).Find(&records).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return nil, err
	}

	res := make([]*schema.http, len(records))
	for k, v := range records {
		res[k] = model.httpTo(v)
	}
	return res, nil
}

func (a *httpRepo) Count(ctx context.Context, req *httpListReq) (int64, error) {
	req.PageIndex = 0
	req.PageIndex = 0
	var count int64
	err := a.listReq(ctx, req).Count(&count).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return 0, err
	}
	return count, nil
}

func (a *httpRepo) listReq(ctx context.Context, req *httpListReq) *gorm.DB {
	db := util.GetDBWithModel(ctx, a.DB, &model.http{})

	//if req.Name != "" {
	//	db.Where("name like ?", "%"+req.Name+"%")
	//}

	db.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize))
	return db
}
