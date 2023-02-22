package cloud

import "github.com/hashicorp/terraform/internal/modsdir"

// ModuleVersion is the official interface for the modules
type ModuleVersion struct {
	Module  string `json:"module"`
	Version string `json:"version"`
	Source  string `json:"source"`
}

type WorkspaceModules struct {
	Modules []ModuleVersion `json:"modules"`
}

func NewWorkspaceModules(manifest modsdir.Manifest) WorkspaceModules {
	modules := make([]ModuleVersion, len(manifest))
	for _, record := range manifest {
		modules = append(modules, ModuleVersion{
			Module:  record.Key,
			Version: record.VersionStr,
			Source:  record.SourceAddr,
		})
	}
	return WorkspaceModules{Modules: modules}
}
