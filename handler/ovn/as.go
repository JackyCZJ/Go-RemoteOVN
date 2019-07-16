package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

var err error

func ASAdd(c *gin.Context) {
	var as ASRequest
	err = c.BindJSON(&as)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if len(as.ExternalID) == 0 {
		as.ExternalID = nil
	}
	cmd, err := ovndbapi.ASAdd(as.Name, as.Addresses, as.ExternalID)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASAdd)
		return
	}
	log.Infof("Address set %s Add", as.Name)
	handler.SendResponse(c, nil, nil)
}

func ASDel(c *gin.Context) {
	as := c.Param("name")
	cmd, err := ovndbapi.ASDel(as)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASDel)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASDel)
		return
	}
	log.Infof("Address Set %s Delete", as)
	handler.SendResponse(c, nil, nil)
}

func ASUpdate(c *gin.Context) {
	var as ASRequest
	err = c.BindJSON(&as)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if len(as.ExternalID) == 0 {
		as.ExternalID = nil
	}
	cmd, err := ovndbapi.ASUpdate(as.Name, as.Addresses, as.ExternalID)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASUpdate)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASUpdate)
		return
	}
	log.Infof("Address Set %s Update", as.Name)
	handler.SendResponse(c, nil, nil)
}

func ASGet(c *gin.Context) {
	as := c.Param("name")
	cmd, err := ovndbapi.ASGet(as)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASGet)
		return
	}
	l := ASStruct(cmd)
	handler.SendResponse(c, nil, l)
}

func ASList(c *gin.Context) {
	cmd, err := ovndbapi.ASList()
	if err != nil {
		handleOvnErr(c, err, errno.ErrASList)
		return
	}
	var as []ASRequest
	for _, v := range cmd {
		l := ASStruct(v)
		as = append(as, l)
	}
	handler.SendResponse(c, nil, as)
}
