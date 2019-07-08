package ovn

import (
	"testing"
)

func TestLRGet(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "LrTest1"
	arg := args{
		arg:    ar,
	}
	ginTestPathTool(LRGet,arg,&req)
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

func TestLRAdd(t *testing.T) {
	param := make(map[string]interface{})
	param["external_id"] = map[string]string{"a": "b"}
	jp := jsonPackage{
		arg: map[string]string{
			"name":"LrTest1",
		},
		method: "POST",
		data:   param,
	}
	ginTestJsonTool(LRAdd,jp,&req)
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

func TestLRDel(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "LrTest1"
	arg := args{
		arg:    ar,
	}
	ginTestPathTool(LRDel,arg,&req)
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
		arg: map[string]string{

		},
	}
	ginTestPathTool(ACLList,ar,&req)
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
