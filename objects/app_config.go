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

package objects

type AppConfig struct {
	UseWhitelist                    bool     `json:"useWhitelist"`
	ExcludeSystemApps               bool     `json:"excludeSystemApps"`
	HideInstallationSource          bool     `json:"hideInstallationSource"`
	HideSystemInstallationSource    bool     `json:"hideSystemInstallationSource"`
	ExcludeTargetInstallationSource bool     `json:"excludeTargetInstallationSource"`
	InvertActivityLaunchProtection  bool     `json:"invertActivityLaunchProtection"`
	ExcludeVoldIsolation            bool     `json:"excludeVoldIsolation"`
	RestrictedZygotePermissions     []int    `json:"restrictedZygotePermissions"`
	ApplyTemplates                  []string `json:"applyTemplates"`
	ApplyPresets                    []string `json:"applyPresets"`
	ApplySettingTemplates           []string `json:"applySettingTemplates"`
	ApplySettingsPresets            []string `json:"applySettingsPresets"`
	ExtraAppList                    []string `json:"extraAppList"`
	ExtraOppositeAppList            []string `json:"extraOppositeAppList"`
}
