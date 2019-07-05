package ovn

import (
	"fmt"
	"testing"
)


func TestACLAdd(t *testing.T) {
	js := jsonPackage{
		arg: map[string]string{
			"name":"test2",
		},
		method:"GET",
		data: map[string]interface{}{
			"ls": "test2",
			"direct":"to-lport",
			"match": "outport == \"96d44061-1823-428b-a7ce-f473d10eb3d0\" && ip && ip.dst == 10.97.183.61",
			"action":"drop",
			"priority":1001,
			"external_ids":nil,
			"logflag":true,
			"meter":"",
		},
	}
	ginTestJsonTool(ACLAdd,js,&req)
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

func TestACLDel(t *testing.T) {
	js := jsonPackage{
		arg: map[string]string{
			"name":"test2",
		},
		method:"GET",
		data: map[string]interface{}{
			"ls": "test2",
			"direct":"to-lport",
			"match": "outport == \"96d44061-1823-428b-a7ce-f473d10eb3d0\" && ip && ip.dst == 10.97.183.61",
			"action":"drop",
			"priority":1001,
			"external_ids":nil,
			"logflag":true,
			"meter":"",
		},
	}
	ginTestJsonTool(ACLDel,js,&req)
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

func TestACLList(t *testing.T) {
	ar := args{
		method:"GET",
		arg: map[string]string{
			"name": "test2",
		},
	}
	ginTestPathTool(ACLList,ar,&req)
	switch req.Code {
	case 0:
		t.Log(req.Message)
		fmt.Print(req.Data)
	case 10001:
		t.Fatal(req.Message)
	case 200200:
		t.Fatal(req.Message)
	default:
		t.Error(req.Message)
	}
}
