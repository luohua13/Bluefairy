package event

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

func AddEndpoints(obj interface{}) {
	endpoints, ok := obj.(*v1.Endpoints)
	if !ok {
		return
	}

	var ipList []string
	for _, epSubset := range endpoints.Subsets {
		for _, addr := range epSubset.Addresses {
			ipList = append(ipList, addr.IP)
		}
	}
	fmt.Println("AddEndpoints:", endpoints.Name, endpoints.Namespace, ipList)
}

func UpdateEndpoints(obj interface{}) {
	endpoints, ok := obj.(*v1.Endpoints)
	if !ok {
		return
	}

	var ipList []string
	if len(endpoints.Subsets) != 0 {
		for _, epSubset := range endpoints.Subsets {
			for _, addr := range epSubset.Addresses {
				ipList = append(ipList, addr.IP)
			}
		}
		fmt.Println("UpdateEndpoints:", endpoints.Name, endpoints.Namespace, ipList)
	}
}

func DeleteEndpoints(obj interface{}) {
	endpoints, ok := obj.(*v1.Endpoints)
	if !ok {
		return
	}
	fmt.Println("DeleteEndpoints:", endpoints.Name, endpoints.Namespace)

}

func ReplaceEndpoints(objs []interface{}) {
	for _, obj := range objs {
		endpoints, ok := obj.(*v1.Endpoints)
		if !ok {
			continue
		}

		var ipList []string
		for _, epSubset := range endpoints.Subsets {
			for _, addr := range epSubset.Addresses {
				ipList = append(ipList, addr.IP)
			}
		}
		fmt.Println("ReplaceEndpoints:", endpoints.Name, endpoints.Namespace, ipList)
	}
}
