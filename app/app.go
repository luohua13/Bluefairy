package app

import (
	"Bluefairy/pkg/controller/application"
	"net/http"
)

func startMicroServicePolicyController(ctx ControllerContext) (http.Handler, bool, error) {

	go application.NewMicroServicePolicyController(
		ctx.Client.CoreV1().Pods("testhl"),
		ctx.Client.CoreV1().Endpoints("testhl"),
	).Start(ctx.Stop)

	return nil, true, nil
}
