// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/10gen/mcli/internal/usage"
	"github.com/10gen/mcli/internal/utils"
	"github.com/spf13/cobra"
)

type iamProjectsListOpts struct {
	*globalOpts
	store store.ProjectLister
}

func (opts *iamProjectsListOpts) init() error {
	if err := opts.loadConfig(); err != nil {
		return err
	}

	s, err := store.New(opts.Config)

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *iamProjectsListOpts) Run() error {
	projects, err := opts.store.GetAllProjects()

	if err != nil {
		return err
	}

	return utils.PrettyJSON(projects)
}

// mcli iam project(s) list [--orgId orgId]
func IAMProjectsListBuilder() *cobra.Command {
	opts := &iamProjectsListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List projects",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.profile, flags.Profile, "p", config.DefaultProfile, usage.Profile)

	return cmd
}