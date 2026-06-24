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
	"hma_oss_conf_editor/consts"
	"hma_oss_conf_editor/objects"
	"slices"
)

func AddToScope(
	config *objects.JsonConfig,
	packageName string,
	appConfig *objects.AppConfig,
	overwrite bool,
) bool {
	if !overwrite {
		if _, ok := config.Scope[packageName]; ok {
			fmt.Println("Scope exists")
			return false
		}
	}

	if slices.Contains(consts.PackagesShouldNotHide, packageName) {
		fmt.Println("This package should not be added as target")
		return false
	}

	config.Scope[packageName] = appConfig
	return true
}
