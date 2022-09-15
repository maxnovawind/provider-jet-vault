package auth

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the mount group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("vault_auth_backend", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "auth"
		r.Kind = "Mount"
	})
}
