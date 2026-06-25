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
	"hma_oss_conf_editor/consts"
	"hma_oss_conf_editor/objects"
)

// The functions in this file is only valid for recent Zygisk module versions

// All of settings will be ignored when the config version was set to no settings
func DeleteAllSettings(config *objects.JsonConfig) {
	config.ConfigVersion = consts.ConfigVersionNoSettings
}

// All of settings will be applied when the config version was set to current version
func RestoreAllSettings(config *objects.JsonConfig) {
	config.ConfigVersion = consts.ConfigVersionCurrent
}
