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

package consts

var PackagesShouldNotHide = []string{
	"android",
	"android.media",
	"android.uid.system",
	"android.uid.shell",
	"android.uid.systemui",
	"com.android.permissioncontroller",
	"com.android.providers.downloads",
	"com.android.providers.downloads.ui",
	"com.android.providers.media",
	"com.android.providers.media.module",
	"com.android.providers.settings",
	"com.google.android.providers.media.module",
}
