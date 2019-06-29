package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
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
	log.Info("Logical Switch Add", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	defer Lsr.Unlock()
	Lsr.Lock()
	Lsr.Ls = c.Param("name")
	ocmd, _ := ovndbapi.LSAdd(Lsr.Ls)
	var err error
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
	req := CreateResponse{
		Name:   Lsr.Ls,
		Type:   "Switch",
		Action: "Create",
	}
	handler.SendResponse(c, nil, req)
}

//	@Summary GET Logical switch By name
//	@Description Get a Logical switch
//	@Tags	Logical switch
//	@Accept json
//	@Produce json
//	@Param	name path string true "Logical Switch Name"
//	@Router /api/v1/esix/ovn/LS/{name} [GET]
func LSGet(c *gin.Context) {
	log.Info("Logical Switch Get", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	Lsr.Ls = c.Param("name")
	ocmd, err := ovndbapi.LSGet(Lsr.Ls)
	if err != nil {
		handler.SendResponse(c, errno.ErrLsGet, nil)
		log.Fatal("err executing command:%v", err)
	}
	var l LogicalSwitch
	for _, v := range ocmd {
		str, _ := jsoniter.Marshal(v)
		err := jsoniter.Unmarshal(str, &l)
		if err != nil {
			log.Fatal("err executing command:%v", err)
		}
	}
	handler.SendResponse(c, nil, l)
}

//	@Summary Delete Logical switch
//	@Description Delete a Logical switch
//	@Tags	Logical switch
//	@Produce json
//	@Param	name path string true "Logical Switch Name"
//	@Router /api/esix/ovn/LS/{name} [DELETE]
func LSDel(c *gin.Context) {
	log.Info("Logical Switch Delete", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	defer Lsr.Unlock()
	Lsr.Lock()
	Lsr.Ls = c.Param("name")
	ocmd, err := ovndbapi.LSDel(Lsr.Ls)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
	req := CreateResponse{
		Name:   Lsr.Ls,
		Type:   "Switch",
		Action: "Delete",
	}

	handler.SendResponse(c, nil, req)

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
		handler.SendResponse(c, errno.ErrLsGet, nil)
		log.Fatal("err executing command:%v", err)
	}
	var lslist []LogicalSwitch
	for _, v := range ocmd {
		var l LogicalSwitch
		str, _ := jsoniter.Marshal(v)
		err := jsoniter.Unmarshal(str, &l)
		if err != nil {
			log.Fatal("err executing command:%v", err)
		}
		lslist = append(lslist, l)
	}
	handler.SendResponse(c, nil, lslist)
}
