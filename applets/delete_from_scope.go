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
	"maps"
	"slices"
)

func DeleteFromScope(config *objects.JsonConfig, packageNames []string) {
	total := len(packageNames)
	if total < 1 {
		fmt.Println("Package name list is empty")
		return
	}

	if len(config.Scope) < 1 {
		fmt.Println("Scope list is empty")
		return
	}

	maps.DeleteFunc(
		config.Scope,
		func(key string, value *objects.AppConfig) bool {
			val := slices.Contains(packageNames, key)
			if val {
				fmt.Printf("Removing %v from scope\n", key)
			}

			return val
		},
	)
}
