package resources 

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/api/extensions/v1beta1",
)

func Deployments(c *kubernetes.Clientset, namespace string) (v1beta1.Deployments, error) {
	return c.ExtensionsV1beta1().Deployments(namespace).List(metav1.ListOptions{})
}

func Daemonset(c *kubernetes.Clientset, namespace string) (v1beta1.DaemonSets, error) {
	return c.ExtensionsV1beta1().DaemonSets(namespace).List(metav1.ListOptions{})
}

func Ingresses(c *kubernetes.Clientset, namespace string) (v1beta1.Ingresses, error) {
	return c.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{})
}

func NetworkPolicies(c *kubernetes.Clientset, namespace string) {
	return c.ExtensionsV1beta1().NetworkPolicies(namespace).List(metav1.ListOptions{})
}

func PodSecurityPolicies(c *kubernetes.Clientset, namespace string)(v1beta1.PodSecurityPolicies, error) {
	return c.ExtensionsV1beta1().PodSecurityPolicies(namespace).List(metav1.ListOptions{})
}

func ReplicaSets(c *kubernetes.Clientset, namepsace string)(v1beta1.ReplicaSets, error) {
	return c.ExtensionsV1beta1().ReplicaSets(namespace).List(metav1.ListOptions{})
}
