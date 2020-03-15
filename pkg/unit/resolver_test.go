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
	"strings"
	"testing"
)

type ResolverTestSuite struct {
	suite.Suite
}

func TestResolverTestSuite(t *testing.T) {
	suite.Run(t, new(ResolverTestSuite))
}

func (s *ResolverTestSuite) TestResolution_SpotCheck() {
	s.Equal(Tablespoon, Resolve("Tbsp"))
	s.Equal(Cup, Resolve("c"))
	s.Equal(Teaspoon, Resolve("tsp"))
}

func (s *ResolverTestSuite) TestResolution_CaseInsensitive() {
	tu := &Unit{
		Display: "Testor",
		Family:  Universal,
		Measure: Number,
	}

	tuAlias := "TEST"
	registerLookup(tu, tuAlias)

	// Exact Match
	s.Equal(tu, Resolve(tuAlias), "Failed to resolve exact match")
	s.Equal(tu, Resolve(strings.ToLower(tuAlias)), "Failed to resolve lowercase alias")
	s.Equal(tu, Resolve(strings.ToTitle(tuAlias)), "Failed to resolve title case")
}
