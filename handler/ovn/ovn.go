package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"bytes"
	"encoding/json"
	"fmt"
	goovn "github.com/eBay/go-ovn"
	"github.com/gin-gonic/gin"
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

// Todo: package 的 init过早了，无法触发viper，待解决。
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
	if err != nil{
		panic(err)
	}
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
	UUID         string `json:"uuid"`
	Name         string		`json:"name"`
	Ports        []string `json:"ports"`
	LoadBalancer []string `json:"load_balancer"`
	ACLs         []string `json:"acls"`
	QoSRules     []string `json:"qos_rules"`
	DNSRecords   []string `json:"dns_records"`
	OtherConfig  map[string]interface{} `json:"other_config"`
	ExternalID   map[string]string	`json:"external_id"`
}

//Map[interface{}]interface{} can't transfer to json , make it to map[string]interface{}
//just make it change to struct again.
func logicalSwitchStruct(v *goovn.LogicalSwitch) LogicalSwitch {
	var l LogicalSwitch
	mapString := make(map[string]string)
	for i,v :=range v.ExternalID{
		strKey := fmt.Sprintf("%v", i)
		strValue := fmt.Sprintf("%v", v)
		mapString[strKey] = strValue
	}
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	l.ExternalID = mapString
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


//gin 测试方法，返回req
type args struct {
	method string
	arg map[string]string
}

//use to test path param func
func ginTestPathTool(todo gin.HandlerFunc,args args,req *handler.Response){
	url := ""
	testUrl:= ""
	for i,arg :=range args.arg{
		url = url+"/:"+ i
		testUrl = testUrl+"/"+arg
	}
	log.Info(url)
	log.Info(testUrl)
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET(url,todo)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(args.method, testUrl, nil)
	router.ServeHTTP(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &req)
	fmt.Print(req.Message)
}

type jsonPackage struct {
	arg map[string]string
	name string
	value string
	method string
	data  map[string]interface{}
}
//use to test json param func
func ginTestJsonTool(todo gin.HandlerFunc,param jsonPackage,req *handler.Response){
	gin.SetMode(gin.TestMode)
	url := ""
	testUrl:= ""
	for i,arg :=range param.arg{
		url = url+"/:"+ i
		testUrl = testUrl+"/"+arg
	}
	router := gin.New()
	jsonByte,_ := jsoniter.Marshal(param.data)
	w := httptest.NewRecorder()
	c := bytes.NewReader(jsonByte)
	router.PUT(url,todo)
	r := httptest.NewRequest("PUT", testUrl, c)
	router.ServeHTTP(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &req)
	fmt.Print(req.Message)
}
