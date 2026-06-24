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
	"encoding/json"
	"hma_oss_conf_editor/objects"
	"io"
	"log"
	"os"
)

func ParseConfig(filePath string) *objects.JsonConfig {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("cannot open file %v due to %v", filePath, err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read file %v due to %v", filePath, err)
	}

	parsed := objects.JsonConfig{
		WebViewProtection: true,
	}
	err = json.Unmarshal(content, &parsed)
	if err != nil {
		log.Fatalf("cannot parse %v due to %v", filePath, err)
	}

	return &parsed
}
