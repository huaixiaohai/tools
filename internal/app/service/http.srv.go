package service

import (
	"context"
	"fpga-controller/internal/app/dao"
	. "fpga-controller/internal/app/schema"
	"github.com/google/wire"
)

var httpSrvSet = wire.NewSet(wire.Struct(new(httpSrv), "*"))

type httpSrv struct {
	httpRepo *dao.httpRepo
}

func (a *httpSrv) Create(ctx context.Context, one *http) error {
	return a.httpRepo.Create(ctx, one)
}

func (a *httpSrv) Update(ctx context.Context, one *http) error {
	return a.httpRepo.Update(ctx, one)
}

func (a *httpSrv) Get(ctx context.Context, id uint64) (*http, error) {
	return a.httpRepo.One(ctx, id)
}

func (a *httpSrv) Has(ctx context.Context, id uint64) (bool, error) {
	return a.httpRepo.Has(ctx, id)
}

func (a *httpSrv) Delete(ctx context.Context, id uint64) error {
	return a.httpRepo.Delete(ctx, id)
}

func (a *httpSrv) List(ctx context.Context, req *dao.httpListReq) ([]*http, error) {
	return a.httpRepo.List(ctx, req)
}

func (a *httpSrv) Count(ctx context.Context, req *dao.httpListReq) (int64, error) {
	return a.httpRepo.Count(ctx, req)
}