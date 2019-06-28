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

func LsAdd(c *gin.Context) {
	log.Info("Logical Switch Add", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	if err := c.Bind(&Lsr); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ocmd, _ := ovndbapi.LSAdd(Lsr.Ls)
	var err error
	err = ovndbapi.Execute(ocmd)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
}

func LsGet(c *gin.Context) {
	log.Info("Logical Switch Get", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	if err := c.Bind(&Lsr); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ocmd, err := ovndbapi.LSGet(Lsr.Ls)
	if err != nil {
		handler.SendResponse(c,errno.ErrLsGet,nil)
		log.Fatal("err executing command:%v", err)
	}

	LogicalSwitchList, err := json.Marshal(ocmd)
	if err != nil {
		log.Fatal("err executing Json:%v", err)
	}
	handler.SendResponse(c, nil, LogicalSwitchList)
}

func LsDel(c *gin.Context){
	log.Info("Logical Switch Get", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var Lsr = LsRequest{}
	if err := c.Bind(&Lsr); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ocmd, err := ovndbapi.LSDel(Lsr.Ls)
	if err != nil {
		log.Fatal("err executing command:%v", err)
	}
}