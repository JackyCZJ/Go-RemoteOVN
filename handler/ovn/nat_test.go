/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRNAT(t *testing.T) {
	cmd, _ := ovndbapi.LRAdd("LR_TEST", nil)
	_ = ovndbapi.Execute(cmd)

	jp := jsonPackage{
		arg: map[string]string{
			"name": "LR_TEST",
		},
		data: map[string]interface{}{
			"type":        "snat",
			"external_ip": "10.127.0.129",
			"logical_ip":  "10.127.24.244/25",
		},
	}
	ginTestJsonTool(LRNATAdd, jp, &req)
	assert.Equal(t, "OK", req.Message)
	ginTestPathTool(LRNATList, args{arg: map[string]string{"name": "LR_TEST"}}, &req)
	a := fmt.Sprint(req.Data)
	ginTestJsonTool(LRNATDel, jp, &req)
	assert.Equal(t, "OK", req.Message)
	ginTestPathTool(LRNATList, args{arg: map[string]string{"name": "LR_TEST"}}, &req)
	b := fmt.Sprint(req.Data)
	if a == b {
		t.Fatal("not delete")
	}
}
