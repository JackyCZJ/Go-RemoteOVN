package ovn

import (
	_ "apiserver/config"
	"apiserver/handler"
	"encoding/json"
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var req handler.Response

func init() {
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
	log.InitWithConfig(&passLagerCfg)
}

func TestLSAdd(t *testing.T) {


//	ovndbapi, err = goovn.NewClient(&goovn.Config{Addr: "tcp://10.1.2.82:2333"})

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/testLSAdd/:name",LSAdd)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/testLSAdd/test3", nil)
	router.ServeHTTP(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	//获得结果，并检查
	_ = json.Unmarshal(body, &req)
	if req.Code == 10001{
		t.Log(req.Message)
		return
	}
	if req.Code == 0{
		t.Log(req.Message)
		return
	}
	if req.Code == 200200{
		t.Log(req.Message)
		return
	}
	t.Fatal(req.Message,req.Code)
}

func TestLSGet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/testLSGET/:name",LSGet)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/testLSGET/test3", nil)
	router.ServeHTTP(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	//获得结果，并检查
	_ = json.Unmarshal(body, &req)
	if req.Code == 10001{
		t.Log(req.Message)
		return
	}
	if req.Code == 0{
		t.Log(req.Message)
		return
	}
	if req.Code == 200200{
		t.Log(req.Message)
		return
	}
	t.Fatal(req.Message,req.Code)
}

func TestLSDel(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/testLSDel/:name",LSDel)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/testLSDel/test3", nil)
	router.ServeHTTP(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	//获得结果，并检查
	_ = json.Unmarshal(body, &req)
	if req.Code == 10001{
		t.Log(req.Message)
		return
	}
	if req.Code == 0{
		t.Log(req.Message)
		return
	}
	if req.Code == 200200{
		t.Log(req.Message)
		return
	}
	t.Fatal(req.Message,req.Code)
}

func TestLSList(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LSList(tt.args.c)
		})
	}
}

func TestLsExtIdsAdd(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LsExtIdsAdd(tt.args.c)
		})
	}
}

func TestLSPAdd(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LSPAdd(tt.args.c)
		})
	}
}

func TestLSPDel(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LSPDel(tt.args.c)
		})
	}
}
