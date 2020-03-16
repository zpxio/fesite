/*
 * Copyright 2020 zpxio (Jeff Sharpe)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"io/ioutil"
	"os"
	"strings"
)

func ScanForExtension(dir string, ext string) ([]os.FileInfo, error) {
	// Make sure the extension starts with a dot
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	var found []os.FileInfo
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return found, err
	}

	for _, fi := range files {
		if !fi.IsDir() && strings.HasSuffix(fi.Name(), ext) {
			found = append(found, fi)
		}
	}

	return found, err
}
