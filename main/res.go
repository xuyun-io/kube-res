package main

import (
	"flag"

	"github.com/xuyun-io/kube-res/main/resources/appsv1"
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

	if list, err := corev1.ConfigMaps(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			utils.Print("api/v1", "ConfigMap", p.Namespace, p.Name)
		}
	}
	/*
		if list, err := corev1.Endpoints(c.Clientset, namespace); err == nil {
			for _, p := range list.Items {
				utils.Print("api/v1", "Endpoints", p.Namespace, p.Name)
			}
		}

		if list, err := corev1.LimitRanges(c.Clientset, namespace); err == nil {
			for _, p := range list.Items {
				utils.Print("api/v1", "LimitRange", p.Namespace, p.Name)
			}
		}
	*/
	if list, err := corev1.PersistentVolumeClaims(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			utils.Print("api/v1", "PersistentVolumeClaim", p.Namespace, p.Name)
		}
	}

	if list, err := corev1.PersistentVolumes(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			utils.Print("api/v1", "PersistentVolume", p.Namespace, p.Name)
		}
	}

	if list, err := corev1.ReplicationControllers(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			utils.Print("api/v1", "ReplicationController", p.Namespace, p.Name)
		}
	}
	/*
		if list, err := corev1.ResourceQuotas(c.Clientset, namespace); err == nil {
			for _, p := range list.Items {
				utils.Print("api/v1", "ResourceQuota", p.Namespace, p.Name)
			}
		}
	*/

	if list, err := corev1.Secrets(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			utils.Print("api/v1", "Secrets", p.Namespace, p.Name)
		}
	}

	if list, err := corev1.ServiceAccounts(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			utils.Print("api/v1", "ServiceAccount", p.Namespace, p.Name)
		}
	}

	if list, err := corev1.Services(c.Clientset, namespace); err == nil {
		for _, p := range list.Items {
			utils.Print("api/v1", "Service", p.Namespace, p.Name)
		}
	}

	utils.Printline()

	if list, err := extensionsv1beta1.Deployments(c.Clientset, namespace); err == nil {
		for _, dp := range list.Items {
			utils.Print("extensions/v1beta1", "Deployment", dp.Namespace, dp.Name)
		}
	}

	if list, err := extensionsv1beta1.Daemonsets(c.Clientset, namespace); err == nil {
		for _, dp := range list.Items {
			utils.Print("extensions/v1beta1", "DaemonSet", dp.Namespace, dp.Name)
		}
	}

	if list, err := extensionsv1beta1.ReplicaSets(c.Clientset, namespace); err == nil {
		for _, dp := range list.Items {
			utils.Print("extensions/v1beta1", "ReplicaSet", dp.Namespace, dp.Name)
		}
	}

	utils.Printline()

	if list, err := appsv1.StatefulSets(c.Clientset, namespace); err == nil {
		for _, dp := range list.Items {
			utils.Print("apps/v1", "StatefulSet", dp.Namespace, dp.Name)
		}
	}
}
