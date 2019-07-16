/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"testing"
)

func TestLBAdd(t *testing.T) {
	jp := jsonPackage{
		arg: map[string]string{
			"name": "lb1",
		},
		data: map[string]interface{}{
			"vipPort":  "192.168.0.19:80",
			"protocol": "tcp",
			"addrs": []string{
				"10.0.0.11:80", "10.0.0.12:80",
			},
		},
	}
	ginTestJsonTool(LBAdd, jp, &req)
	switch req.Code {
	case 0:
		t.Log(req.Message)
		t.Log(req.Data)
	case 10001:
		t.Fatal(req.Message)
	case 200200:
		t.Fatal(req.Message)
	default:
		t.Error(req.Message)
	}
}

func TestLBGet(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"name": "lb1",
		},
	}
	ginTestPathTool(LBGet, ar, &req)
	switch req.Code {
	case 0:
		t.Log(req.Message)
		t.Log(req.Data)
	case 10001:
		t.Fatal(req.Message)
	case 200200:
		t.Fatal(req.Message)
	default:
		t.Error(req.Message)
	}
}

func TestLBUpdate(t *testing.T) {
	jp := jsonPackage{
		arg: map[string]string{
			"name": "lb1",
		},
		data: map[string]interface{}{
			"vipPort":  "192.168.0.19:80",
			"protocol": "tcp",
			"addrs": []string{
				"10.0.0.11:80", "10.0.0.12:80", "10.0.0.13:80",
			},
		},
	}
	ginTestJsonTool(LBUpdate, jp, &req)
	switch req.Code {
	case 0:
		t.Log(req.Message)
		t.Log(req.Data)
	case 10001:
		t.Fatal(req.Message)
	case 200200:
		t.Fatal(req.Message)
	default:
		t.Error(req.Message)
	}
}

func TestLBDel(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"name": "lb1",
		},
	}
	ginTestPathTool(LBDel, ar, &req)
	switch req.Code {
	case 0:
		t.Log(req.Message)
	case 10001:
		t.Fatal(req.Message)
	case 200200:
		t.Fatal(req.Message)
	default:
		t.Error(req.Message)
	}
}
