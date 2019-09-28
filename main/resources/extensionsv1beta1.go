package resources

import (
	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Deployments(c *kubernetes.Clientset, namespace string) (*v1beta1.DeploymentList, error) {
	return c.ExtensionsV1beta1().Deployments(namespace).List(metav1.ListOptions{})
}

func Daemonset(c *kubernetes.Clientset, namespace string) (*v1beta1.DaemonSetList, error) {
	return c.ExtensionsV1beta1().DaemonSets(namespace).List(metav1.ListOptions{})
}

func Ingresses(c *kubernetes.Clientset, namespace string) (*v1beta1.IngressList, error) {
	return c.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{})
}

func NetworkPolicies(c *kubernetes.Clientset, namespace string) (*v1beta1.NetworkPolicyList, error) {
	return c.ExtensionsV1beta1().NetworkPolicies(namespace).List(metav1.ListOptions{})
}

func PodSecurityPolicies(c *kubernetes.Clientset, namespace string) (*v1beta1.PodSecurityPolicyList, error) {
	return c.ExtensionsV1beta1().PodSecurityPolicies().List(metav1.ListOptions{})
}

func ReplicaSets(c *kubernetes.Clientset, namespace string) (*v1beta1.ReplicaSetList, error) {
	return c.ExtensionsV1beta1().ReplicaSets(namespace).List(metav1.ListOptions{})
}
