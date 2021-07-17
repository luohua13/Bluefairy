package informer

import (
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
)

type ResourceEventHandler interface {
	cache.ResourceEventHandler
	OnReplace(objs []interface{})
}

type ResourceEventHandlerFuncs struct {
	cache.ResourceEventHandlerFuncs
	ReplaceFunc func(objs []interface{})
}

func (r ResourceEventHandlerFuncs) OnReplace(objs []interface{}) {
	if r.ReplaceFunc != nil {
		r.ReplaceFunc(objs)
	}
}

type Config struct {
	cache.Queue

	cache.ListerWatcher

	name string

	ObjectType runtime.Object

	FullResyncPeriod time.Duration

	ShouldResync cache.ShouldResyncFunc
}

type controller struct {
	config         Config
	reflector      *cache.Reflector
	reflectorMutex sync.RWMutex
}

func NewController(c *Config) *controller {
	return &controller{
		config: *c,
	}
}

func (c *controller) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	r := cache.NewNamedReflector(
		c.config.name,
		c.config.ListerWatcher,
		c.config.ObjectType,
		c.config.Queue,
		c.config.FullResyncPeriod,
	)
	r.ShouldResync = c.config.ShouldResync

	c.reflectorMutex.Lock()
	c.reflector = r
	c.reflectorMutex.Unlock()

	var wg wait.Group
	defer wg.Wait()

	wg.StartWithChannel(stopCh, r.Run)

	<-stopCh
}

// Returns true once this controller has completed an initial resource listing
func (c *controller) HasSynced() bool {
	return c.config.Queue.HasSynced()
}

func (c *controller) LastSyncResourceVersion() string {
	if c.reflector == nil {
		return ""
	}
	return c.reflector.LastSyncResourceVersion()
}

func NewInformer(
	name string,
	lw cache.ListerWatcher,
	objType runtime.Object,
	resyncPeriod time.Duration,
	h ResourceEventHandler,
) cache.Controller {
	queue := NewProxyQueue(h)

	cfg := &Config{
		name:             name,
		Queue:            queue,
		ListerWatcher:    lw,
		ObjectType:       objType,
		FullResyncPeriod: resyncPeriod,
	}
	return NewController(cfg)
}
