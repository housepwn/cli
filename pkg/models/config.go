/*
Copyright © 2019 Doppler <support@doppler.com>

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
package models

import "time"

// ConfigFile structure of the config file
type ConfigFile struct {
	Scoped       map[string]FileScopedOptions `yaml:"scoped"`
	VersionCheck VersionCheck                 `yaml:"version-check"`
}

// FileScopedOptions config options
type FileScopedOptions struct {
	Token          string `json:"token,omitempty" yaml:"token,omitempty"`
	APIHost        string `json:"api-host,omitempty" yaml:"api-host,omitempty"`
	DashboardHost  string `json:"dashboard-host,omitempty" yaml:"dashboard-host,omitempty"`
	VerifyTLS      string `json:"verify-tls,omitempty" yaml:"verify-tls,omitempty"`
	EnclaveProject string `json:"enclave.project,omitempty" yaml:"enclave.project,omitempty"`
	EnclaveConfig  string `json:"enclave.config,omitempty" yaml:"enclave.config,omitempty"`
}

// VersionCheck info about the last check for the latest cli version
type VersionCheck struct {
	LatestVersion string    `yaml:"latest-version,omitempty"`
	CheckedAt     time.Time `yaml:"checked-at,omitempty"`
}

// ScopedOptions options with their scope
type ScopedOptions struct {
	Token          ScopedOption `json:"token,omitempty" yaml:"token,omitempty"`
	APIHost        ScopedOption `json:"api-host,omitempty" yaml:"api-host,omitempty"`
	DashboardHost  ScopedOption `json:"dashboard-host,omitempty" yaml:"dashboard-host,omitempty"`
	VerifyTLS      ScopedOption `json:"verify-tls,omitempty" yaml:"verify-tls,omitempty"`
	EnclaveProject ScopedOption `json:"enclave.project,omitempty" yaml:"enclave.project,omitempty"`
	EnclaveConfig  ScopedOption `json:"enclave.config,omitempty" yaml:"enclave.config,omitempty"`
}

// ScopedOption value and its scope
type ScopedOption struct {
	Value  string `json:"value"`
	Scope  string `json:"scope"`
	Source string `json:"source"`
}

// Source where the value came from
type Source int

// the source of the value
const (
	FlagSource Source = iota
	ConfigFileSource
	EnvironmentSource
	DefaultValueSource
)

func (s Source) String() string {
	return [...]string{"Flag", "Config File", "Environment", "Default Value"}[s]
}

// ConfigOptions all supported options
var ConfigOptions = []string{
	"token",
	"api-host",
	"dashboard-host",
	"verify-tls",
	"enclave.project",
	"enclave.config",
}

// Pairs get the pairs for the given config
func Pairs(conf FileScopedOptions) map[string]string {
	return map[string]string{
		"token":           conf.Token,
		"api-host":        conf.APIHost,
		"dashboard-host":  conf.DashboardHost,
		"verify-tls":      conf.VerifyTLS,
		"enclave.project": conf.EnclaveProject,
		"enclave.config":  conf.EnclaveConfig,
	}
}

// ScopedPairs get the pairs for the given scoped config
func ScopedPairs(conf *ScopedOptions) map[string]*ScopedOption {
	return map[string]*ScopedOption{
		"token":           &conf.Token,
		"api-host":        &conf.APIHost,
		"dashboard-host":  &conf.DashboardHost,
		"verify-tls":      &conf.VerifyTLS,
		"enclave.project": &conf.EnclaveProject,
		"enclave.config":  &conf.EnclaveConfig,
	}
}

// EnvPairs get the scoped config pairs for each environment variable
func EnvPairs(conf *ScopedOptions) map[string]*ScopedOption {
	return map[string]*ScopedOption{
		"DOPPLER_TOKEN":          &conf.Token,
		"DOPPLER_API_HOST":       &conf.APIHost,
		"DOPPLER_DASHBOARD_HOST": &conf.DashboardHost,
		"DOPPLER_VERIFY_TLS":     &conf.VerifyTLS,
		"ENCLAVE_PROJECT":        &conf.EnclaveProject,
		"ENCLAVE_CONFIG":         &conf.EnclaveConfig,
	}
}
