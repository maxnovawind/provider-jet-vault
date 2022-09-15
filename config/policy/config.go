package policy

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the mount group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("vault_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "rbac"
		r.Kind = "Policy"
	})
}
