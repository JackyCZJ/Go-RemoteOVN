/*
 * Copyright (c) 2019. eSix Inc.
 */

package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	jsoniter "github.com/json-iterator/go"

	"github.com/gin-gonic/gin"
)

type DHCPOptions struct {
	UUID       string				`json:"uuid"`
	CIDR       string				`json:"cidr"`
	Options    map[string]string	`json:"options"`
	ExternalID map[string]string	`json:"external_id"`
}

func DHCPOptionsAdd(c *gin.Context) {
	dhcp := DHCPOptions{}
	if err := c.BindJSON(&dhcp); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}
	cmd, err := ovndbapi.DHCPOptionsAdd(dhcp.CIDR, dhcp.Options, dhcp.ExternalID)
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

func DHCPOptionSet(c *gin.Context) {
	dhcp := DHCPOptions{}
	if err := c.BindJSON(&dhcp); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}
	cmd, err := ovndbapi.DHCPOptionsSet(dhcp.CIDR, dhcp.Options, dhcp.ExternalID)
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

func DHCPOptionsDel(c *gin.Context) {
	uuid := c.Param("uuid")
	cmd, err := ovndbapi.DHCPOptionsDel(uuid)
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

func DHCPOptionsList(c *gin.Context) {
	cmd, err := ovndbapi.DHCPOptionsList()
	if err != nil {
		handleOvnErr(c, err, err)
	}
	var dhcpLs []DHCPOptions
	var dhcp DHCPOptions
	for _, v := range cmd {
		str, _ := jsoniter.Marshal(v)
		_ = jsoniter.Unmarshal(str, &dhcp)
		dhcp.Options = MapInterfaceToMapString(v.Options)
		dhcp.ExternalID = MapInterfaceToMapString(v.ExternalID)
		dhcpLs = append(dhcpLs, dhcp)
	}
	handler.SendResponse(c, nil, dhcpLs)
}

func LSPSetDHCPv4Options(c *gin.Context) {
	lsp := c.Param("name")
	uuid := c.Param("uuid")
	cmd, err := ovndbapi.LSPSetDHCPv4Options(lsp, uuid)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func LSPGetDHCPv4Options(c *gin.Context) {
	lsp := c.Param("name")
	cmd, err := ovndbapi.LSPGetDHCPv4Options(lsp)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	dhcp := DHCPOptions{}
	str, _ := jsoniter.Marshal(cmd)
	_ = jsoniter.Unmarshal(str, &dhcp)
	dhcp.Options = MapInterfaceToMapString(cmd.Options)
	dhcp.ExternalID = MapInterfaceToMapString(cmd.ExternalID)

	handler.SendResponse(c, nil, dhcp)
}

// TODO : May need to use an DHCP V6 URL
func LSPSetDHCPv6Options(c *gin.Context) {
	lsp := c.Param("name")
	uuid := c.Param("uuid")
	cmd, err := ovndbapi.LSPSetDHCPv6Options(lsp, uuid)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	handler.SendResponse(c, nil, nil)
}

// TODO : May need to use an DHCP V6 URL
func LSPGetDHCPv6Options(c *gin.Context) {
	lsp := c.Param("name")
	cmd, err := ovndbapi.LSPGetDHCPv6Options(lsp)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	dhcp := DHCPOptions{}
	str, _ := jsoniter.Marshal(cmd)
	_ = jsoniter.Unmarshal(str, &dhcp)
	dhcp.Options = MapInterfaceToMapString(cmd.Options)
	dhcp.ExternalID = MapInterfaceToMapString(cmd.ExternalID)

	handler.SendResponse(c, nil, dhcp)
}
