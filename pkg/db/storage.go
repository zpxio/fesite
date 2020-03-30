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

package db

import (
	"github.com/apex/log"
	"github.com/zpxio/fesite/pkg/config"
	"github.com/zpxio/fesite/pkg/model"
	"os"
	"path/filepath"
	"strings"
)

var recipeStore = make(map[string]*model.Recipe)

func Rescan() {
	rs := make(map[string]*model.Recipe)

	root := config.SiteRoot()

	// Scan for recipe.yml files.
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(info.Name(), ".recipe.yml") {
			log.Infof("Reading recipe: %s", path)
			r, err := model.ParseFile(path)
			if err != nil {
				log.Errorf("Failed to parse recipe [%s]: %s", path, err)
			} else {
				// Add the recipe to the store
				rs[r.Permalink] = r
				log.Infof("Loaded recipe [%s] @ %s", r.Permalink, path)
			}
		}

		return nil
	})

	if err != nil {
		log.Errorf("Failed to scan for recipe files: %s", err)
	}

	// Replace the global store
	recipeStore = rs
}

func Retrieve(permalink string) *model.Recipe {
	r, ok := recipeStore[permalink]
	if ok {
		return r
	}

	return nil
}
