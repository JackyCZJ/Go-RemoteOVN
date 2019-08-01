package ovn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeter(t *testing.T) {
	var jp = jsonPackage{
		arg: map[string]string{
			"name": "meter1",
		},
		data: map[string]interface{}{
			"action": "drop",
			"rate":   1,
			"unit":   "kbps",
		},
	}
	ginTestJsonTool(MeterAdd, jp, &req)

	assert.Equal(t, req.Message, "OK")

	ginTestPathTool(MeterList, args{}, &req)

	fmt.Println(req.Data)

	ginTestPathTool(MeterBandsList, args{}, &req)

	fmt.Println(req.Data)

	ginTestPathTool(MeterDel, args{arg: map[string]string{"name": "meter1"}}, &req)

}
