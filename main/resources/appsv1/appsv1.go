package appsv1

import (
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Deployments(c *kubernetes.Clientset, namespace string) (*v1.DeploymentList, error) {
	// return c.ExtensionsV1beta1().Deployments(namespace).List(metav1.ListOptions{})
	return c.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
}

func DaemonSets(c *kubernetes.Clientset, namespace string) (*v1.DaemonSetList, error) {
	return c.AppsV1().DaemonSets(namespace).List(metav1.ListOptions{})
}

func ReplicaSets(c *kubernetes.Clientset, namespace string) (*v1.ReplicaSetList, error) {
	return c.AppsV1().ReplicaSets(namespace).List(metav1.ListOptions{})
}

func StatefulSets(c *kubernetes.Clientset, namespace string) (*v1.StatefulSetList, error) {
	return c.AppsV1().StatefulSets(namespace).List(metav1.ListOptions{})
}
