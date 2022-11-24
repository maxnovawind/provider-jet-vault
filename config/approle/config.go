package approle

import (
	"github.com/crossplane/terrajet/pkg/config"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the mount group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("vault_approle_auth_backend_role", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "approle"
		r.Kind = "Role"
	})
	p.AddResourceConfigurator("vault_approle_auth_backend_role_secret_id", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "approle"
		r.Kind = "Secret"
		r.References["role_name"] = config.Reference{
			Type: "Role",
		}
	})
}
