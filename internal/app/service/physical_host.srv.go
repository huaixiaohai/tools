package service

import (
	"context"
	"fpga-controller/internal/app/dao"
	. "fpga-controller/internal/app/schema"
	"github.com/google/wire"
)

var PhysicalHostSrvSet = wire.NewSet(wire.Struct(new(PhysicalHostSrv), "*"))

type PhysicalHostSrv struct {
	PhysicalHostRepo *dao.PhysicalHostRepo
}

func (a *PhysicalHostSrv) Create(ctx context.Context, one *PhysicalHost) error {
	return a.PhysicalHostRepo.Create(ctx, one)
}

func (a *PhysicalHostSrv) Update(ctx context.Context, one *PhysicalHost) error {
	return a.PhysicalHostRepo.Update(ctx, one)
}

func (a *PhysicalHostSrv) Get(ctx context.Context, id uint64) (*PhysicalHost, error) {
	return a.PhysicalHostRepo.One(ctx, id)
}

func (a *PhysicalHostSrv) Has(ctx context.Context, id uint64) (bool, error) {
	return a.PhysicalHostRepo.Has(ctx, id)
}

func (a *PhysicalHostSrv) Delete(ctx context.Context, id uint64) error {
	return a.PhysicalHostRepo.Delete(ctx, id)
}

func (a *PhysicalHostSrv) List(ctx context.Context, req *dao.PhysicalHostListReq) ([]*PhysicalHost, error) {
	return a.PhysicalHostRepo.List(ctx, req)
}

func (a *PhysicalHostSrv) Count(ctx context.Context, req *dao.PhysicalHostListReq) (int64, error) {
	return a.PhysicalHostRepo.Count(ctx, req)
}