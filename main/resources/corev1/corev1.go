package corev1 

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Pods(c *kubernetes.Clientset, namespace string) (*corev1.PodList, error) {
	return c.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}

func ConfigMaps(c *kubernetes.Clientset, namespace string) (*corev1.ConfigMapList, error) {
	return c.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
}

func Endpoints(c *kubernetes.Clientset, namespace string) (*corev1.EndpointsList, error) {
	return c.CoreV1().Endpoints(namespace).List(metav1.ListOptions{})
}

func LimitRanges(c *kubernetes.Clientset, namespace string) (*corev1.LimitRangeList, error) {
	return c.CoreV1().LimitRanges(namespace).List(metav1.ListOptions{})
}

func PersistentVolumeClaims(c *kubernetes.Clientset, namespace string) (*corev1.PersistentVolumeClaimList, error) {
	return c.CoreV1().PersistentVolumeClaims(namespace).List(metav1.ListOptions{})
}

func PersistentVolumes(c *kubernetes.Clientset, namespace string) (*corev1.PersistentVolumeList, error) {
	return c.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
}

func ReplicationControllers(c *kubernetes.Clientset, namespace string) (*corev1.ReplicationControllerList, error) {
	return c.CoreV1().ReplicationControllers(namespace).List(metav1.ListOptions{})
}

func ResourceQuotas(c *kubernetes.Clientset, namespace string) (*corev1.ResourceQuotaList, error) {
	return c.CoreV1().ResourceQuotas(namespace).List(metav1.ListOptions{})
}

func Secrets(c *kubernetes.Clientset, namespace string) (*corev1.SecretList, error) {
	return c.CoreV1().Secrets(namespace).List(metav1.ListOptions{})
}

func ServiceAccounts(c *kubernetes.Clientset, namespace string) (*corev1.ServiceAccountList, error) {
	return c.CoreV1().ServiceAccounts(namespace).List(metav1.ListOptions{})
}

func Services(c *kubernetes.Clientset, namespace string) (*corev1.ServiceList, error) {
	return c.CoreV1().Services(namespace).List(metav1.ListOptions{})
}
