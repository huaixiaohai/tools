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

var testRepoSet = wire.NewSet(wire.Struct(new(testRepo), "*"))

type testRepo struct {
	DB *gorm.DB
}

func (a *testRepo) Create(ctx context.Context, one *schema.test) error {
	err := util.GetDBWithModel(ctx, a.DB, &model.test{}).Create(model.testFrom(one)).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return err
	}
	return nil
}

func (a *testRepo) Update(ctx context.Context, one *schema.test) error {
	err := util.GetDBWithModel(ctx, a.DB, &model.test{}).Where("id=?", one.ID).Updates(model.testFrom(one)).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return err
	}
	return nil
}

func (a *testRepo) One(ctx context.Context, id uint64) (*schema.test, error) {
	one := &model.test{}
	err := util.GetDBWithModel(ctx, a.DB, &model.test{}).Where("id=?", id).First(one).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return model.testTo(one), nil
}

func (a *testRepo) Has(ctx context.Context, id uint64) (bool, error) {
	var count int64
	err := util.GetDBWithModel(ctx, a.DB, &model.test{}).Where("id=?", id).Count(&count).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return false, err
	}
	return count > 0, nil
}

func (a *testRepo) Delete(ctx context.Context, id uint64) error {
	err := util.GetDBWithModel(ctx, a.DB, new(model.test)).Delete(&model.test{}, id).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
	}
	return err
}

type testListReq struct {
	PageSize  int64
	PageIndex int64
	//Name      string
}

func (a *testRepo) List(ctx context.Context, req *testListReq) ([]*schema.test, error) {
	records := make([]*model.test, 0)
	err := a.listReq(ctx, req).Find(&records).Error
	if err != nil {
		logger.WithContext(ctx).Error(err)
		return nil, err
	}

	res := make([]*schema.test, len(records))
	for k, v := range records {
		res[k] = model.testTo(v)
	}
	return res, nil
}

func (a *testRepo) Count(ctx context.Context, req *testListReq) (int64, error) {
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

func (a *testRepo) listReq(ctx context.Context, req *testListReq) *gorm.DB {
	db := util.GetDBWithModel(ctx, a.DB, &model.test{})

	//if req.Name != "" {
	//	db.Where("name like ?", "%"+req.Name+"%")
	//}

	db.Limit(int(req.PageSize)).Offset(int((req.PageIndex - 1) * req.PageSize))
	return db
}
