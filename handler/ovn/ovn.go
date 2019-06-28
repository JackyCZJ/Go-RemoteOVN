package ovn

import (
	goovn "github.com/eBay/go-ovn"
	"github.com/spf13/viper"
)

var ovndbapi goovn.Client

const (
	OVS_RUNDIR   = "/var/run/openvswitch"
	OVNNB_SOCKET = "ovnnb_db.sock"
)

func init() {
	var err error
	var url string
	url = viper.GetString("ovn.REMOTEURL")
	if viper.GetString("ovn.runmode")=="local"{
		var ovs_rundir = viper.GetString("ovn.OVS_RUNDIR")
		if ovs_rundir == "" {
			ovs_rundir = OVS_RUNDIR
		}
		url = "unix:"+ovs_rundir+"/"+OVNNB_SOCKET
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
	Ls        string `json:"ls"`
	Lsp       string `json:"lsp"`
	addresses string `json:"addresses"`
	security  string `json:"security"`
}

//Address Set struct
type ASRequest struct {
	Name        string            `json:"name"`
	Addrs       string            `json:"addrs"`
	ExternalIds map[string]string `json:"external_ids"`
}

//Logical Router struct
type LRRequest struct {
	Name string `json:"name"`
}

//Logical Router Port struct
type LRPRequest struct {
	Lr          string            `json:"lr"`
	Lrp         string            `json:"lrp"`
	Mac         string            `json:"mac"`
	Network     []string          `json:"network"`
	Peer        string            `json:"peer"`
	ExternalIds map[string]string `json:"external_ids"`
}

//Logical Router And Logical Bridge	operate struct
type LRLBRequest struct {
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


//type LogicalSwitch struct {
//	UUID         string
//	Name         string
//	Ports        []string
//	LoadBalancer []string
//	ACLs         []string
//	QoSRules     []string
//	DNSRecords   []string
//	OtherConfig  map[interface{}]interface{}
//	ExternalID   map[interface{}]interface{}
//}