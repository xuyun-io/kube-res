package main

import (
	"fmt"

	"github.com/xuyun-io/kube-res/main/resources"

	"github.com/mlycore/log"
)

const (
	testns = "observability"
)

func main() {
	c := KubernetesClientset()
	log.Infof("this is a log")

	if list, err := resources.Pod(c.Clientset, testns); err == nil {
		for _, p := range list.Items {
			fmt.Printf("pod/%s/%s\n", p.Namespace, p.Name)
		}
	}

	if list, err := resources.Deployments(c.Clientset, testns); err == nil {
		for _, dp := range list.Items {
			fmt.Printf("deployments/%s/%s\n", dp.Namespace, dp.Name)
		}
	}
}
