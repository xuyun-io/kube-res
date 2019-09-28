package main

import (
	"fmt"

	"github.com/xuyun-io/kube-res/main/resources/corev1"
	"github.com/xuyun-io/kube-res/main/resources/extensionsv1beta1"

	"github.com/mlycore/log"
)

const (
	testns = "observability"
)

func main() {
	c := KubernetesClientset()
	log.Infof("this is a log")

	if list, err := corev1.Pods(c.Clientset, testns); err == nil {
		for _, p := range list.Items {
			fmt.Printf("pod/%s/%s\n", p.Namespace, p.Name)
		}
	}

	if list, err := extensionsv1beta1.Deployments(c.Clientset, testns); err == nil {
		for _, dp := range list.Items {
			fmt.Printf("deployments/%s/%s\n", dp.Namespace, dp.Name)
		}
	}
}
