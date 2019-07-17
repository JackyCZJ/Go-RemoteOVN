/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	goovn "github.com/jackyczj/go-ovn"
	jsoniter "github.com/json-iterator/go"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
)

var ovndbapi goovn.Client

const (
	OVS_RUNDIR   = "/var/run/openvswitch"
	OVNNB_SOCKET = "ovnnb_db.sock"
)

func init() {
	var err error
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	err = log.InitWithConfig(&passLagerCfg)
	if err != nil {
		panic(err)
	}
	viper.SetDefault("ovn.remoteurl", "tcp://10.1.2.82:2333")
	//if err := config.Init(""); err != nil {
	//	panic(err)
	//}
	var url string
	url = viper.GetString("ovn.remoteurl")
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
	Priority    int               `json:"priority"`
	ExternalIds map[string]string `json:"external_ids"`
	Logflag     bool              `json:"logflag"`
	Meter       string            `json:"meter"`
}

type ACL struct {
	UUID       string
	Action     string
	Direction  string
	Match      string
	Priority   int
	Log        bool
	ExternalID map[string]string
}

//Logical switch struct
type LsRequest struct {
	Ls string `json:"ls"`
}

//Address Set struct
type ASRequest struct {
	Name       string            `json:"name"`
	Addresses  []string          `json:"addresses"`
	ExternalID map[string]string `json:"external_id"`
	UUID       string            `json:"uuid"`
}

//Logical Router struct
type LRRequest struct {
	Name       string            `json:"name"`
	ExternalID map[string]string `json:"external_id"`
}

//Logical Router Port struct
type LRPRequest struct {
	Lrp         string            `json:"lrp"`
	Mac         string            `json:"mac"`
	Network     []string          `json:"network"`
	Peer        string            `json:"peer"`
	ExternalIds map[string]string `json:"external_ids"`
}

type LogicalRouterPort struct {
	UUID           string
	Name           string
	GatewayChassis []string
	Networks       []string
	MAC            string
	Enabled        bool
	IPv6RAConfigs  map[string]string
	Options        map[string]string
	Peer           string
	ExternalID     map[string]string
}

//Logical Bridge struct
type LBRequest struct {
	VipPort  string   `json:"vipPort"`
	Protocol string   `json:"protocol"`
	Addrs    []string `json:"addrs"`
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
	UUID         string            `json:"uuid"`
	Name         string            `json:"name"`
	Ports        []string          `json:"ports"`
	LoadBalancer []string          `json:"load_balancer"`
	ACLs         []string          `json:"acls"`
	QoSRules     []string          `json:"qos_rules"`
	DNSRecords   []string          `json:"dns_records"`
	OtherConfig  map[string]string `json:"other_config"`
	ExternalID   map[string]string `json:"external_id"`
}

type LogicalSwitchPort struct {
	UUID          string            `json:"uuid"`
	Name          string            `json:"name"`
	Type          string            `json:"type"`
	Options       map[string]string `json:"options"`
	Addresses     []string          `json:"addresses"`
	PortSecurity  []string
	DHCPv4Options string
	DHCPv6Options string
	ExternalID    map[string]string `json:"external_id"`
}

type LogicalRouter struct {
	UUID    string
	Name    string
	Enabled bool

	Ports        []string
	StaticRoutes []string
	NAT          []string
	LoadBalancer []string

	Options    map[string]string
	ExternalID map[string]string
}

type AS struct {
	Addresses []string `json:"addresses"`
}

type Security struct {
	Security []string `json:"security"`
}

type StaticRouter struct {
	UUID       string            `json:"uuid"`
	IPPrefix   string            `json:"ip_prefix"`
	Nexthop    string            `json:"nexthop"`
	OutputPort []string          `json:"output_port"`
	Policy     []string          `json:"policy"`
	ExternalID map[string]string `json:"external_id"`
}

type DHCPOptions struct {
	UUID       string
	CIDR       string
	Options    map[string]string
	ExternalID map[string]string
}

