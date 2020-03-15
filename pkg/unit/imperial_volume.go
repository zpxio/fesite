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

// Create units
var Teaspoon = &Unit{
	Display: "tsp",
	Family:  Imperial,
	Measure: Volume,
}

var Tablespoon = &Unit{
	Display: "Tbsp",
	Family:  Imperial,
	Measure: Volume,
}

var Cup = &Unit{
	Display: "c",
	Family:  Imperial,
	Measure: Volume,
}

var Ounce = &Unit{
	Display: "oz",
	Family:  Imperial,
	Measure: Volume,
}

// Register Units
func registerImperialVolume() {
	registerLookup(Teaspoon, "tsp", "teaspoon")
	registerLookup(Tablespoon, "tbsp", "tablespoon")
	registerLookup(Cup, "cup", "c")
	registerLookup(Ounce, "oz", "ounce")
}
