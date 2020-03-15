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

var Dictionary = make(map[string]Unit)

type Family string

const (
	Metric    Family = "metric"
	Imperial  Family = "imperial"
	Universal Family = "universal"
)

const DefaultFamily = Universal

type Measure int

const (
	Weight Measure = iota
	Volume
	Temperature
	Number
)

const DefaultMeasure = Number

type Unit struct {
	Display string
	Family  Family
	Measure Measure
}

func init() {
	reinitialize()
}

func reinitialize() {
	Dictionary = make(map[string]Unit)
	registerImperialVolume()
	registerImperialWeight()
	registerMetricVolume()
	registerMetricWeight()

	registerGeneric()
}

func (u *Unit) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var name string
	err := unmarshal(&name)

	if err != nil {
		return err
	}

	resolvedUnit := Resolve(name)

	u.Display = resolvedUnit.Display
	u.Measure = resolvedUnit.Measure
	u.Family = resolvedUnit.Family

	return nil
}
