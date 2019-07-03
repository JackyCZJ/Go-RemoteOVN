package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	goovn "github.com/eBay/go-ovn"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"sync"
)

var ovndbapi goovn.Client

const (
	OVS_RUNDIR   = "/var/run/openvswitch"
	OVNNB_SOCKET = "ovnnb_db.sock"
)

// Todo: package 的 init过早了，无法触发viper，待解决。
func init() {
	var err error
	var url string
	//y.SetDefault("ovn.remoteurl","tcp://10.1.2.82:2333")
	viper.SetDefault("ovn.remoteurl", "tcp://10.1.2.82:2333")
	url = viper.GetString("ovn.remoteurl")
	//url = viper.GetString("ovn.remoteurl")
	if viper.GetString("ovn.runmode") == "local" {
		var ovs_rundir = viper.GetString("ovn.OVS_RUNDIR")
		if ovs_rundir == "" {
			ovs_rundir = OVS_RUNDIR
		}
		url = "unix:" + ovs_rundir + "/" + OVNNB_SOCKET
	}
	ovndbapi, err = goovn.NewClient(&goovn.Config{Addr: url})
	if err != nil {
		panic(err)
	}
}

//ACL Request struct
type AclRequest struct {
	Ls          string            `json:"ls"`
	Direct      string            `json:"direct"`
	Match       string            `json:"match"`
	Action      string            `json:"action"`
	Priority    string            `json:"priority"`
	ExternalIds map[string]string `json:"external_ids"`
	Logflag     bool              `json:"logflag"`
	Meter       string            `json:"meter"`
}

//Logical switch struct
type LsRequest struct {
	Ls string `json:"ls"`
}

//Logical switch port struct
type LspRequest struct {
	LsRequest
	Lsp       string `json:"lsp"`
	addresses string `json:"addresses"`
	security  string `json:"security"`
}

//Address Set struct
type ASRequest struct {
	sync.Mutex
	Name        string            `json:"name"`
	Addrs       string            `json:"addrs"`
	ExternalIds map[string]string `json:"external_ids"`
}

//Logical Router struct
type LRRequest struct {
	sync.Mutex
	Name string `json:"name"`
}

//Logical Router Port struct
type LRPRequest struct {
	sync.Mutex
	Lr          string            `json:"lr"`
	Lrp         string            `json:"lrp"`
	Mac         string            `json:"mac"`
	Network     []string          `json:"network"`
	Peer        string            `json:"peer"`
	ExternalIds map[string]string `json:"external_ids"`
}

//Logical Router And Logical Bridge	operate struct
type LRLBRequest struct {
	sync.Mutex
	Lr string `json:"lr"`
	Lb string `json:"lb"`
}

//Logical Bridge struct
type LBRequest struct {
	Lb       string   `json:"lb"`
	VipPort  string   `json:"vipPort"`
	Protocol string   `json:"protocol"`
	addrs    []string `json:"addrs"`
}

//dhcp4_options on lsp
type LspDHCPv4 struct {
}

//dhcp6_options on lsp
type LspDHCPv6 struct {
}

//QOS struct
type QoSRequest struct {
}

type CreateResponse struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Action string `json:"action"`
	Status string `json:"status"`
}

type LogicalSwitch struct {
	UUID         string
	Name         string
	Ports        []string
	LoadBalancer []string
	ACLs         []string
	QoSRules     []string
	DNSRecords   []string
	OtherConfig  map[string]interface{}
	ExternalID   map[string]string
}

//Map[interface{}]interface{} can't transfer to json , make it to map[string]interface{}
//just make it change to struct again.
func logicalSwitchStruct(v interface{}) LogicalSwitch {
	var l LogicalSwitch
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
	return l
}

//Only use to handle OVN api error!
func handleOvnErr(c *gin.Context, err error, errn error) {
	erro := &errno.Errno{
		Message: err.Error(),
		Code:    20200,
	}
	if viper.GetString("runmode") == "debug" {
		handler.SendResponse(c, erro, nil)
		log.Errorf(erro, "err executing command:")
		return
	}
	log.Errorf(erro, "err executing command:")
	handler.SendResponse(c, errn, nil)
	return
}
