/*
Copyright 2021 The Crossplane Authors.

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
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	role "github.com/maxnovawind/provider-jet-vault/internal/controller/approle/role"
	secret "github.com/maxnovawind/provider-jet-vault/internal/controller/approle/secret"
	mount "github.com/maxnovawind/provider-jet-vault/internal/controller/auth/mount"
	providerconfig "github.com/maxnovawind/provider-jet-vault/internal/controller/providerconfig"
	policy "github.com/maxnovawind/provider-jet-vault/internal/controller/rbac/policy"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		role.Setup,
		secret.Setup,
		mount.Setup,
		providerconfig.Setup,
		policy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
