/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestQoSAdd(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QoSAdd(tt.args.c)
		})
	}
}

func TestQoSDel(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QoSDel(tt.args.c)
		})
	}
}

func TestQoSList(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QoSList(tt.args.c)
		})
	}
}
