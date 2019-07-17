/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jp = jsonPackage{
	arg: map[string]string{
		"name": "ls1",
	},
	data: map[string]interface{}{
		"direction": "to-lport",
		"priority":  1001,
		"match":     `inport=="lp3"`,
		"bandwidth": map[string]int{
			"rate": 1234, "burst": 12345,
		},
	},
}

func TestQoSAdd(t *testing.T) {

	ginTestJsonTool(QoSAdd, jp, &req)
	fmt.Print(req)
	assert.Equal(t, "OK", req.Message)
}

func TestQoSList(t *testing.T) {
	ginTestPathTool(QoSList, args{arg: map[string]string{"name": "ls1"}}, &req)
	fmt.Print(req.Data)
}

func TestQoSDel(t *testing.T) {
	ginTestJsonTool(QoSDel, jp, &req)
	assert.Equal(t, "OK", req.Message)

}
