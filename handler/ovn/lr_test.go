package ovn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRAdd(t *testing.T) {
	param := make(map[string]interface{})
	param["external_id"] = map[string]string{"a": "b"}
	jp := jsonPackage{
		arg: map[string]string{
			"name": "LrTest1",
		},
		data: param,
	}
	ginTestJsonTool(LRAdd, jp, &req)
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

func TestLRList(t *testing.T) {
	ar := args{
		arg: map[string]string{},
	}
	ginTestPathTool(LRList, ar, &req)
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
func TestLRGet(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "LrTest1"
	arg := args{
		arg: ar,
	}
	ginTestPathTool(LRGet, arg, &req)
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

func TestLRPAdd(t *testing.T) {
	param := make(map[string]interface{})
	param["external_ids"] = map[string]string{"a": "b",
		"foo": "bar"}
	param["mac"] = "54:54:54:54:54:56"
	param["network"] = []string{"192.168.0.1/24"}
	param["peer"] = "lrp3"
	jp := jsonPackage{
		arg: map[string]string{
			"name": "LrTest1",
			"port": "br-int1",
		},
		data: param,
	}
	ginTestJsonTool(LRPAdd, jp, &req)
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

func TestLRPList(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"name": "LrTest1",
		},
	}
	ginTestPathTool(LRPList, ar, &req)
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

var srt = map[string]interface{}{
	"ip_prefix": "10.0.1.1/24",
	"nexthop":   "10.3.0.1",
}

func TestLRSRAdd(t *testing.T) {
	jp := jsonPackage{
		arg: map[string]string{
			"name": "LrTest1",
		},
		data: srt,
	}
	ginTestJsonTool(LRSRAdd, jp, &req)
	assert.Equal(t, req.Message, "OK")
}

func TestLRSRList(t *testing.T) {
	arg := make(map[string]string)
	arg["name"] = "LrTest1"
	ar := args{
		arg: arg,
	}
	ginTestPathTool(LRSRList, ar, &req)
	assert.Equal(t, req.Message, "OK")
}

func TestLRSRDel(t *testing.T) {
	arg := make(map[string]string)
	arg["name"] = "LrTest1"
	jp := jsonPackage{
		arg:  arg,
		data: srt,
	}
	ginTestJsonTool(LRSRDel, jp, &req)
	assert.Equal(t, req.Message, "OK")
	fmt.Printf("%s", req.Data)
}

func TestLRLBAdd(t *testing.T) {
	cmd, _ := ovndbapi.LBAdd("lb1", "192.168.0.20:80", "tcp", []string{"10.0.0.21:80", "10.0.0.22:80"})
	_ = ovndbapi.Execute(cmd)
	arg := make(map[string]string)
	arg["name"] = "LrTest1"
	arg["lb"] = "lb1"
	ar := args{
		arg: arg,
	}
	ginTestPathTool(LRLBAdd, ar, &req)
	assert.Equal(t, req.Message, "OK")
}

func TestLRLBlist(t *testing.T) {
	arg := make(map[string]string)
	arg["name"] = "LrTest1"
	ar := args{
		arg: arg,
	}
	ginTestPathTool(LRLBlist, ar, &req)
	assert.Equal(t, req.Message, "OK")
}

func TestLRLBDel(t *testing.T) {
	defer func() {
		cmd, _ := ovndbapi.LBDel("lb1")
		_ = ovndbapi.Execute(cmd)
	}()
	arg := make(map[string]string)
	arg["name"] = "LrTest1"
	arg["lb"] = "lb1"
	ar := args{
		arg: arg,
	}
	ginTestPathTool(LRLBDel, ar, &req)
	assert.Equal(t, req.Message, "OK")
}

func TestLRPDel(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"name": "LrTest1",
			"port": "br-int1",
		},
	}
	ginTestPathTool(LRDel, ar, &req)
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

func TestLRDel(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "LrTest1"
	arg := args{
		arg: ar,
	}
	ginTestPathTool(LRDel, arg, &req)
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