//Map[interface{}]interface{} can't transfer to json , make it to map[string]interface{}
//just make it change to struct again.
func logicalSwitchStruct(v *goovn.LogicalSwitch) LogicalSwitch {
	var l LogicalSwitch
	mapString := make(map[string]string)
	for i, v := range v.ExternalID {
		strKey := fmt.Sprintf("%v", i)
		strValue := fmt.Sprintf("%v", v)
		mapString[strKey] = strValue
	}
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	l.ExternalID = mapString
	if err != nil {
		log.Fatal("struct unmarshal error :%v", err)
	}
	return l
}

func ACLStruct(v *goovn.ACL) ACL {
	var l ACL
	mapString := make(map[string]string)
	for i, v := range v.ExternalID {
		strKey := fmt.Sprintf("%v", i)
		strValue := fmt.Sprintf("%v", v)
		mapString[strKey] = strValue
	}
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	l.ExternalID = mapString
	if err != nil {
		log.Fatal("struct unmarshal error :%v", err)
	}
	return l
}

func ASStruct(v *goovn.AddressSet) ASRequest {
	var l ASRequest
	mapString := make(map[string]string)
	for i, v := range v.ExternalID {
		strKey := fmt.Sprintf("%v", i)
		strValue := fmt.Sprintf("%v", v)
		mapString[strKey] = strValue
	}
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	l.ExternalID = mapString
	if err != nil {
		log.Fatal("struct unmarshal error :%v", err)
	}
	return l
}

func LRStruct(v *goovn.LogicalRouter) (l LogicalRouter) {
	mapString := make(map[string]string)
	optString := make(map[string]string)
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i, v := range v.ExternalID {
			strKey := fmt.Sprintf("%v", i)
			strValue := fmt.Sprintf("%v", v)
			mapString[strKey] = strValue
		}
	}()
	go func() {
		defer wg.Done()
		for i, v := range v.Options {
			optionKey := fmt.Sprintf("%v", i)
			optValue := fmt.Sprintf("%v", v)
			optString[optionKey] = optValue
		}
	}()
	wg.Wait()
	if err != nil {
		log.Fatal("struct unmarshal error :%v", err)
	}
	l.ExternalID = mapString
	l.Options = optString
	return
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

type args struct {
	arg map[string]string
}

//use to test path param func
func ginTestPathTool(todo gin.HandlerFunc, args args, req *handler.Response) {
	url := ""
	testUrl := ""
	if len(args.arg) != 0 {
		for i, arg := range args.arg {
			url = url + "/:" + i
			testUrl = testUrl + "/" + arg
		}
	} else {
		url = "/"
		testUrl = "/"
	}
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET(url, todo)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", testUrl, nil)
	router.ServeHTTP(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &req)
	fmt.Print(req.Message)
}

//method do nothing.
type jsonPackage struct {
	arg  map[string]string
	data map[string]interface{}
}

//use to test json param func
func ginTestJsonTool(todo gin.HandlerFunc, param jsonPackage, req *handler.Response) {
	gin.SetMode(gin.TestMode)
	url := ""
	testUrl := ""
	if len(param.arg) == 0 {
		url = "/"
		testUrl = "/"
	} else {
		for i, arg := range param.arg {
			url = url + "/:" + i
			testUrl = testUrl + "/" + arg
		}
	}
	router := gin.New()
	jsonByte, _ := jsoniter.Marshal(param.data)
	w := httptest.NewRecorder()
	c := bytes.NewReader(jsonByte)
	router.PUT(url, todo)
	r := httptest.NewRequest("PUT", testUrl, c)
	router.ServeHTTP(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &req)
	//	fmt.Print(req.Message)
}

func MapInterfaceToMapString(m map[interface{}]interface{}) map[string]string {
	mapString := make(map[string]string)
	for i, v := range m {
		strKey := fmt.Sprintf("%v", i)
		strValue := fmt.Sprintf("%v", v)
		mapString[strKey] = strValue
	}
	return mapString
}
