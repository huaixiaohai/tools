package model

import (
	"fpga-controller/internal/app/schema"
	"time"
)

type PhysicalHost struct {
	Model
}

func PhysicalHostFrom(one *schema.PhysicalHost) *PhysicalHost {
	if one == nil {
		return nil
	}
	return &PhysicalHost{
		Model: Model{
			ID: one.ID,
		},
	}

}

func PhysicalHostTo(one *PhysicalHost) *schema.PhysicalHost {
	if one  == nil {
		return nil
	}
	return &schema.PhysicalHost{
		ID: one.ID,
	}
}
