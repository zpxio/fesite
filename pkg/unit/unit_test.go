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

package unit

import (
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

type UnitSuite struct {
	suite.Suite
}

func TestUnitSuite(t *testing.T) {
	suite.Run(t, new(UnitSuite))
}

func (s *UnitSuite) TestYamlUnmarshal_Simple() {
	unitName := "mL"
	u := Unit{}

	// Pull the unit details to ensure that we don't accidentally test a new unit registration
	// as a side effect of resolution
	expUnit := Resolve(unitName)
	s.Require().Equal(Volume, expUnit.Measure)
	s.Require().Equal(Metric, expUnit.Family)

	err := yaml.Unmarshal([]byte(unitName), &u)

	s.Require().Nil(err, "Failed to unmarshal unit name")
	s.Equal(expUnit, u)
}

func (s *UnitSuite) TestYamlUnmarshal_New() {
	unitName := "foo"
	u := Unit{}

	// Pull the unit details to ensure that we don't accidentally test a new unit registration
	// as a side effect of resolution
	s.Require().NotContains(Dictionary, strings.ToLower(unitName), "Test unit already registered")

	err := yaml.Unmarshal([]byte(unitName), &u)

	s.Require().Nil(err, "Failed to unmarshal unit name")
	s.Equal(unitName, u.Display)
	s.Equal(DefaultMeasure, u.Measure)
	s.Equal(DefaultFamily, u.Family)
}
