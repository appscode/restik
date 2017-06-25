package framework

import (
	"github.com/appscode/go/crypto/rand"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

func (f *Framework) DaemonSet() extensions.DaemonSet {
	return extensions.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.WithUniqSuffix("stash"),
			Namespace: f.namespace,
			Labels: map[string]string{
				"app": "stash-e2e",
			},
		},
		Spec: extensions.DaemonSetSpec{
			Template: f.PodTemplate(),
		},
	}
}

func (f *Framework) CreateDaemonSet(obj extensions.DaemonSet) error {
	_, err := f.kubeClient.ExtensionsV1beta1().DaemonSets(obj.Namespace).Create(&obj)
	return err
}

func (f *Framework) DeleteDaemonSet(meta metav1.ObjectMeta) error {
	return f.kubeClient.ExtensionsV1beta1().DaemonSets(meta.Namespace).Delete(meta.Name, &metav1.DeleteOptions{})
}

func (f *Framework) WaitForDaemonSetCondition(meta metav1.ObjectMeta, condition GomegaMatcher) {
	Eventually(func() *extensions.DaemonSet {
		obj, err := f.kubeClient.ExtensionsV1beta1().DaemonSets(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		return obj
	}).Should(condition)
}

func (f *Framework) WaitUntilDaemonSetCondition(meta metav1.ObjectMeta, condition GomegaMatcher) {
	Eventually(func() *extensions.DaemonSet {
		obj, err := f.kubeClient.ExtensionsV1beta1().DaemonSets(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		return obj
	}).ShouldNot(condition)
}
