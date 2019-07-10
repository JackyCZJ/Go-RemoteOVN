package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/lexkong/log"
	"sync"
)

var lr LRRequest

//Logical Router Get
func LRGet(c *gin.Context) {
	lr ,err := ovndbapi.LRGet(c.Param("name"))
	if err != nil{
		handleOvnErr(c,err,errno.ErrLRGet)
		return
	}
	var l  LogicalRouter

	for _,v :=range lr {
		l = LRStruct(v)
	}
	handler.SendResponse(c,nil,l)
}

//Logical Router Add
func LRAdd(c *gin.Context) {
	l := c.Param("name")
	err := c.BindJSON(&lr)
	if err!=nil {
		log.Info("LR add with nil external id")
	}
	lr.Name = l
	if len(lr.ExternalID)==0{
		lr.ExternalID = nil
	}
	cmd, err := ovndbapi.LRAdd(lr.Name,lr.ExternalID)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	log.Infof("Logical Router Add or Update: %s",l)
	handler.SendResponse(c, nil, nil)
}

//Logical Router Delete
func LRDel(c *gin.Context) {
	l := c.Param("name")
	ocmd, err := ovndbapi.LRDel(l)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
		}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
	}
	log.Infof("Logical Router Delete: %s",l)
	handler.SendResponse(c, nil, nil)
}

//Logical Router List
func LRList(c *gin.Context) {
	lrList ,err := ovndbapi.LRList()
	if err != nil{
		handleOvnErr(c,err,errno.ErrLRList)
		return
	}
	var rList []LogicalRouter
	for _,v :=range lrList{
		rList = append(rList,LRStruct(v))
	}
	handler.SendResponse(c,nil,rList)
}

var LRP LRPRequest

//Logical Router Port Add
func LRPAdd(c *gin.Context){
	err := c.BindJSON(&LRP);if err != nil {
		handler.SendResponse(c,errno.ErrBind,nil)
		return
	}
	LR := c.Param("name")
	LRP.Lrp = c.Param("port")
	a := LRP.ExternalIds
	fmt.Printf("%v \n",a)
	cmd, err := ovndbapi.LRPAdd(LR, LRP.Lrp, LRP.Mac, LRP.Network, LRP.Peer, LRP.ExternalIds)
	if err !=nil{
		handleOvnErr(c,err,err)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err !=nil{
		handleOvnErr(c,err,err)
		return
	}
	log.Infof("Logical Router %s Add or Update Port : %s ",LR,LRP.Lrp)
	handler.SendResponse(c,nil,nil)
}

func LRPDel(c *gin.Context){
	LR:=c.Param("name")
	LP:=c.Param("port")
	cmd,err := ovndbapi.LRPDel(LR,LP)
	if err !=nil{
		handleOvnErr(c,err,err)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err !=nil{
		handleOvnErr(c,err,err)
		return
	}
	log.Infof("Logical Router %s  Delete Port %s",LR,LP)
	handler.SendResponse(c,nil,nil)
}


func LRPList(c *gin.Context){
	LR:=c.Param("name")
	cmd,err := ovndbapi.LRPList(LR)
	if err !=nil{
		handleOvnErr(c,err,err)
		return
	}
	var LRPs []LogicalRouterPort
	var LRP LogicalRouterPort
//	ch := make(chan map[string]string)

	for _,v :=range cmd{
		wg := sync.WaitGroup{}
		str, _ := jsoniter.Marshal(v)
		err := jsoniter.Unmarshal(str, &LRP)
		if err != nil {
			log.Fatal("struct unmarshal error :%v", err)
		}
		wg.Add(3)
		go func() {
			defer wg.Done()
			mapString := make(map[string]string)
			for i, v := range v.ExternalID {
				strKey := fmt.Sprintf("%v", i)
				strValue := fmt.Sprintf("%v", v)
				mapString[strKey] = strValue
			}
		LRP.ExternalID  = mapString
		}()
		go func() {
			defer wg.Done()
			mapString := make(map[string]string)
			for i, v := range v.IPv6RAConfigs {
				strKey := fmt.Sprintf("%v", i)
				strValue := fmt.Sprintf("%v", v)
				mapString[strKey] = strValue
			}
			LRP.IPv6RAConfigs =  mapString
		}()
		go func() {
			defer wg.Done()
			mapString := make(map[string]string)
			for i, v := range v.Options  {
				strKey := fmt.Sprintf("%v", i)
				strValue := fmt.Sprintf("%v", v)
				mapString[strKey] = strValue
			}
			LRP.Options  =  mapString
		}()
		wg.Wait()

		LRPs = append(LRPs,LRP)
	}
	handler.SendResponse(c,nil,LRPs)
}