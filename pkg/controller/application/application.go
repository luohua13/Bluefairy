package application

import (
	"Bluefairy/common"
	"Bluefairy/pkg/event"
	"Bluefairy/pkg/kubernetes/informer"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"log"
)

type AppController struct {
	podInformer       cache.Controller
	endpointsInformer cache.Controller
	syncHandler       func()
}

func NewMicroServicePolicyController(podClient v1.PodInterface, endpointsClient v1.EndpointsInterface) *AppController {

	podInformer := informer.NewPodInformer(podClient)
	endpointsInformer := informer.NewEndpointsInformer(endpointsClient)

	app := &AppController{
		podInformer:       podInformer,
		endpointsInformer: endpointsInformer,
	}
	app.syncHandler = app.Reconcile
	return app
}

func (app *AppController) Start(stopCh <-chan struct{}) {
	go app.endpointsInformer.Run(stopCh)
	go app.podInformer.Run(stopCh)
	//go wait.Until(app.syncHandler, time.Minute, stopCh)
}

func (app *AppController) Reconcile() {
	var appList []string
	var danger float64
	for _, v := range event.GlobalPodPool {
		appList = append(appList, v.App)
	}

	dupApp := common.DupCount(appList)

	for app, count := range dupApp {
		log.Printf("Item : %s , Count : %d\n", app, count)
		var hostList []string
		for _, v := range event.GlobalPodPool {
			if v.App == app {
				hostList = append(hostList, v.EsxiName)
			}
		}
		dupHost := common.DupCount(hostList)
		for EsxiName, c := range dupHost {
			log.Printf("hostName: %s Count: %d\n", EsxiName, c)
			if c == 1 {
				danger = 0
				log.Println("Only one pod for the app:", app)
			} else {
				danger = (float64)(c) / (float64)(count)
				log.Printf("bluefairy watch: %f\n", danger)
			}
			common.PodsExistOnePhysicalMachine.WithLabelValues(app).Add(danger)
		}
	}
}
