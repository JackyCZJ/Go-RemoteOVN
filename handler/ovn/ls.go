package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

//	@Summary Add new Logical switch
//	@Description Add new Logical switch
//	@Tags	Logical switch
//	@Accept	string
//	@Produce err or nil
//	@Param	id body ovn.LsAdd true
//	@Router /:id PUT
func LSAdd(c *gin.Context) {
	log.Info("Logical Switch Add", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	Lsr.Ls = c.Param("id")
	ocmd, _ := ovndbapi.LSAdd(Lsr.Ls)
	var err error
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
	req:=CreateResponse{
		Name:Lsr.Ls,
		Type:"Switch",
		Action:"Create",
	}

	handler.SendResponse(c,nil,req)

}

//	@Summary Get Logical switch
//	@Description Add new Logical switch by Name
//	@Tags	Logical switch
//	@Accept	string
//	@Produce json
//	@Param	id body ovn.LsAdd true
//	@Router /:id PUT
func LSGet(c *gin.Context) {
	log.Info("Logical Switch Get", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	Lsr.Ls = c.Param("id")
	ocmd, err := ovndbapi.LSGet(Lsr.Ls)
	if err != nil {
		handler.SendResponse(c, errno.ErrLsGet, nil)
		log.Fatal("err executing command:%v", err)
	}

	LogicalSwitchList, err := json.Marshal(ocmd)
	if err != nil {
		log.Fatal("err executing Json:%v", err)
	}
	log.Info("Action Success !", lager.Data{"X-Request-Id": util.GetReqID(c)})
	handler.SendResponse(c, nil, LogicalSwitchList)
}

func LSDel(c *gin.Context) {
	log.Info("Logical Switch Delete", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	Lsr.Ls = c.Param("id")
	ocmd, err := ovndbapi.LSDel(Lsr.Ls)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
	req:=CreateResponse{
		Name:Lsr.Ls,
		Type:"Switch",
		Action:"Delete",
	}

	handler.SendResponse(c,nil,req)

}

func LSList(c *gin.Context) {
	log.Info("Logical Switch Get List", lager.Data{"X-Request-Id": util.GetReqID(c)})
	ocmd, err := ovndbapi.LSList()
	if err != nil {
		handler.SendResponse(c, errno.ErrLsGet, nil)
		log.Fatal("err executing command:%v", err)
	}

	LogicalSwitchList, err := json.Marshal(ocmd)
	if err != nil {
		log.Fatal("err executing Json:%v", err)
	}
	handler.SendResponse(c, nil, LogicalSwitchList)
	log.Info("Action Success !", lager.Data{"X-Request-Id": util.GetReqID(c)})

}

func LsExtIdsAdd(c *gin.Context) {

}
