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
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/zpxio/fesite/pkg/config"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func AttachStyling(d *Dispatcher) {
	d.engine.GET("/css/unified.css", UnifiedStyling)
}

func UnifiedStyling(c *gin.Context) {

	// Scan the directory for css files
	siteRoot := config.SiteRoot()
	cssFiles, err := ScanForExtension(config.SiteRoot(), "css")
	if err != nil {
		log.Errorf("Failed to scan site directory for stylesheets: %s", err)
	}

	page := strings.Builder{}

	for _, fi := range cssFiles {
		cssFile := filepath.Join(siteRoot, fi.Name())
		data, err := ioutil.ReadFile(cssFile)
		if err != nil {
			log.Errorf("Failed to read site stylesheet [%s]: %s", fi.Name(), err)
		}

		_, wErr := page.Write(data)
		if wErr != nil {
			log.Errorf("Error while composing data [%s]: %s", fi.Name(), wErr)
		}
		_, wErr = page.WriteString("\n")
		if wErr != nil {
			log.Errorf("Error while delimiting data [%s]: %s", fi.Name(), wErr)
		}
	}

	content := page.String()

	c.String(http.StatusOK, content)
}
