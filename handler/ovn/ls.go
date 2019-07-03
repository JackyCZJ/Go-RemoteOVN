package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

//	@Summary Add new Logical switch
//	@Description Add new Logical switch
//	@Tags	Logical switch
//	@Produce json
//	@Param	name path string true "Logical Switch Name"
//	@Router /api/v1/esix/ovn/LS/{name} [PUT]
func LSAdd(c *gin.Context) {
//	log.Info("Logical Switch Add")
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
	handler.SendResponse(c, nil, nil)
}

//	@Summary GET Logical switch By name
//	@Description Get a Logical switch
//	@Tags	Logical switch
//	@Accept json
//	@Produce json
//	@Param	name path string true "Logical Switch Name"
//	@Router /api/v1/esix/ovn/LS/{name} [GET]
func LSGet(c *gin.Context) {
	log.Info("Logical Switch Get")
	var Lsr = LsRequest{}
	Lsr.Ls = c.Param("name")
	ocmd, err := ovndbapi.LSGet(Lsr.Ls)
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
//	@Param	name path string true "Logical Switch Name"
//	@Router /api/v1/esix/ovn/LS/{name} [DELETE]
func LSDel(c *gin.Context) {
	log.Info("Logical Switch Delete", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	Lsr.Ls = c.Param("name")
	ocmd, err := ovndbapi.LSDel(Lsr.Ls)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsDel)
		return

	}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsDel)
		return
	}

	handler.SendResponse(c, nil, nil)

}

//	@Summary Get List Of Logical switch
//	@Description  get Logical switch list
//	@Tags	Logical switch
//	@Produce json
//	@Router /api/esix/ovn/LS [GET]
func LSList(c *gin.Context) {
	log.Info("Logical Switch Get List", lager.Data{"X-Request-Id": util.GetReqID(c)})
	ocmd, err := ovndbapi.LSList()
	if err != nil {
		handleOvnErr(c, err, errno.ErrLsListGet)
		return
	}
	var lslist []LogicalSwitch
	for _, v := range ocmd {
		var l LogicalSwitch
		l = logicalSwitchStruct(v)
		lslist = append(lslist, l)
	}
	handler.SendResponse(c, nil, lslist)
}

//json
//	ls logical Switch
//	"Extid":{
//			"key":"value"
//			}
//
//
func LsExtIdsAdd(c *gin.Context) {
	//Ext id map[string][string]
	var r LogicalSwitch
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
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
	req := CreateResponse{
		Name:   r.Name,
		Type:   "External ID",
		Action: "Add",
	}

	handler.SendResponse(c, nil, req)
}

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

	handler.SendResponse(c, nil, nil)
}

//Delete Port from its attached switch 把网口从绑定的逻辑交换机上删除
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

	handler.SendResponse(c, nil, nil)
}
