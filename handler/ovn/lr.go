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

//Logical Router struct
type LogicalRouter struct {
	UUID    string
	Name    string 						`json:"name"`
	Enabled bool

	Ports        []string
	StaticRoutes []string
	NAT          []string
	LoadBalancer []string

	Options    map[string]string
	ExternalID map[string]string 		`json:"external_id"`
}

type LogicalRouterPort struct {
	UUID           string   			`json:"uuid"`
	Name           string   			`json:"name"`
	GatewayChassis []string 			`json:"gateway_chassis"`
	Networks       []string 			`json:"networks"`
	MAC            string   			`json:"mac"`
	Enabled        bool     			`json:"enabled"`
	IPv6RAConfigs  map[string]string
	Options        map[string]string 	`json:"options"`
	Peer           string            	`json:"peer"`
	ExternalID     map[string]string 	`json:"external_id"`
}
//Logical Router Port struct
type StaticRouter struct {
	UUID       string            `json:"uuid"`
	IPPrefix   string            `json:"ip_prefix"`
	Nexthop    string            `json:"nexthop"`
	OutputPort []string          `json:"output_port"`
	Policy     []string          `json:"policy"`
	ExternalID map[string]string `json:"external_id"`
}

//Logical Router Get
func LRGet(c *gin.Context) {
	lr, err := ovndbapi.LRGet(c.Param("name"))
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRGet)
		return
	}
	var l LogicalRouter

	for _, v := range lr {
		log.Infof("%s", v.StaticRoutes)
		l = LRStruct(v)
	}
	handler.SendResponse(c, nil, l)
}

//Logical Router Add
func LRAdd(c *gin.Context) {
	var lr LogicalRouter
	l := c.Param("name")
	err := c.BindJSON(&lr)
	if err != nil {
		log.Info("LR add with nil external id")
	}
	lr.Name = l
	if len(lr.ExternalID) == 0 {
		lr.ExternalID = nil
	}
	cmd, err := ovndbapi.LRAdd(lr.Name, lr.ExternalID)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	log.Infof("Logical Router Add or Update: %s", l)
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
	log.Infof("Logical Router Delete: %s", l)
	handler.SendResponse(c, nil, nil)
}

//Logical Router List
func LRList(c *gin.Context) {
	lrList, err := ovndbapi.LRList()
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRList)
		return
	}
	var rList []LogicalRouter
	for _, v := range lrList {
		log.Infof("%s", v.StaticRoutes)
		rList = append(rList, LRStruct(v))
	}
	handler.SendResponse(c, nil, rList)
}

var LRP LogicalRouterPort

//Logical Router Port Add
func LRPAdd(c *gin.Context) {
	err := c.BindJSON(&LRP)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	LR := c.Param("name")
	LRP.Name = c.Param("port")
	a := LRP.ExternalID
	fmt.Printf("%v \n", a)
	cmd, err := ovndbapi.LRPAdd(LR, LRP.Name, LRP.MAC, LRP.Networks, LRP.Peer, LRP.ExternalID)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	log.Infof("Logical Router %s Add or Update Port : %s ", LR, LRP.Name)
	handler.SendResponse(c, nil, nil)
}

func LRPDel(c *gin.Context) {
	LR := c.Param("name")
	LP := c.Param("port")
	cmd, err := ovndbapi.LRPDel(LR, LP)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	log.Infof("Logical Router %s  Delete Port %s", LR, LP)
	handler.SendResponse(c, nil, nil)
}

func LRPList(c *gin.Context) {
	LR := c.Param("name")
	cmd, err := ovndbapi.LRPList(LR)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	var LRPs []LogicalRouterPort
	var LRP LogicalRouterPort
	//	ch := make(chan map[string]string)

	for _, v := range cmd {
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
			LRP.ExternalID = mapString
		}()
		go func() {
			defer wg.Done()
			mapString := make(map[string]string)
			for i, v := range v.IPv6RAConfigs {
				strKey := fmt.Sprintf("%v", i)
				strValue := fmt.Sprintf("%v", v)
				mapString[strKey] = strValue
			}
			LRP.IPv6RAConfigs = mapString
		}()
		go func() {
			defer wg.Done()
			mapString := make(map[string]string)
			for i, v := range v.Options {
				strKey := fmt.Sprintf("%v", i)
				strValue := fmt.Sprintf("%v", v)
				mapString[strKey] = strValue
			}
			LRP.Options = mapString
		}()
		wg.Wait()

		LRPs = append(LRPs, LRP)
	}
	handler.SendResponse(c, nil, LRPs)
}

var sr StaticRouter

func LRSRAdd(c *gin.Context) {
	lr := c.Param("name")
	err := c.BindJSON(&sr)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	cmd, err := ovndbapi.LRSRAdd(lr, sr.IPPrefix, sr.Nexthop, sr.OutputPort, sr.Policy, sr.ExternalID)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	log.Infof("Logical Router :%s ,Add Static Router:  %s", lr, sr.IPPrefix)
	handler.SendResponse(c, nil, nil)
}

func LRSRDel(c *gin.Context) {
	lr := c.Param("name")
	err := c.BindJSON(&sr)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ip := sr.IPPrefix
	cmd, err := ovndbapi.LRSRDel(lr, ip)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	log.Infof("Logical Router :%s ,Delete Static Router:  %s", lr, ip)
	handler.SendResponse(c, nil, nil)
}

func LRSRList(c *gin.Context) {
	lr := c.Param("name")
	cmd, err := ovndbapi.LRSRList(lr)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRGet)
		return
	}
	var lrsrList []StaticRouter
	var sr StaticRouter
	for _, v := range cmd {
		log.Infof("%v", v)
		str, _ := jsoniter.Marshal(v)
		err = jsoniter.Unmarshal(str, &sr)
		if err != nil {
			log.Fatal("struct unmarshal error :%v", err)
		}
		sr.ExternalID = MapInterfaceToMapString(v.ExternalID)
		lrsrList = append(lrsrList, sr)
	}
	handler.SendResponse(c, nil, lrsrList)
}

func LRLBAdd(c *gin.Context) {
	lr := c.Param("name")
	lb := c.Param("lb")
	cmd, err := ovndbapi.LRLBAdd(lr, lb)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func LRLBDel(c *gin.Context) {
	lr := c.Param("name")
	lb := c.Param("lb")
	cmd, err := ovndbapi.LRLBDel(lr, lb)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func LRLBlist(c *gin.Context) {
	lr := c.Param("name")
	cmd, err := ovndbapi.LRLBList(lr)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
	}
	pack := make(map[string]interface{})
	var packs []map[string]interface{}
	for _, v := range cmd {
		pack["UUID"] = v.UUID
		pack["Name"] = v.Name
		pack["ExternalID"] = MapInterfaceToMapString(v.ExternalID)
		packs = append(packs, pack)
	}
	fmt.Printf("%s", packs)
	handler.SendResponse(c, nil, packs)
}
