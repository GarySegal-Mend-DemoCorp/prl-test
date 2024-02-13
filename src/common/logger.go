package common

import (
	"context"
	"fmt"

	"github.com/Parallels/prl-devops-service/constants"

	log "github.com/cjlapao/common-go-logger"
)

var Logger = log.Get().WithTimestamp()

func LogInfo(ctx context.Context, format string, args ...string) {
	id := ctx.Value(constants.REQUEST_ID_KEY)
	if id != nil && id.(string) != "" {
		Logger.Info(fmt.Sprintf("[%s] %s", id.(string), fmt.Sprintf(format, args)))
	} else {
		Logger.Info(format, args)
	}
}
