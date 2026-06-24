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

package main

import (
	"hma_oss_conf_editor/applets"
	"hma_oss_conf_editor/objects"
)

func main() {
	cfg := applets.ParseConfig("example.json")

	// list scopes
	applets.ListScope(cfg)

	applets.ShowScopeDetails("com.zhenxi.hunter", cfg.Scope["com.zhenxi.hunter"])

	applets.AddToScope(cfg, "com.zhenxi.hunter", &objects.AppConfig{
		UseWhitelist: false,
	}, false)

	// try to delete one scope twice
	applets.DeleteFromScope(cfg, []string{"com.zhenxi.hunter"})
	applets.DeleteFromScope(cfg, []string{"com.zhenxi.hunter"})

	applets.AddToScope(cfg, "com.zhenxi.hunter", &objects.AppConfig{
		UseWhitelist: false,
	}, false)
	applets.AddToScope(cfg, "com.zhenxi.hunter", &objects.AppConfig{
		UseWhitelist: false,
	}, false)
}
