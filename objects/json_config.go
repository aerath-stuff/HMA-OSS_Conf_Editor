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

type JsonConfig struct {
	ConfigVersion                   int                          `json:"configVersion"`
	DetailLog                       bool                         `json:"detailLog"`
	ErrorOnlyLog                    bool                         `json:"errorOnlyLog"`
	MaxLogSize                      int                          `json:"maxLogSize"`
	ForceMountData                  bool                         `json:"forceMountData"`
	DisableActivityLaunchProtection bool                         `json:"disableActivityLaunchProtection"`
	AltAppDataIsolation             bool                         `json:"altAppDataIsolation"`
	AltVoldAppDataIsolation         bool                         `json:"altVoldAppDataIsolation"`
	SkipSystemAppDataIsolation      bool                         `json:"skipSystemAppDataIsolation"`
	PackageQueryWorkaround          bool                         `json:"packageQueryWorkaround"`
	WebViewProtection               bool                         `json:"webViewProtection"`
	Templates                       map[string]*AppTemplate      `json:"templates"`
	SettingsTemplates               map[string]*SettingsTemplate `json:"settingsTemplates"`
	DisabledHooks                   []*DisabledHook              `json:"disabledHooks"`
	Scope                           map[string]*AppConfig        `json:"scope"`
}
