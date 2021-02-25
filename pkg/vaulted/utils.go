// Copyright 2018 SumUp Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vaulted

import (
	"path/filepath"
	"regexp"
	"strings"
)

var replaceWindowsDriveRegex = regexp.MustCompile(`(?i)[a-z]:\\`)

func Contains(array []string, needle string) bool {
	for _, v := range array {
		if v == needle {
			return true
		}
	}

	return false
}

func SanitizeFilename(filename string) string {
	sanitizedName := strings.ReplaceAll(filename, ".", "_")
	sanitizedName = replaceWindowsDriveRegex.ReplaceAllString(sanitizedName, "")
	sanitizedName = strings.ReplaceAll(
		sanitizedName,
		string(filepath.Separator),
		"_",
	)
	// NOTE: We need to replace the unix file separator, too because on windows the "/" will not be replaced
	sanitizedName = strings.ReplaceAll(
		sanitizedName,
		"/",
		"_",
	)

	return sanitizedName
}
