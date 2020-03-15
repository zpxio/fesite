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
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"testing"
)

type RecipeSuite struct {
	suite.Suite
}

func TestRecipeSuite(t *testing.T) {
	suite.Run(t, new(RecipeSuite))
}

func (s *RecipeSuite) TestParse_Sample() {
	sampleFilename := "fluffy-pancakes.yml"
	file := filepath.Join("..", "..", "testdata", sampleFilename)

	r, err := ParseFile(file)

	s.Require().Nil(err, "Failed to parse sample file: %s", sampleFilename)
	s.Require().NotNil(r, "Parse returned nil")

	s.NotEmpty(r.Name)
	s.NotEmpty(r.Intro)
	s.Len(r.IngredientGroups, 3, "Incorrect number of ingredient groups")
	s.Len(r.InstructionSets, 1, "Incorrect number of instruction sets")
	s.Len(r.InstructionSets[0].Steps, 7, "Incorrect number of steps in instruction set")
}
