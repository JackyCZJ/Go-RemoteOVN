package ovn

import (
	_ "apiserver/config"
	"apiserver/handler"
	"fmt"
	"testing"
)

var req handler.Response

func TestLSAdd(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "test1"
	arg := args{
		arg: ar,
	}
	ginTestPathTool(LSAdd, arg, &req)
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

func TestLSGet(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "test1"
	arg := args{
		arg: ar,
	}
	ginTestPathTool(LSGet, arg, &req)
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

func TestLSDel(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "test1"
	arg := args{
		arg: ar,
	}
	ginTestPathTool(LSDel, arg, &req)
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

func TestLsExtIds(t *testing.T) {
	param := make(map[string]interface{})
	param["external_id"] = map[string]string{"a": "b"}
	jp := jsonPackage{
		arg: map[string]string{
			"name": "test2",
		},
		data: param,
	}
	ginTestJsonTool(LsExtIdsAdd, jp, &req)
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
	ginTestJsonTool(LsExtIdsDel, jp, &req)
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

//just execute command
func TestEverything(t *testing.T) {
	cmd, _ := ovndbapi.LRSRDel("LrTest1", "10.0.1.1/24")
	err := ovndbapi.Execute(cmd)

	//cmd, _ := ovndbapi.LSExtIdsDel("test2", map[string]string{"a": "b"})
	//err := ovndbapi.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("PASS")
}

func TestLSList(t *testing.T) {
	ar := make(map[string]string)
	arg := args{
		arg: ar,
	}
	ginTestPathTool(LSList, arg, &req)
	switch req.Code {
	case 0:
		fmt.Println(req.Data)
		t.Log(req.Message)
	case 10001:
		t.Fatal(req.Message)
	case 200200:
		t.Fatal(req.Message)
	default:
		t.Error(req.Message)
	}
}
