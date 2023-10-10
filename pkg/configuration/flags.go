/*
Copyright © 2023 Doppler <support@doppler.com>

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
package configuration

import (
	"slices"

	"github.com/DopplerHQ/cli/pkg/models"
)

func GetFlag(flag string) bool {
	flags := configContents.Flags
	switch flag {
	case models.FlagAnalytics:
		if flags.Analytics != nil {
			return *flags.Analytics
		}
		return GetFlagDefault(models.FlagAnalytics)
	}

	return false
}

func SetFlag(flag string, enable bool) {
	switch flag {
	case models.FlagAnalytics:
		configContents.Flags.Analytics = &enable
	}
	writeConfig(configContents)
}

func GetFlagDefault(flag string) bool {
	switch flag {
	case models.FlagAnalytics:
		return true
	}

	return false
}

func IsValidFlag(flag string) bool {
	flags := models.GetFlags()
	return slices.Contains(flags, flag)
}

func IsAnalyticsEnabled() bool {
	return GetFlag(models.FlagAnalytics)
}
