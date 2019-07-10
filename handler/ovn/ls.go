package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)


//	@Summary Add new Logical switch
//	@Description Add new Logical switch
//	@Tags	Logical switch
//	@Produce json
//  @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":nil"}
//	@Param	name path string true "Logical Switch Name"
//	@Router /api/v1/esix/ovn/LS/{name} [PUT]
func LSAdd(c *gin.Context) {
	ls := c.Param("name")
	var err error
	cmd, err := ovndbapi.LSAdd(ls)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsAdd)
		return
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsAdd)
		return
	}
	log.Infof("Logical Switch Add: %s",ls)
	handler.SendResponse(c, nil, nil)
}

//	@Summary GET Logical switch By name
//	@Description Get a Logical switch
//	@Tags	Logical switch
//	@Accept json
//	@Produce json
//  @Success 200 {object} handler.Response "{"code":0,"message":"OK","data": { "uuid": "a6b50553-9366-45d6-9e62-37335144b6c3", "name": "test2", "ports": [], "load_balancer": null, "acls": [], "qos_rules": null, "dns_records": null, "other_config": null, "external_id": {}}"}
//	@Param	name path string true "Logical Switch Name"
//	@Router /api/v1/esix/ovn/LS/{name} [GET]
func LSGet(c *gin.Context) {
	log.Info("Logical Switch Get")
	Ls := c.Param("name")
	ocmd, err := ovndbapi.LSGet(Ls)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsGet)
		return
	}
	var l LogicalSwitch
	for _, v := range ocmd {
		l = logicalSwitchStruct(v)
	}
	handler.SendResponse(c, nil, l)
}

//	@Summary Delete Logical switch
//	@Description Delete a Logical switch
//	@Tags	Logical switch
//	@Produce json
//  @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":nil"}
//	@Param	name path string true "Logical Switch Name"
//	@Router /api/v1/esix/ovn/LS/{name} [DELETE]
func LSDel(c *gin.Context) {
	Ls := c.Param("name")
	ocmd, err := ovndbapi.LSDel(Ls)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsDel)
		return

	}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsDel)
		return
	}
	log.Infof("Logical Switch Delete : %s", Ls)
	handler.SendResponse(c, nil, nil)

}

//	@Summary Get List Of Logical switch
//	@Description  get Logical switch list
//	@Tags	Logical switch
//	@Produce json
//  @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":{[ "uuid": "a6b50553-9366-45d6-9e62-37335144b6c3", "name": "test2", "ports": [], "load_balancer": null, "acls": [], "qos_rules": null, "dns_records": null, "other_config": null, "external_id": {} }]"}
//	@Router /api/esix/ovn/LS [GET]
func LSList(c *gin.Context) {
	lss, err := ovndbapi.LSList()
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsListGet)
		return
	}
	var lslist []LogicalSwitch
	for _, v := range lss {
		l := logicalSwitchStruct(v)
		lslist = append(lslist, l)
	}
	handler.SendResponse(c, nil, lslist)
}

//	@Summary Ls Ext IDs add
//  @Description add extends ids to ls
//	@Tags Logical switch
//  @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":nil"}
//	@Router /api/esix/ovn/LsExt/{name} [PUT]
func LsExtIdsAdd(c *gin.Context) {
	//Ext id map[string][string]
	var r LogicalSwitch
	if err := c.BindJSON(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	r.Name = c.Param("name")
	command, err := ovndbapi.LSExtIdsAdd(r.Name, r.ExternalID)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsExidOprate)
		return
	}
	err = ovndbapi.Execute(command)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsExidOprate)
		return
	}
	log.Infof("Logical Switch %s add External Id: %v",r.Name,r.ExternalID)
	handler.SendResponse(c, nil, nil)
}

//	@Summary Ls Ext IDs Delete
//  @Description Delete extends ids form ls
//	@Tags Logical switch
//  @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":nil"}
//	@Router /api/esix/ovn/LsExt/{name} [Delete]
func LsExtIdsDel(c *gin.Context) {
	var r LogicalSwitch
	if err := c.BindJSON(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	r.Name = c.Param("name")
	command, err := ovndbapi.LSExtIdsDel(r.Name, nil)
	if len(r.ExternalID) != 0{
		command, err = ovndbapi.LSExtIdsDel(r.Name, r.ExternalID)
	}
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsExidOprate)
		return
	}
	err = ovndbapi.Execute(command)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsExidOprate)
		return
	}
	log.Infof("Logical Switch %s Delete External Id: %v",r.Name,r.ExternalID)
	handler.SendResponse(c, nil, nil)
}

//	@Summary Add Port to a logical switch
//	@Description  add Port to switch
//	@Tags	Logical switch Port
//	@Produce json
//	@Router /api/esix/ovn/Lsp/{name}/{port} [DELETE]
func LSPAdd(c *gin.Context) {
	ls := c.Param("name")
	lp := c.Param("port")
	ocmd, err := ovndbapi.LSPAdd(ls, lp)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	log.Infof("Logical Switch %s Add Port : %s ",ls,lp)
	handler.SendResponse(c, nil, nil)
}

//	把网口从绑定的逻辑交换机上删除
//	@Summary Delete Port from its attached switch
//	@Description  Delete Port from its attached switch
//	@Tags	Logical switch Port
//	@Produce json
//	@Router /api/esix/ovn/Lsp/{port} [DELETE]
func LSPDel(c *gin.Context) {
	lp := c.Param("port")
	ocmd, err := ovndbapi.LSPDel(lp)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		handleOvnErr(c, err, err)
		return
	}
	log.Infof("Logical Switch Port unattached: %s ",lp)
	handler.SendResponse(c, nil, nil)
}
