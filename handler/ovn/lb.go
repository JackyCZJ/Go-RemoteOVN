/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

//Logical Bridge struct
type LBRequest struct {
	VipPort  string   `json:"vipPort"`
	Protocol string   `json:"protocol"`
	Addrs    []string `json:"addrs"`
}

func LBAdd(c *gin.Context) {
	var lb LBRequest
	err = c.BindJSON(&lb)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	l := c.Param("name")
	cmd, err := ovndbapi.LBAdd(l, lb.VipPort, lb.Protocol, lb.Addrs)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASAdd)
		return
	}
	log.Infof("LoadBlancer %s Add", l)
	handler.SendResponse(c, nil, nil)
}

func LBDel(c *gin.Context) {
	l := c.Param("name")
	ocmd, err := ovndbapi.LBDel(l)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
	}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
	}
	log.Infof("LoadBlancer Delete: %s", l)
	handler.SendResponse(c, nil, nil)
}

func LBGet(c *gin.Context) {
	l := c.Param("name")
	cmd, err := ovndbapi.LBGet(l)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLRDel)
		return
	}
	pack := make(map[string]interface{}, len(cmd))

	for _, v := range cmd {
		mapString := make(map[string]string)
		mapString = MapInterfaceToMapString(v.ExternalID)
		pack["UUID"] = v.UUID
		pack["Name"] = v.Name
		pack["ExternalID"] = mapString
	}

	handler.SendResponse(c, nil, pack)
}

func LBUpdate(c *gin.Context) {
	var lb LBRequest
	err = c.BindJSON(&lb)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	l := c.Param("name")
	cmd, err := ovndbapi.LBUpdate(l, lb.VipPort, lb.Protocol, lb.Addrs)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrASAdd)
		return
	}
	log.Infof("LoadBlancer %s Update", l)
	handler.SendResponse(c, nil, nil)
}
