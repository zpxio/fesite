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
	"github.com/zpxio/fesite/pkg/db"
	"github.com/zpxio/fesite/pkg/model"
	"strings"
)

func AttachRecipeHandler(d *Dispatcher) {
	d.engine.GET("/recipe/:rid", RecipeData)
}

func RecipeData(c *gin.Context) {

	r := db.Retrieve(c.Param("rid"))
	format := c.DefaultQuery("format", "html")
	log.Infof("Rendering [%s] with format %s", r.Permalink, format)

	switch strings.ToLower(format) {
	case "yaml":
		{
			RecipeDataYaml(c, r)
		}
	case "json":
		{
			RecipeDataJson(c, r)
		}
	default:
		RecipeDataHtml(c, r)
	}
}

func RecipeDataHtml(c *gin.Context, r *model.Recipe) {
	log.Debugf("Rendering [%s] as HTML", r.Permalink)
	c.YAML(200, r)
}

func RecipeDataJson(c *gin.Context, r *model.Recipe) {
	log.Debugf("Rendering [%s] as YAML", r.Permalink)
	c.JSON(200, r)
}

func RecipeDataYaml(c *gin.Context, r *model.Recipe) {
	log.Debugf("Rendering [%s] as YAML", r.Permalink)
	c.YAML(200, r)
}
