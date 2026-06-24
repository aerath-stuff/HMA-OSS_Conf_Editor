/*
 * HMA-OSS Conf Editor
 *
 * Copyright (c) 2026 Furkan Karcioglu <krc440002@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package applets

import (
	"fmt"
	"hma_oss_conf_editor/objects"
)

const workModeBlacklist = "Blacklist"
const workModeWhitelistExclude = "Whitelist + excluding system apps"
const workModeWhitelistNormal = "Whitelist"
const instSourceUserApps = "User apps"
const instSourceSystemApps = "System apps"
const instSourceExcludeCurrent = "Exclude current"
const instSourceNone = "None"

var template = `%v

Work mode: %v
Installation source spoofing: %v
Invert Activity launch protection: %v
Exclude vold isolation: %v
Restricted Zygote permissions: %v

Applied templates:
%v

Applied presets:
%v

Applied settings templates:
%v

Applied settings presets:
%v

Extra app list:
%v

Extra opposite app list:
%v
`

func ShowScopeDetails(packageName string, appConfig *objects.AppConfig) {
	workMode := workModeBlacklist
	if appConfig.UseWhitelist {
		if appConfig.ExcludeSystemApps {
			workMode = workModeWhitelistExclude
		} else {
			workMode = workModeWhitelistNormal
		}
	}

	instSource := instSourceNone
	if appConfig.HideInstallationSource {
		instSource = instSourceUserApps

		if appConfig.HideSystemInstallationSource {
			instSource += fmt.Sprintf(" + %v", instSourceSystemApps)
		}

		if appConfig.ExcludeTargetInstallationSource {
			instSource += fmt.Sprintf(" + %v", instSourceExcludeCurrent)
		}
	}

	applyTemplates := ""
	for _, k := range appConfig.ApplyTemplates {
		applyTemplates += fmt.Sprintln("-", k)
	}

	applyPresets := ""
	for _, k := range appConfig.ApplyPresets {
		applyPresets += fmt.Sprintln("-", k)
	}

	applySettingsTemplates := ""
	for _, k := range appConfig.ApplySettingTemplates {
		applySettingsTemplates += fmt.Sprintln("-", k)
	}

	applySettingsPresets := ""
	for _, k := range appConfig.ApplySettingsPresets {
		applySettingsPresets += fmt.Sprintln("-", k)
	}

	extraApps := ""
	for _, k := range appConfig.ExtraAppList {
		extraApps += fmt.Sprintln("-", k)
	}

	extraOppositeApps := ""
	for _, k := range appConfig.ExtraOppositeAppList {
		extraOppositeApps += fmt.Sprintln("-", k)
	}

	fmt.Printf(
		template,
		packageName,
		workMode,
		instSource,
		appConfig.InvertActivityLaunchProtection,
		appConfig.ExcludeVoldIsolation,
		appConfig.RestrictedZygotePermissions,
		applyTemplates,
		applyPresets,
		applySettingsTemplates,
		applySettingsPresets,
		extraApps,
		extraOppositeApps,
	)
}
