package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
)


var err error

func ASAdd(c *gin.Context){
	var as ASRequest
	err = c.BindJSON(&as);if err!=nil{
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if len(as.ExternalID) == 0{
		as.ExternalID = nil
	}
	defer  ovndbapi.Unlock()
	ovndbapi.Lock()
	cmd, err := ovndbapi.ASAdd(as.Name,as.Addresses,as.ExternalID)
	if err != nil{
		handleOvnErr(c,err,errno.ErrASAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil{
		handleOvnErr(c,err,errno.ErrASAdd)
		return
	}
	handler.SendResponse(c,nil,nil)
}

func ASDel(c *gin.Context){
	as := c.Param("name")
	defer  ovndbapi.Unlock()
	ovndbapi.Lock()
	cmd, err := ovndbapi.ASDel(as)
	if err != nil{
		handleOvnErr(c,err,errno.ErrASDel)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil{
		handleOvnErr(c,err,errno.ErrASDel)
		return
	}
	handler.SendResponse(c,nil,nil)
}

func ASUpdate(c *gin.Context){
	var as ASRequest
	err = c.BindJSON(&as);if err!=nil{
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if len(as.ExternalID) == 0{
		as.ExternalID = nil
	}
	defer  ovndbapi.Unlock()
	ovndbapi.Lock()
	cmd, err := ovndbapi.ASUpdate(as.Name,as.Addresses,as.ExternalID)
	if err != nil{
		handleOvnErr(c,err,errno.ErrASUpdate)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil{
		handleOvnErr(c,err,errno.ErrASUpdate)
		return
	}
	handler.SendResponse(c,nil,nil)
}

func ASGet(c *gin.Context){
	as := c.Param("name")
	defer ovndbapi.RUnlock()
	ovndbapi.RLock()
	cmd, err := ovndbapi.ASGet(as)
	if err != nil {
		handleOvnErr(c,err,errno.ErrASGet)
		return
	}
	l := ASStruct(cmd)
	handler.SendResponse(c,nil,l)
}

func ASList(c *gin.Context){
	defer ovndbapi.RUnlock()
	ovndbapi.RLock()
	cmd, err := ovndbapi.ASList()
	if err != nil {
		handleOvnErr(c,err,errno.ErrASList)
		return
	}
	var as []ASRequest
	for _ , v := range cmd{
		l := ASStruct(v)
		as = append(as,l)
	}
	handler.SendResponse(c,nil,as)
}