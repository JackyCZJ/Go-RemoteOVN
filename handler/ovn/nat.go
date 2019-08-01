/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"fmt"

	"github.com/gin-gonic/gin"
	goovn "github.com/jackyczj/go-ovn"
	jsoniter "github.com/json-iterator/go"
)

type NAT struct {
	UUID        string            `json:"uuid"`
	Type        string            `json:"type"`
	ExternalIP  string            `json:"external_ip"`
	ExternalMAC string            `json:"external_mac"`
	LogicalIP   string            `json:"logical_ip"`
	LogicalPort string            `json:"logical_port"`
	ExternalID  map[string]string `json:"external_ids"`
}

var nat NAT
var cmd *goovn.OvnCommand

func LRNATAdd(c *gin.Context) {
	if err := c.BindJSON(&nat); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}
	lr := c.Param("name")
	cmd, err = ovndbapi.LRNATAdd(lr, nat.Type, nat.ExternalIP, nat.LogicalIP, nat.ExternalID, nat.LogicalPort, nat.ExternalMAC)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	handler.SendResponse(c, nil, nil)
}

func LRNATDel(c *gin.Context) {
	if err := c.BindJSON(&nat); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}
	lr := c.Param("name")
	var ip string
	switch nat.Type {
	case "snat":
		ip = nat.LogicalIP
	case "dnat":
		ip = nat.ExternalIP
	case "dnat_and_snat":
		ip = nat.ExternalIP
	default:
		handleOvnErr(c, fmt.Errorf("ERROR OPTION"), fmt.Errorf("ERROR OPTION"))
	}
	cmd, err = ovndbapi.LRNATDel(lr, nat.Type, ip)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	handler.SendResponse(c, nil, nil)
}

func LRNATList(c *gin.Context) {
	lr := c.Param("name")
	natList, err := ovndbapi.LRNATList(lr)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	NATlist := make([]NAT, len(natList))
	for i := 0; i < len(natList); i++ {
		str, _ := jsoniter.Marshal(natList[i])
		_ = jsoniter.Unmarshal(str, &nat)
		nat.ExternalID = MapInterfaceToMapString(natList[i].ExternalID)
		NATlist = append(NATlist, nat)
	}
	handler.SendResponse(c, nil, NATlist)
}
