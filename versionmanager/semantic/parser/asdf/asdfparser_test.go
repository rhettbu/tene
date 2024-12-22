/*
 *
 * Copyright 2024 tofuutils authors.
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
 *
 */

package asdfparser

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/tofuutils/tenv/v3/pkg/loghelper"
)

//go:embed testdata/.tool-versions
var toolFileData []byte

func TestParseVersionFromToolFileReader(t *testing.T) {
	t.Parallel()

	t.Run("BasicLine", func(t *testing.T) {
		version, err := parseVersionFromToolFileReader("", bytes.NewReader(toolFileData), "nodejs", loghelper.InertDisplayer)
		if err != nil {
			t.Fatal("Unexpected parse error : ", err)
		}

		if version != "10.15.0" {
			t.Fatal("Unexpected version : ", version)
		}
	})

	t.Run("LineWithComment", func(t *testing.T) {
		version, err := parseVersionFromToolFileReader("", bytes.NewReader(toolFileData), "ruby", loghelper.InertDisplayer)
		if err != nil {
			t.Fatal("Unexpected parse error : ", err)
		}

		if version != "2.5.3" {
			t.Fatal("Unexpected version : ", version)
		}
	})
}
