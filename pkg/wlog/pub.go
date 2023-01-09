package wlog

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func Error(ctx context.Context, err error) error {
	file, line := getFileLine()
	err = fmt.Errorf("file:%v line:%v %v",
		file, line, err)
	obj := map[string]interface{}{
		"file": file,
		"line": line,
		"info": err.Error(),
		//"reqId": wUtil.GetReqId(ctx),
		//"hardwareId": wUtil.GetHarwareId(ctx),
		//"ip":    wUtil.GetClientIp(ctx),
		"level": "error",
	}
	jBuff, _ := json.Marshal(obj)
	log.Println(string(jBuff))
	return err
}
