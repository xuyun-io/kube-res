package main

import (
	"flag"

	"github.com/xuyun-io/kube-res/main/resources/corev1"
	"github.com/xuyun-io/kube-res/main/resources/extensionsv1beta1"
	"github.com/xuyun-io/kube-res/main/utils"
)

const (
	defaultns = "default"
)

func main() {
	var namespace string
	flag.StringVar(&namespace, "n", "default", "")
	flag.StringVar(&namespace, "namespace", "default", "")

	c := KubernetesClientset()

	utils.PrintHeader()
	if list, err := corev1.Pods(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			utils.Print("api/v1", "Pod", p.Namespace, p.Name)
		}
	}

	if list, err := extensionsv1beta1.Deployments(c.Clientset, namespace); err == nil {
		for _, dp := range list.Items {
			utils.Print("extensions/v1beta1", "Deployment", dp.Namespace, dp.Name)
		}
	}
}
