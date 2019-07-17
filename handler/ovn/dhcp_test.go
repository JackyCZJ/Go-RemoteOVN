/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"github.com/gin-gonic/gin/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDHCPOptionsAdd(t *testing.T) {
	jp := jsonPackage{
		arg: map[string]string{},
		data: map[string]interface{}{
			"CIDR": "192.168.0.0/24",
			"OPTIONS": map[string]string{
				"server_id":  "192.168.1.1",
				"server_mac": "54:54:54:54:54:54",
				"lease_time": "6000",
			},
		},
	}
	ginTestJsonTool(DHCPOptionsAdd, jp, &req)
	assert.Equal(t, "OK", req.Message)
}

//todo: It Fail, seems need to modify ovs db api.
func TestDHCPOptionSet(t *testing.T) {
	jp := jsonPackage{
		arg: map[string]string{},
		data: map[string]interface{}{
			"CIDR": "192.168.0.0/24",
			"OPTIONS": map[string]string{
				"server_id":  "192.168.1.2",
				"server_mac": "54:54:54:54:54:55",
				"lease_time": "6000",
			},
		},
	}
	ginTestJsonTool(DHCPOptionSet, jp, &req)
	assert.Equal(t, "OK", req.Message)
}

var uuid string

func TestDHCPOptionsList(t *testing.T) {
	ginTestPathTool(DHCPOptionsList, args{}, &req)
	assert.Equal(t, "OK", req.Message)
	//uuid = req.Data.(DHCPOptions).UUID
	switch req.Data.(type) {
	case []interface{}:
		for _, v := range req.Data.([]interface{}) {
			switch v.(type) {
			case interface{}:
				str, _ := json.Marshal(v)
				dhcp := DHCPOptions{}
				_ = jsoniter.Unmarshal(str, &dhcp)
				uuid = dhcp.UUID
			default:
				t.Fatal("Fail to GET uuid")
			}
		}
	default:
		t.Fatal("Fail to GET uuid")
	}

}

func TestLSPSetDHCPv4Options(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"name": "br-int",
			"uuid": uuid,
		},
	}
	ginTestPathTool(LSPSetDHCPv4Options, ar, &req)
	assert.Equal(t, "OK", req.Message)

}

func TestLSPGetDHCPv4Options(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"name": "br-int",
		},
	}
	ginTestPathTool(LSPGetDHCPv4Options, ar, &req)
	switch req.Data.(type) {
	case []interface{}:
		for _, v := range req.Data.([]interface{}) {
			switch v.(type) {
			case interface{}:
				str, _ := json.Marshal(v)
				dhcp := DHCPOptions{}
				_ = jsoniter.Unmarshal(str, &dhcp)
				assert.Equal(t, uuid, dhcp.UUID)
			default:
				t.Fatal("Fail to GET uuid")
			}
		}
	case interface{}:
		{
			v := req.Data.(interface{})
			str, _ := json.Marshal(v)
			dhcp := DHCPOptions{}
			_ = jsoniter.Unmarshal(str, &dhcp)
			assert.Equal(t, uuid, dhcp.UUID)
		}
	default:
		t.Fatal("Fail to GET uuid")
	}

}

func TestLSPSetDHCPv6Options(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"name": "br-int1",
			"uuid": "5a5e48fc-0698-405d-a3ba-aae68fa85e6f",
		},
	}
	ginTestPathTool(LSPSetDHCPv6Options, ar, &req)
	assert.Equal(t, "OK", req.Message)
}

func TestLSPGetDHCPv6Options(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"name": "br-int1",
		},
	}
	ginTestPathTool(LSPGetDHCPv6Options, ar, &req)
	assert.Equal(t, "OK", req.Message)

}

func TestDHCPOptionsDel(t *testing.T) {
	ar := args{
		arg: map[string]string{
			"uuid": uuid,
		},
	}
	ginTestPathTool(DHCPOptionsDel, ar, &req)
	assert.Equal(t, "OK", req.Message)

}
