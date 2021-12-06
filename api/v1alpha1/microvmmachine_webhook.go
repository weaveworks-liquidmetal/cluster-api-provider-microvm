// Copyright 2021 Weaveworks or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MPL-2.0

package v1alpha1

import (
	"fmt"
	"reflect"

	"github.com/yitsushi/macpot"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var _ = logf.Log.WithName("microvmmachine-resource")

func (r *MicrovmMachine) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1alpha1-microvmmachine,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=microvmmachine,versions=v1alpha1,name=validation.microvmmachine.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1
// +kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1alpha1-microvmmachine,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=microvmmachine,versions=v1alpha1,name=default.microvmmachine.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

var (
	_ webhook.Validator = &MicrovmMachine{}
	_ webhook.Defaulter = &MicrovmMachine{}
)

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *MicrovmMachine) ValidateCreate() error {
	for i := range r.Spec.NetworkInterfaces {
		iface := r.Spec.NetworkInterfaces[i]

		if iface.GuestMAC == "" {
			mac, err := macpot.New(macpot.AsLocal())
			if err != nil {
				return fmt.Errorf("creating mac address client: %w", err)
			}

			iface.GuestMAC = mac.ToString()
		}
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *MicrovmMachine) ValidateDelete() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *MicrovmMachine) ValidateUpdate(old runtime.Object) error {
	var allErrs field.ErrorList

	previous, _ := old.(*MicrovmMachine)

	if !reflect.DeepEqual(r.Spec, previous.Spec) {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec"), "microvm machine spec is immutable"))
	}

	return aggregateObjErrors(r.GroupVersionKind().GroupKind(), r.Name, allErrs)
}

// Default satisfies the defaulting webhook interface.
func (r *MicrovmMachine) Default() {
}
