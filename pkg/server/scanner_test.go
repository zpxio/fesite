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
	"github.com/stretchr/testify/suite"
	"path/filepath"
	"strings"
	"testing"
)

type ScannerSuite struct {
	suite.Suite
}

func TestScannerSuite(t *testing.T) {
	suite.Run(t, new(ScannerSuite))
}

func (s *ScannerSuite) TestSimpleScan() {
	testDir := filepath.Join("..", "..", "testdata", "site01")

	files, err := ScanForExtension(testDir, "css")

	s.Require().NoError(err, "Failed to scan directory: %s", testDir)
	s.Len(files, 2)

	for _, f := range files {
		s.True(strings.HasSuffix(f.Name(), ".css"))
	}
}

func (s *ScannerSuite) TestEmptyScan() {
	testDir := filepath.Join("..", "..", "testdata", "site01")

	files, err := ScanForExtension(testDir, "nope")

	s.Require().NoError(err, "Failed to scan directory: %s", testDir)
	s.Len(files, 0)
}

func (s *ScannerSuite) TestBadScan() {
	testDir := filepath.Join("..", "..", "testdata", "missingSite")

	files, err := ScanForExtension(testDir, "yml")

	s.Error(err, "Directory somehow scanned: %s", testDir)
	s.Len(files, 0)
}
