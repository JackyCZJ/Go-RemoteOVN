/*
 * Copyright (c) 2019. eSix Inc.
 */
package ovn

import (
	"apiserver/handler"
	"apiserver/pkg/errno"

	goovn "github.com/jackyczj/go-ovn"

	jsoniter "github.com/json-iterator/go"

	"github.com/gin-gonic/gin"
)

type Meter struct {
	UUID        string
	Name        string            `json:"name"`
	Unit        string            `json:"unit"`
	Bands       []string          `json:"bands"`
	ExternalIds map[string]string `json:"external_ids"`
}

type MeterBand struct {
	UUID      string
	Action    string `json:"action"`
	Rate      int    `json:"rate"`
	BurstSize int    `json:"burst_size"`
}

type MeterRequest struct {
	Action      string            `json:"action"`
	Rate        int               `json:"rate"`
	Unit        string            `json:"unit"`
	BurstSize   int               `json:"burst_size"`
	ExternalIds map[string]string `json:"external_ids"`
}

func MeterAdd(c *gin.Context) {
	var meter *MeterRequest
	name := c.Param("name")
	if err := c.BindJSON(&meter); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}
	cmd, err := ovndbapi.MeterAdd(name, meter.Action, meter.Rate, meter.Unit, meter.ExternalIds, meter.BurstSize)
	if err != nil {
		handleOvnErr(c, err, err)
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handler.SendResponse(c, err, err)
	}
	handler.SendResponse(c, nil, nil)

}

func MeterDel(c *gin.Context) {
	name := c.Param("name")
	var cmd *goovn.OvnCommand
	var err error

	cmd, err = ovndbapi.MeterDel()
	if name != "" {
		cmd, err = ovndbapi.MeterDel(name)
	}

	if err != nil {
		handleOvnErr(c, err, err)
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		handler.SendResponse(c, err, err)
	}
	handler.SendResponse(c, nil, nil)
}

func MeterList(c *gin.Context) {
	var meter *Meter
	var meterList []*Meter
	meters, err := ovndbapi.MeterList()
	if err != nil {
		handleOvnErr(c, err, err)
	}
	for _, v := range meters {
		str, _ := jsoniter.Marshal(v)
		err = jsoniter.Unmarshal(str, &meter)
		meter.ExternalIds = MapInterfaceToMapString(v.ExternalIds)
		meterList = append(meterList, meter)
	}
	handler.SendResponse(c, nil, meterList)
}

func MeterBandsList(c *gin.Context) {
	var meterBand *MeterBand
	var meterBandList []*MeterBand

	meterBands, err := ovndbapi.MeterBandsList()
	if err != nil {
		handleOvnErr(c, err, err)
	}
	for _, v := range meterBands {
		str, _ := jsoniter.Marshal(v)
		err = jsoniter.Unmarshal(str, &meterBand)
		meterBandList = append(meterBandList, meterBand)
	}
	handler.SendResponse(c, nil, meterBandList)
}
