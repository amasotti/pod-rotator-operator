package controller

import (
	"context"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	appsv1alpha1 "github.com/amasotti/pod-rotator-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

var _ = Describe("CustomPodRotator Controller", func() {
	Context("When scheduling a restart", func() {
		It("Should update deployment annotations", func() {
			ctx := context.Background()

			// Create deployment first
			deployment := &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-deploy",
					Namespace: "default",
				},
				Spec: appsv1.DeploymentSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "test-deploy",
						},
					},
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"app": "test-deploy",
							},
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Name:  "nginx",
								Image: "nginx:1.25",
							}},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, deployment)).To(Succeed())

			// Create rotator with immediate trigger
			rotator := &appsv1alpha1.CustomPodRotator{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-rotator",
					Namespace: "default",
				},
				Spec: appsv1alpha1.CustomPodRotatorSpec{
					TargetDeployment: "test-deploy",
					Schedule:         "* * * * *",
				},
			}
			Expect(k8sClient.Create(ctx, rotator)).To(Succeed())

			// Manually trigger reconciliation
			reconciler := &CustomPodRotatorReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			// Trigger reconciliation manually
			_, err := reconciler.Reconcile(ctx, ctrl.Request{
				NamespacedName: types.NamespacedName{
					Name:      "test-rotator",
					Namespace: "default",
				},
			})
			Expect(err).NotTo(HaveOccurred())

			// Check the result
			updatedDeployment := &appsv1.Deployment{}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, types.NamespacedName{
					Name:      "test-deploy",
					Namespace: "default",
				}, updatedDeployment)
				if err != nil {
					return false
				}
				return updatedDeployment.Spec.Template.Annotations != nil &&
					updatedDeployment.Spec.Template.Annotations["custompodrotator.tonihacks.com/restarted-at"] != ""
			}, 5*time.Second, 100*time.Millisecond).Should(BeTrue())

			// Cleanup
			Expect(k8sClient.Delete(ctx, deployment)).To(Succeed())
			Expect(k8sClient.Delete(ctx, rotator)).To(Succeed())
		})
	})
})
