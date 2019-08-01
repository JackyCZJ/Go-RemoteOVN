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

//ACL Request struct
type AclRequest struct {
	UUID        string            `json:"uuid"`
	Direct      string            `json:"direct"`
	Match       string            `json:"match"`
	Action      string            `json:"action"`
	Priority    int               `json:"priority"`
	ExternalIds map[string]string `json:"external_ids"`
	Logflag     bool              `json:"logflag"`
	Meter       string            `json:"meter"`
}

func ACLAdd(c *gin.Context) {
	var acl AclRequest
	var err error
	err = c.BindJSON(&acl)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		log.Fatal("JSON Error:", err)
		return
	}
	Ls := c.Param("name")
	cmd, err := ovndbapi.ACLAdd(Ls, acl.Direct, acl.Match, acl.Action, acl.Priority, acl.ExternalIds, acl.Logflag, acl.Meter)
	if err != nil {
		log.Fatal("ERROR:", err)
		handleOvnErr(c, err, errno.ErrACLAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		log.Fatal("ERROR:", err)
		handleOvnErr(c, err, errno.ErrACLAdd)
		return
	}
	handler.SendResponse(c, nil, nil)

}

func ACLDel(c *gin.Context) {
	var acl AclRequest
	var err error
	err = c.BindJSON(&acl)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		log.Fatal("JSON Error:", err)
		return
	}
	Ls := c.Param("name")
	cmd, err := ovndbapi.ACLDel(Ls, acl.Direct, acl.Match, acl.Priority, acl.ExternalIds)
	if err != nil {
		handleOvnErr(c, err, errno.ErrACLDel)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrACLDel)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func ACLList(c *gin.Context) {
	ls := c.Param("name")
	acls, err := ovndbapi.ACLList(ls)
	if err != nil {
		handleOvnErr(c, err, errno.ErrACLList)
		return
	}
	var acl []AclRequest
	for _, v := range acls {
		l := ACLStruct(v)
		acl = append(acl, l)
	}
	handler.SendResponse(c, nil, acl)
}
