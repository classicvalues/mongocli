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

package opsmanager

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type MetricsProcessOpts struct {
	cli.GlobalOpts
	cli.MetricsOpts
	hostID string
	store  store.HostMeasurementLister
}

func (opts *MetricsProcessOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *MetricsProcessOpts) Run() error {
	listOpts := opts.NewProcessMetricsListOptions()
	result, err := opts.store.HostMeasurements(opts.ConfigProjectID(), opts.hostID, listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli om|cm metric(s) process(es) hostId [--granularity granularity] [--period period] [--start start] [--end end] [--type type][--projectId projectId]
func MetricsProcessBuilder() *cobra.Command {
	opts := &MetricsProcessOpts{}
	cmd := &cobra.Command{
		Use:     "process [hostId]",
		Short:   description.ProcessMeasurements,
		Aliases: []string{"processes"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.hostID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.Granularity, flag.Granularity, "", usage.Granularity)
	cmd.Flags().StringVar(&opts.Period, flag.Period, "", usage.Period)
	cmd.Flags().StringVar(&opts.Start, flag.Start, "", usage.MeasurementStart)
	cmd.Flags().StringVar(&opts.End, flag.End, "", usage.MeasurementEnd)
	cmd.Flags().StringSliceVar(&opts.MeasurementType, flag.Type, nil, usage.MeasurementType)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}