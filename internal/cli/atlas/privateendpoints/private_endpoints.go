// Copyright 2020 MongoDB Inc
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

package privateendpoints

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/atlas/privateendpoints/interfaces"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "privateEndpoints"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   privateEndpoints,
	}
	cmd.AddCommand(ListBuilder())
	cmd.AddCommand(DescribeBuilder())
	cmd.AddCommand(CreateBuilder())
	cmd.AddCommand(DeleteBuilder())
	cmd.AddCommand(WatchBuilder())
	cmd.AddCommand(interfaces.Builder())

	return cmd
}
