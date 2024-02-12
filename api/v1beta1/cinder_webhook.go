/*
Copyright 2022.

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

// Generated by:
//
// operator-sdk create webhook --group cinder --version v1beta1 --kind Cinder --programmatic-validation --defaulting
//

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/util"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// CinderDefaults -
type CinderDefaults struct {
	APIContainerImageURL       string
	BackupContainerImageURL    string
	SchedulerContainerImageURL string
	VolumeContainerImageURL    string
	DBPurgeAge                 int
	DBPurgeSchedule            string
}

var cinderDefaults CinderDefaults

// log is for logging in this package.
var cinderlog = logf.Log.WithName("cinder-resource")

// SetupDefaults - initialize Cinder spec defaults for use with either internal or external webhooks
func SetupDefaults() {
	cinderDefaults = CinderDefaults{
		APIContainerImageURL:       util.GetEnvVar("RELATED_IMAGE_CINDER_API_IMAGE_URL_DEFAULT", CinderAPIContainerImage),
		BackupContainerImageURL:    util.GetEnvVar("RELATED_IMAGE_CINDER_BACKUP_IMAGE_URL_DEFAULT", CinderBackupContainerImage),
		SchedulerContainerImageURL: util.GetEnvVar("RELATED_IMAGE_CINDER_SCHEDULER_IMAGE_URL_DEFAULT", CinderSchedulerContainerImage),
		VolumeContainerImageURL:    util.GetEnvVar("RELATED_IMAGE_CINDER_VOLUME_IMAGE_URL_DEFAULT", CinderVolumeContainerImage),
		DBPurgeAge:                 DBPurgeDefaultAge,
		DBPurgeSchedule:            DBPurgeDefaultSchedule,
	}

	cinderlog.Info("Cinder defaults initialized", "defaults", cinderDefaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *Cinder) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-cinder-openstack-org-v1beta1-cinder,mutating=true,failurePolicy=fail,sideEffects=None,groups=cinder.openstack.org,resources=cinders,verbs=create;update,versions=v1beta1,name=mcinder.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Cinder{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Cinder) Default() {
	cinderlog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this Cinder spec
func (spec *CinderSpec) Default() {
	if spec.CinderAPI.ContainerImage == "" {
		spec.CinderAPI.ContainerImage = cinderDefaults.APIContainerImageURL
	}

	if spec.CinderBackup.ContainerImage == "" {
		spec.CinderBackup.ContainerImage = cinderDefaults.BackupContainerImageURL
	}

	if spec.CinderScheduler.ContainerImage == "" {
		spec.CinderScheduler.ContainerImage = cinderDefaults.SchedulerContainerImageURL
	}

	for index, cinderVolume := range spec.CinderVolumes {
		if cinderVolume.ContainerImage == "" {
			cinderVolume.ContainerImage = cinderDefaults.VolumeContainerImageURL
		}
		// This is required, as the loop variable is a by-value copy
		spec.CinderVolumes[index] = cinderVolume
	}

	if spec.DBPurge.Age == 0 {
		spec.DBPurge.Age = cinderDefaults.DBPurgeAge
	}
	if spec.DBPurge.Schedule == "" {
		spec.DBPurge.Schedule = cinderDefaults.DBPurgeSchedule
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-cinder-openstack-org-v1beta1-cinder,mutating=false,failurePolicy=fail,sideEffects=None,groups=cinder.openstack.org,resources=cinders,verbs=create;update,versions=v1beta1,name=vcinder.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Cinder{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Cinder) ValidateCreate() error {
	cinderlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Cinder) ValidateUpdate(old runtime.Object) error {
	cinderlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Cinder) ValidateDelete() error {
	cinderlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
