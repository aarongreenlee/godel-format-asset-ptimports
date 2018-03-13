// Copyright 2016 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/palantir/godel-format-plugin/assetapi"
	"github.com/spf13/cobra"

	"github.com/palantir/godel-format-asset-ptimports/generated_src"
)

const assetName = "ptimports"

func main() {
	rootCmd := assetapi.RootCommand(
		assetName,
		func(cmd *cobra.Command, args []string) error {
			return assetapi.RunAmalgomatedFormatCommand(assetName, args, assetapi.ListFlagVal, []string{"--refactor"}, cmd.OutOrStdout(), cmd.OutOrStderr())
		},
	)
	assetapi.AmalgomatedMain(assetName, rootCmd, amalgomatedformatter.Instance())
}
