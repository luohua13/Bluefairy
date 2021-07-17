package informer

import (
	"k8s.io/client-go/tools/cache"
)

type ProxyQueue struct {
	resynced bool
	h        ResourceEventHandler
}

func NewProxyQueue(h ResourceEventHandler) *ProxyQueue {
	return &ProxyQueue{
		resynced: false,
		h:        h,
	}
}

func (p *ProxyQueue) Add(obj interface{}) error {
	if p.h != nil {
		p.h.OnAdd(obj)
	}
	return nil
}

func (p *ProxyQueue) Update(obj interface{}) error {
	if p.h != nil {
		p.h.OnUpdate(nil, obj)
	}
	return nil
}

func (p *ProxyQueue) Delete(obj interface{}) error {
	if p.h != nil {
		p.h.OnDelete(obj)
	}
	return nil
}

func (p *ProxyQueue) List() []interface{} {
	return nil
}

func (p *ProxyQueue) ListKeys() []string {
	return nil
}

func (p *ProxyQueue) Get(obj interface{}) (item interface{}, exists bool, err error) {
	return nil, false, nil
}

func (p *ProxyQueue) GetByKey(key string) (item interface{}, exists bool, err error) {
	return nil, false, nil
}

func (p *ProxyQueue) Replace(objs []interface{}, resourceversion string) error {
	if p.h != nil {
		p.h.OnReplace(objs)
	}
	p.resynced = true
	return nil
}

// Resync is not necessary here, cause ProxyQueue store nothing but proxy the events.
func (p *ProxyQueue) Resync() error {
	return nil
}

func (p *ProxyQueue) Pop(pop cache.PopProcessFunc) (interface{}, error) {
	return nil, nil
}

func (p *ProxyQueue) AddIfNotPresent(interface{}) error {
	return nil
}

// Return true if the first batch of itemp has been popped
func (p *ProxyQueue) HasSynced() bool {
	return p.resynced
}

// Close queue
func (p *ProxyQueue) Close() { /* Nothing to close */ }
