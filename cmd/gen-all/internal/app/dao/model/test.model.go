package model

import (
	"fpga-controller/internal/app/schema"
)

type test struct {
	Model
}

func testFrom(one *schema.test) *test {
	if one == nil {
		return nil
	}
	return &test{
		Model: Model{
			ID: one.ID,
		},
	}

}

func testTo(one *test) *schema.test {
	if one  == nil {
		return nil
	}
	return &schema.test{
		ID: one.ID,
	}
}
