package transcodingapi

// ProviderName is a custom string type with the name of the provider
type ProviderName string

// ProviderNames holds many ProviderName instances
type ProviderNames []ProviderName

// Description fully describes a provider.
type ProviderDescription struct {
	Name         string               `json:"name"`
	Capabilities ProviderCapabilities `json:"capabilities"`
	Health       ProviderHealth       `json:"health"`
	Enabled      bool                 `json:"enabled"`
}

// Capabilities describes the available features in the provider.
type ProviderCapabilities struct {
	InputFormats  []string `json:"input"`
	OutputFormats []string `json:"output"`
	Destinations  []string `json:"destinations"`
}

// Health describes the current health status of the provider.
type ProviderHealth struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}
