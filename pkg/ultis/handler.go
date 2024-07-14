package ultis

import (
	"fmt"

	"go.uber.org/zap"
)

func HandleErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, zap.Error(err))
		return
	}
}

func HandleShouldBindJSONErr(err error, msg string) {
	if err != nil {
		HandleErr(err, msg)
		return
	}
}
