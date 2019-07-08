package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
)

var lr LRRequest

func LRGet(c *gin.Context) {
	defer  ovndbapi.RUnlock()
	ovndbapi.RLock()
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

func LRAdd(c *gin.Context) {
	l := c.Param("name")
	err := c.BindJSON(&lr);if err!=nil{
		return
	}
	lr.Name = l
	cmd, err := ovndbapi.LRAdd(lr.Name,lr.ExternalID)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	defer ovndbapi.Unlock()
	ovndbapi.Lock()
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRAdd)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func LRDel(c *gin.Context) {
	lr.Name = c.Param("name")
	defer ovndbapi.Unlock()
	ovndbapi.Lock()
	ocmd, err := ovndbapi.LRDel(lr.Name)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
		}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func LRList(c *gin.Context) {
	defer  ovndbapi.RUnlock()
	ovndbapi.RLock()
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
