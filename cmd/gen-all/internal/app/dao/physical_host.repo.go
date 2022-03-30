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

var PhysicalHostRepoSet = wire.NewSet(wire.Struct(new(PhysicalHostRepo), "*"))

type PhysicalHostRepo struct {
	DB *gorm.DB
}

func (a *PhysicalHostRepo) Create(ctx context.Context, one *schema.PhysicalHost) error {
	err := util.GetDBWithModel(ctx, a.DB, &model.PhysicalHost{}).Create(model.PhysicalHostFrom(one)).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return err
	}
	return nil
}

func (a *PhysicalHostRepo) Update(ctx context.Context, one *schema.PhysicalHost) error {
	err := util.GetDBWithModel(ctx, a.DB, &model.PhysicalHost{}).Where("id=?", one.ID).Updates(model.PhysicalHostFrom(one)).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return err
	}
	return nil
}

func (a *PhysicalHostRepo) One(ctx context.Context, id uint64) (*schema.PhysicalHost, error) {
	one := &model.PhysicalHost{}
	err := util.GetDBWithModel(ctx, a.DB, &model.PhysicalHost{}).Where("id=?", id).First(one).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return model.PhysicalHostTo(one), nil
}

func (a *PhysicalHostRepo) Has(ctx context.Context, id uint64) (bool, error) {
	var count int64
	err := util.GetDBWithModel(ctx, a.DB, &model.PhysicalHost{}).Where("id=?", id).Count(&count).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return false, err
	}
	return count > 0, nil
}

func (a *PhysicalHostRepo) Delete(ctx context.Context, id uint64) error {
	err := util.GetDBWithModel(ctx, a.DB, new(model.PhysicalHost)).Delete(&model.PhysicalHost{}, id).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
	}
	return err
}

type PhysicalHostListReq struct {
	PageSize  int64
	PageIndex int64
	//Name      string
}

func (a *PhysicalHostRepo) List(ctx context.Context, req *PhysicalHostListReq) ([]*schema.PhysicalHost, error) {
	records := make([]*model.PhysicalHost, 0)
	err := a.listReq(ctx, req).Find(&records).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return nil, err
	}

	res := make([]*schema.PhysicalHost, len(records))
	for k, v := range records {
		res[k] = model.PhysicalHostTo(v)
	}
	return res, nil
}

func (a *PhysicalHostRepo) Count(ctx context.Context, req *PhysicalHostListReq) (int64, error) {
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

func (a *PhysicalHostRepo) listReq(ctx context.Context, req *PhysicalHostListReq) *gorm.DB {
	db := util.GetDBWithModel(ctx, a.DB, &model.PhysicalHost{})

	//if req.Name != "" {
	//	db.Where("name like ?", "%"+req.Name+"%")
	//}

	db.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize))
	return db
}
