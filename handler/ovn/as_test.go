package ovn

import (
	"testing"
)

func TestASAdd(t *testing.T) {
	js := jsonPackage{
		arg: map[string]string{
			"name": "asTest",
		},
		data: map[string]interface{}{
			"name": "asTest",
			"Addresses": []string{
				"127.0.0.1",
			},
			"ExternalIds": nil,
		},
	}
	ginTestJsonTool(ASAdd, js, &req)
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


func TestASGet(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "asTest"
	arg := args{
		arg: ar,
	}
	ginTestPathTool(ASGet, arg, &req)
	switch req.Code {
	case 0:
		t.Log(req.Data)
	case 10001:
		t.Fatal(req.Message)
	case 200200:
		t.Fatal(req.Message)
	default:
		t.Error(req.Message)
	}
}

func TestASList(t *testing.T) {
	arg := args{
		arg: nil,
	}
	ginTestPathTool(ASList, arg, &req)
	switch req.Code {
	case 0:
		t.Log(req.Data)
	case 10001:
		t.Fatal(req.Message)
	case 200200:
		t.Fatal(req.Message)
	default:
		t.Error(req.Message)
	}

}

func TestASUpdate(t *testing.T) {
	js := jsonPackage{
		arg: map[string]string{
			"name": "asTest",
		},
		data: map[string]interface{}{
			"name": "asTest",
			"Addresses": []string{
				"127.0.0.1",
				"233.233.233,233",
			},
			"ExternalIds": nil,
		},
	}
	ginTestJsonTool(ASUpdate, js, &req)
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

func TestASDel(t *testing.T) {
	ar := make(map[string]string)
	ar["name"] = "asTest"
	arg := args{
		arg: ar,
	}
	ginTestPathTool(ASDel, arg, &req)
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
