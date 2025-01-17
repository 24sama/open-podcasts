/*
Copyright 2022 The open-podcasts Authors. All rights reserved.

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

package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	osf2fv1alpha1 "github.com/linuxsuren/open-podcasts/api/v1alpha1"
)

// ProfileReconciler reconciles a Profile object
type ProfileReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=osf2f.my.domain,resources=profiles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=osf2f.my.domain,resources=profiles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=osf2f.my.domain,resources=profiles/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Profile object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *ProfileReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, err error) {
	_ = log.FromContext(ctx)

	profile := &osf2fv1alpha1.Profile{}
	if err = r.Client.Get(ctx, req.NamespacedName, profile); err != nil {
		err = client.IgnoreNotFound(err)
		return
	}

	err = r.handleLaterPlayList(profile)
	return
}

func (r *ProfileReconciler) handleLaterPlayList(profile *osf2fv1alpha1.Profile) (err error) {
	playList := profile.Spec.LaterPlayList
	if len(playList) == 0 {
		return
	}

	needToUpdate := false
	for i, _ := range playList {
		item := &playList[i]
		if item.DisplayName != "" {
			continue
		}

		episode := &osf2fv1alpha1.Episode{}
		if err = r.Client.Get(context.Background(), types.NamespacedName{
			Namespace: profile.Namespace,
			Name:      item.Name,
		}, episode); err == nil {
			// TODO need to consider the case when episode does not exist
			item.DisplayName = strings.TrimSpace(episode.Spec.Title)
			needToUpdate = true
		}
	}

	if needToUpdate {
		err = r.Client.Update(context.Background(), profile)
	}
	return
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProfileReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&osf2fv1alpha1.Profile{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
