/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"github.com/robfig/cron/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1alpha1 "github.com/amasotti/pod-rotator-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
)

// CustomPodRotatorReconciler reconciles a CustomPodRotator object
type CustomPodRotatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
//
// +kubebuilder:rbac:groups=apps.tonihacks.com,resources=custompodrotators,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.tonihacks.com,resources=custompodrotators/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps.tonihacks.com,resources=custompodrotators/finalizers,verbs=get,update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;update
func (r *CustomPodRotatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// First try to fetch the CustomPodRotator instance
	rotator := &appsv1alpha1.CustomPodRotator{}
	if err := r.Get(ctx, req.NamespacedName, rotator); err != nil {
		logger.Error(err, "unable to fetch CustomPodRotator. Has it been deleted?")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Check if the CustomPodRotator has a schedule
	now := time.Now()
	schedule, err := cron.ParseStandard(rotator.Spec.Schedule)
	if err != nil {
		logger.Error(err, "unable to parse schedule. Check your cron expression")
		return ctrl.Result{}, err
	}

	// Check if it's time to restart the pods
	nextRun := schedule.Next(rotator.Status.LastRestartTime.Time)
	if now.Before(nextRun) {
		logger.Info("Not time to restart yet. I'll go sleep again", "nextRun", nextRun)
		return ctrl.Result{RequeueAfter: nextRun.Sub(now)}, nil
	}

	// Trigger the restart
	deployment := &appsv1.Deployment{}
	if err := r.Get(ctx, client.ObjectKey{
		Namespace: rotator.Namespace,
		Name:      rotator.Spec.TargetDeployment,
	}, deployment); err != nil {
		logger.Error(err, "unable to fetch target deployment. Check the passed name")
		return ctrl.Result{}, err
	}

	// Update the deployment annotations
	annotations := deployment.Spec.Template.Annotations
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations["custompodrotator.tonihacks.com/restarted-at"] = now.Format(time.RFC3339)
	deployment.Spec.Template.Annotations = annotations

	if err := r.Update(ctx, deployment); err != nil {
		logger.Error(err, "Error while trying to update deployment")
		return ctrl.Result{}, err
	}

	rotator.Status.LastRestartTime = metav1.Time{Time: now}
	if err := r.Status().Update(ctx, rotator); err != nil {
		logger.Error(err, "Error while trying to update status")
		return ctrl.Result{}, err
	}

	logger.Info("Deployment restarted successfully", "deployment", rotator.Spec.TargetDeployment)

	return ctrl.Result{RequeueAfter: schedule.Next(now).Sub(now)}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CustomPodRotatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.CustomPodRotator{}).
		Complete(r)
}
