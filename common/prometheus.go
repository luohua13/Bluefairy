package common

import (
	"github.com/prometheus/client_golang/prometheus"
)

var PodsExistOnePhysicalMachine = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "pods_exist_one_physical_machine",
		Help: "pods exist one physical machine",
	},
	[]string{"application"},
)

// Register metrics before server start
func init() {
	prometheus.MustRegister(PodsExistOnePhysicalMachine)
}
