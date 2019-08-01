/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

//QOS struct
type QoS struct {
	UUID       string            `json:"uuid"`
	Priority   int               `json:"priority"`
	Direction  string            `json:"direction"`
	Match      string            `json:"match"`
	Action     map[string]int    `json:"action"`
	Bandwidth  map[string]int    `json:"bandwidth"`
	ExternalID map[string]string `json:"external_id"`
}

func QoSAdd(c *gin.Context) {
	ls := c.Param("name")
	var qos QoS
	if err := c.BindJSON(&qos); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}
	cmd, err := ovndbapi.QoSAdd(ls, qos.Direction, qos.Priority, qos.Match, qos.Action, qos.Bandwidth, qos.ExternalID)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func QoSDel(c *gin.Context) {
	ls := c.Param("name")
	var qos QoS
	if err := c.BindJSON(&qos); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	cmd, err := ovndbapi.QoSDel(ls, qos.Direction, qos.Priority, qos.Match)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	handler.SendResponse(c, nil, nil)
}

func QoSList(c *gin.Context) {
	ls := c.Param("name")
	qoslist, err := ovndbapi.QoSList(ls)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	var qosList []QoS
	var qos QoS
	for _, v := range qoslist {
		str, _ := jsoniter.Marshal(v)
		_ = jsoniter.Unmarshal(str, &qos)
		qos.ExternalID = MapInterfaceToMapString(v.ExternalID)
		qos.Bandwidth = MapInterfaceToMapint(v.Bandwidth)
		qos.Action = MapInterfaceToMapint(v.Action)
		qosList = append(qosList, qos)
	}
	handler.SendResponse(c, nil, qosList)

}
