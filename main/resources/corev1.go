package resources

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Pod(c *kubernetes.Clientset, namespace string) (*corev1.PodList, error) {
	/*
		// c.CoreV1().ComponentStatuses()
		c.CoreV1().ConfigMaps()
		c.CoreV1().Endpoints()
		// c.CoreV1().Events()
		c.CoreV1().LimitRanges()
		// c.CoreV1().Namespaces()
		// c.CoreV1().Nodes()
		c.CoreV1().PersistentVolumeClaims()
		c.CoreV1().PersistentVolumes()
		// c.CoreV1().PodTemplates()
		c.CoreV1().ReplicationControllers()
		c.CoreV1().ResourceQuotas()
		c.CoreV1().Secrets()
		c.CoreV1().ServiceAccounts()
		c.CoreV1().Services()
	*/
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

func PersistentVolumeClaim(c *kubernetes.Clientset, namespace string) (*corev1.PersistentVolumeClaimList, error) {
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
