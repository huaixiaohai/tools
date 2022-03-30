package service

import (
	"context"
	"fpga-controller/internal/app/dao"
	"github.com/google/wire"
)

var testSrvSet = wire.NewSet(wire.Struct(new(testSrv), "*"))

type testSrv struct {
	testRepo *dao.testRepo
}

func (a *testSrv) Create(ctx context.Context, one *test) error {
	return a.testRepo.Create(ctx, one)
}

func (a *testSrv) Update(ctx context.Context, one *test) error {
	return a.testRepo.Update(ctx, one)
}

func (a *testSrv) Get(ctx context.Context, id uint64) (*test, error) {
	return a.testRepo.One(ctx, id)
}

func (a *testSrv) Has(ctx context.Context, id uint64) (bool, error) {
	return a.testRepo.Has(ctx, id)
}

func (a *testSrv) Delete(ctx context.Context, id uint64) error {
	return a.testRepo.Delete(ctx, id)
}

func (a *testSrv) List(ctx context.Context, req *dao.testListReq) ([]*test, error) {
	return a.testRepo.List(ctx, req)
}

func (a *testSrv) Count(ctx context.Context, req *dao.testListReq) (int64, error) {
	return a.testRepo.Count(ctx, req)
}