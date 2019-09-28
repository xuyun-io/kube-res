package main

import (
	"flag"
	"fmt"

	"github.com/mlycore/log"
	"github.com/xuyun-io/kube-res/main/resources/corev1"
	"github.com/xuyun-io/kube-res/main/resources/extensionsv1beta1"
)

const (
	defaultns = "default"
)

func main() {
	var namespace string
	flag.StringVar(&namespace, "n", "default", "")
	flag.StringVar(&namespace, "namespace", "default", "")

	c := KubernetesClientset()
	log.Infof("this is a log")

	if list, err := corev1.Pods(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			fmt.Printf("pod/%s/%s\n", p.Namespace, p.Name)
		}
	}

	if list, err := extensionsv1beta1.Deployments(c.Clientset, namespace); err == nil {
		for _, dp := range list.Items {
			fmt.Printf("deployments/%s/%s\n", dp.Namespace, dp.Name)
		}
	}
}
