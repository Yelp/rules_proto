package protogen

import (
	"html/template"
)

// ProtoPlugin loosely represents the starlark ProtoRuleInfo provider.
type ProtoRule struct {
	// Name of the rule
	Name string

	// Package of the rule (label package)
	Package string

	// Kind of the rule (proto|grpc)
	Kind string

	// Description
	Doc string

	// All templates parsed together
	Templates *template.Template

	// ImplementationFilename of the rule (full path)
	ImplementationFilename string

	// Filename of the implementation template
	ImplementationTmpl string

	// Filename of the workspace template
	WorkspaceExampleTmpl string

	// WorkspaceExampleFilename of the rule (full path)
	WorkspaceExampleFilename string

	// BuildExampleFilename of the rule (full path)
	BuildExampleFilename string

	// Filename of the BuildExample template
	BuildExampleTmpl string

	// MarkdownFilename of the rule (full path)
	MarkdownFilename string

	// Filename of the Markdown template
	MarkdownTmpl string

	// DepsFilename of the rule (full path)
	DepsFilename string

	// Filename of the Deps template
	DepsTmpl string

	// List of attributes
	Attrs []*Attr

	// List of plugins
	Plugins []*ProtoPlugin

	// Not expected to be functional
	Experimental bool

	// Bazel build flags required / suggested
	Flags []*Flag

	// Additional CI-specific env vars in the form "K=V"
	PresubmitEnvVars map[string]string

	// Platforms for which to skip testing this rule, overrides language level
	// The special value 'all' will skip app platforms
	SkipTestPlatforms []string

	// Flag indicating if the merge_directories flag should be set to false for
	// the generated rule
	SkipDirectoriesMerge bool

	// The language this rule belongs to
	Language *ProtoLanguage
}

// ProtoPlugin represents the starlark ProtoPluginInfo provider.
type ProtoPlugin struct {
	Name                       string
	Label                      string
	Options                    []string
	Outputs                    []string
	OutputDirectory            string
	Tool                       string
	ToolExecutable             string
	UseBuiltInShellEnvironment bool
	ProtocPluginName           string
	Exclusions                 []string
	SeparateOptionsFlag        bool
	Deps                       []*ProtoDependency
}

// ProtoDependency represents the starlark ProtoDependencyInfo provider.
type ProtoDependency struct {
	Name           string
	RepositoryRule string
	Urls           []string
	Sha256         string
	StripPrefix    string
}

// ProtoLanguage represents the starlark ProtoLanguageInfo provider.
type ProtoLanguage struct {
	Name  string
	Rules []string
	Deps  []*ProtoDependency

	// MarkdownFilename of the rule (full path)
	MarkdownFilename string

	// Filename of the Markdown template
	MarkdownTmpl string

	// RulesFilename (full path)
	RulesFilename string

	// Filename of the Rules template
	RulesTmpl string
	
	// Templates is populated by the NewProtoLanguageFromJSON constructor.
	Templates *template.Template
}

// Flag captures information about a bazel build flag.
type Flag struct {
	Category    string
	Name        string
	Value       string
	Description string
}

type Attr struct {
	Name      string
	Type      string
	Default   string
	Doc       string
	Mandatory bool
}
