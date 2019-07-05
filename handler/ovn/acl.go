package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)


func ACLAdd(c *gin.Context) {
	var acl AclRequest
	var err error
	err = c.BindJSON(&acl)
	if err != nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		log.Fatal("JSON Error:",err)
		return
	}
	acl.Ls = c.Param("name")
	cmd,err := ovndbapi.ACLAdd(acl.Ls,acl.Direct,acl.Match,acl.Action,acl.Priority,acl.ExternalIds,acl.Logflag,acl.Meter)
	if err != nil{
		handleOvnErr(c,err,errno.ErrACLAdd)
		return
	}
	defer ovndbapi.Unlock()
	ovndbapi.Lock()
	err = ovndbapi.Execute(cmd)
	if err != nil{
		handleOvnErr(c,err,errno.ErrACLAdd)
		return
	}
	handler.SendResponse(c,nil,nil)

}

func ACLDel(c *gin.Context) {
	var acl AclRequest
	var err error
	err = c.BindJSON(&acl)
	if err != nil{
		handler.SendResponse(c,errno.ErrBind,nil)
		log.Fatal("JSON Error:",err)
		return
	}
	acl.Ls = c.Param("name")
	cmd,err := ovndbapi.ACLDel(acl.Ls,acl.Direct,acl.Match,acl.Priority,acl.ExternalIds)
	if err != nil{
		handleOvnErr(c,err,errno.ErrACLDel)
		return
	}
	defer ovndbapi.Unlock()
	ovndbapi.Lock()
	err = ovndbapi.Execute(cmd)
	if err != nil{
		handleOvnErr(c,err,errno.ErrACLDel)
		return
	}
	handler.SendResponse(c,nil,nil)
}

func ACLList(c *gin.Context) {
	ls := c.Param("name")
	defer ovndbapi.RUnlock()
	ovndbapi.RLock()
	acls, err := ovndbapi.ACLList(ls)
	if err != nil{
		handleOvnErr(c,err,errno.ErrACLList)
		return
	}
	var acl []ACL
	for _, v := range acls {
		l := ACLStruct(v)
		acl = append(acl, l)
	}
	handler.SendResponse(c,nil,acl)
}
