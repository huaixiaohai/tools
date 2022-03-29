package model

import (
	"fpga-controller/internal/app/schema"
	"time"
)

type http struct {
	Model
}

func httpFrom(one *schema.http) *http {
	if one == nil {
		return nil
	}
	return &http{
		Model: Model{
			ID: one.ID,
		},
	}

}

func httpTo(one *http) *schema.http {
	if one  == nil {
		return nil
	}
	return &schema.http{
		ID: one.ID,
	}
}
