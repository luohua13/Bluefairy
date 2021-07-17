package app

import (
	"Bluefairy/infra"
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

type ControllerContext struct {
	// InformerFactory gives access to informers for the controller.
	InformerFactory informers.SharedInformerFactory
	// Stop is the stop channel
	Stop   <-chan struct{}
	Client kubernetes.Interface
}

type InitFunc func(ctx ControllerContext) (debuggingHandler http.Handler, enabled bool, err error)

// GetTimeDuration generate a time duration of seconds
func getTimeDuration(t int) time.Duration {
	return time.Duration(t) * time.Second
}

func Run(ctx context.Context) {
	clustersCfg, err := infra.GetKubernetesConfig()
	if err != nil {
		log.Fatalf("Get cluster config error:%s", err)
	}

	for _, cfg := range clustersCfg {
		//cfg.Insecure = true
		controllerContext, err := CreateControllerContext(cfg, ctx.Done())
		if err != nil {
			log.Fatalf("error building controller context: %v", err)
		}

		if err := StartControllers(controllerContext, NewControllerInitializers()); err != nil {
			log.Fatalf("error starting controllers: %v", err)
		}
	}
	select {}
}

func CreateControllerContext(c *restclient.Config, stop <-chan struct{}) (ControllerContext, error) {
	versionedClient, err := kubernetes.NewForConfig(c)
	if err != nil {
		return ControllerContext{}, err
	}

	ctx := ControllerContext{
		Stop:   stop,
		Client: versionedClient,
	}
	return ctx, nil
}

func NewControllerInitializers() map[string]InitFunc {
	controllers := map[string]InitFunc{}
	controllers["microservicePolicy"] = startMicroServicePolicyController

	return controllers
}

// StartControllers starts a set of controllers with a specified ControllerContext
func StartControllers(ctx ControllerContext, controllers map[string]InitFunc) error {
	for _, initFn := range controllers {
		_, started, err := initFn(ctx)
		if err != nil {
			return err
		}
		if !started {
			continue
		}
	}
	return nil
}
