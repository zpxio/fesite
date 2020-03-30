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

package model

import (
	"github.com/apex/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Recipe struct {
	Name             string
	Permalink        string
	Intro            string
	IngredientGroups []IngredientGroup `yaml:"ingredients"`
	InstructionSets  []InstructionSet  `yaml:"directions"`
}

func ParseFile(path string) (*Recipe, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Errorf("Failed to read file [%s]: %s", path, err)
		return nil, err
	}

	return Parse(data)
}

func Parse(data []byte) (*Recipe, error) {
	r := Recipe{}

	err := yaml.Unmarshal(data, &r)

	if err != nil {
		return nil, err
	}

	log.Debugf("Parsed recipe data: %+v", r)

	r.InferDefaults()
	r.Normalize()

	return &r, nil
}

func (r *Recipe) Normalize() {

}

func (r *Recipe) InferDefaults() {

}
