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

var ovndbapi = newClient()

const (
	OVS_RUNDIR   = "/var/run/openvswitch"
	OVNNB_SOCKET = "ovnnb_db.sock"
)

func newClient() goovn.Client{
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
	viper.SetDefault("ovn.remoteurl", "tcp://10.1.2.11:2333")
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
	client, err := goovn.NewClient(&goovn.Config{Addr: url})
	if err != nil {
		panic(err)
	}
	return client
}

//Map[interface{}]interface{} can't transfer to json , make it to map[string]interface{}
//just make it change to struct again.
func logicalSwitchStruct(v *goovn.LogicalSwitch) LogicalSwitch {
	var l LogicalSwitch
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	l.ExternalID = MapInterfaceToMapString(v.ExternalID)
	if err != nil {
		log.Fatal("struct unmarshal error :%v", err)
	}
	return l
}

func ACLStruct(v *goovn.ACL) AclRequest {
	var l AclRequest
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	l.ExternalIds = MapInterfaceToMapString(v.ExternalID)
	if err != nil {
		log.Fatal("struct unmarshal error :%v", err)
	}
	return l
}

func ASStruct(v *goovn.AddressSet) ASRequest {
	var l ASRequest
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	l.ExternalID = MapInterfaceToMapString(v.ExternalID)
	if err != nil {
		log.Fatal("struct unmarshal error :%v", err)
	}
	return l
}

func LRStruct(v *goovn.LogicalRouter) (l LogicalRouter) {
	str, _ := jsoniter.Marshal(v)
	err := jsoniter.Unmarshal(str, &l)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		l.ExternalID = MapInterfaceToMapString(v.ExternalID)
	}()
	go func() {
		defer wg.Done()
		l.Options = MapInterfaceToMapString(v.Options)
	}()
	wg.Wait()
	if err != nil {
		log.Fatal("struct unmarshal error :%v", err)
	}
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

func MapInterfaceToMapint(m map[interface{}]interface{}) map[string]int {
	mapString := make(map[string]int)
	for i, v := range m {
		strKey := fmt.Sprintf("%v", i)
		strValue := int(v.(float64))
		mapString[strKey] = strValue
	}
	return mapString
}
