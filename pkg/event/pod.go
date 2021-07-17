package event

import (
	log "github.com/sirupsen/logrus"

	v1 "k8s.io/api/core/v1"
)

type PodPool struct {
	App      string
	HostIP   string
	PodName  string
	EsxiName string
}

var GlobalPodPool = make(map[string]*PodPool)

func AddPod(obj interface{}) {
	newPod, ok := obj.(*v1.Pod)
	if !ok {
		log.Println("Unexpected pod update Event")
		return
	}
	log.Println("pod delete:", newPod.Status.PodIP, newPod.Status.HostIP)
	// go func(name string, client kubernetes.Interface) {
	// 	time.Sleep(30)
	// 	truePod, err := client.CoreV1().Pods("").Get(name, metav1.GetOptions{})
	// 	if err != nil {
	// 		return
	// 	}
	// 	for _, owner := range truePod.GetOwnerReferences() {
	// 		if owner.Kind == "StatefulSet" || owner.Kind == "ReplicaSet" {
	// 			if truePod.Status.HostIP == "" {
	// 				log.Println("Failed to get new pod HostIP.")
	// 				return
	// 			}
	// 			podG := &PodPool{
	// 				podName: truePod.GetName(),
	// 				hostIP:  truePod.Status.HostIP,
	// 			}
	// 			podG.app = owner.Name
	// 			GlobalPodPool[podG.podName] = append(GlobalPodPool[podG.podName], podG)
	// 			log.Println("luohua:::add pod:", podG.podName)
	// 		}
	// 	}
	// }(name, client)
}

func UpdatePod(newObj interface{}) {
	newPod, ok := newObj.(*v1.Pod)
	if !ok {
		log.Println("Unexpected pod update Event")
		return
	}
	log.Println("pod update:", newPod.Status.PodIP, newPod.Status.HostIP)

	//
	//for _, owner := range newPod.GetOwnerReferences() {
	//	if owner.Kind == "StatefulSet" || owner.Kind == "ReplicaSet" {
	//		if newPod.Status.HostIP == "" {
	//			log.Println("Failed to get new pod HostIP.")
	//			return
	//		}
	//		podG := &PodPool{
	//			PodName: newPod.GetName(),
	//			HostIP:  newPod.Status.HostIP,
	//		}
	//		podG.App = owner.Name
	//		GlobalPodPool[podG.PodName] = podG
	//		log.Println("luohua:::update pod:", podG.PodName)
	//	}
	//}

}

func DeletePod(obj interface{}) {
	pod, ok := obj.(*v1.Pod)
	if !ok {
		return
	}
	log.Println("pod delete:", pod.Status.PodIP, pod.Status.HostIP)

	//for _, owner := range pod.GetOwnerReferences() {
	//	if owner.Kind == "StatefulSet" || owner.Kind == "ReplicaSet" {
	//		delete(GlobalPodPool, pod.GetName())
	//		log.Println("luohua:::delete pod :", pod.GetName())
	//	}
	//}
}

func ReplacePods(objs []interface{}) {
	//for k := range GlobalPodPool {
	//	delete(GlobalPodPool, k)
	//}
	for _, obj := range objs {
		pod, ok := obj.(*v1.Pod)
		if pod.Status.Phase != "Running" && !ok {
			continue
		}
		log.Println("pod replace:", pod.Status.PodIP, pod.Status.HostIP)
		//for _, owner := range pod.GetOwnerReferences() {
		//	if owner.Kind == "StatefulSet" || owner.Kind == "ReplicaSet" {
		//		podG := &PodPool{
		//			PodName: pod.GetName(),
		//			HostIP:  pod.Status.HostIP,
		//		}
		//		podG.App = owner.Name
		//		GlobalPodPool[podG.PodName] = podG
		//		log.Println("luohua:::replace:", podG)
		//	}
		//}
	}
}
