// Copyright 2021 MongoDB Inc
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

package quickstart

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongocli/internal/mongosh"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/validate"
)

func newClusterNameQuestion(clusterName string) *survey.Question {
	return &survey.Question{
		Name: "clusterName",
		Prompt: &survey.Input{
			Message: "Cluster Name [This can't be changed later]",
			Help:    usage.ClusterName,
			Default: clusterName,
		},
		Validate: survey.ComposeValidators(survey.Required, validate.ClusterName),
	}
}

func newClusterProviderQuestion() *survey.Question {
	return &survey.Question{
		Name: "provider",
		Prompt: &survey.Select{
			Message: "Cloud Provider",
			Help:    usage.Provider,
			Options: []string{"AWS", "GCP", "AZURE"},
		},
		Validate: survey.Required,
	}
}

func newAccessListQuestion(publicIP, message string) survey.Prompt {
	return &survey.Input{
		Message: fmt.Sprintf("Access List Entry%s", message),
		Help:    usage.NetworkAccessListIPEntry,
		Default: publicIP,
	}
}

func newDBUsernameQuestion(dbUser string, validation survey.Validator) *survey.Question {
	q := &survey.Question{
		Name: "dbUsername",
		Prompt: &survey.Input{
			Message: "Database User Username",
			Help:    usage.DBUsername,
			Default: dbUser,
		},
		Validate: survey.ComposeValidators(survey.Required, validate.DBUsername, validation),
	}
	return q
}

func newDBUserPasswordQuestion(password, message string) *survey.Question {
	return &survey.Question{
		Name: "DBUserPassword",
		Prompt: &survey.Input{
			Message: fmt.Sprintf("Database User Password%s", message),
			Help:    usage.Password,
			Default: password,
		},
		Validate: survey.ComposeValidators(survey.Required, validate.WeakPassword),
	}
}

func newSampleDataQuestion(clusterName string) survey.Prompt {
	return &survey.Confirm{
		Message: fmt.Sprintf("Do you want to load sample data into %s?", clusterName),
		Help:    "Load sample data to help you test cluster features. See: https://docs.atlas.mongodb.com/sample-data/available-sample-datasets/.",
		Default: true,
	}
}

func newMongoShellQuestionAccessDeployment(clusterName string) survey.Prompt {
	return &survey.Confirm{
		Message: fmt.Sprintf("Do you want to access %s with MongoDB Shell?", clusterName),
		Help:    "MongoDB CLI will use your installed version of MongoDB Shell to access your deployments.",
		Default: true,
	}
}

func newMongoShellPathQuestion() survey.Prompt {
	return &survey.Confirm{
		Message: "Do you want to provide the path to your MongoDB Shell binary?",
		Help:    "MongoDB CLI will store the path in your profile, type ‘mongocli config’ to change it.",
		Default: true,
	}
}

func newIsMongoShellInstalledQuestion() survey.Prompt {
	return &survey.Confirm{
		Message: "Do you have a MongoDB Shell version installed on your machine?",
		Default: true,
	}
}

func newMongoShellPathInput() survey.Prompt {
	return &survey.Input{
		Message: "Default MongoDB Shell Path",
		Help:    "MongoDB CLI will use your installed version of MongoDB Shell to access your deployments.",
		Default: mongosh.Path(),
	}
}

func newMongoShellQuestionOpenBrowser() survey.Prompt {
	return &survey.Confirm{
		Message: "Do you want to download MongoDB Shell [This will open www.mongodb.com on your browser]?",
		Default: true,
	}
}

func newAtlasAccountQuestionOpenBrowser() survey.Prompt {
	return &survey.Confirm{
		Message: "Do you want to create an Atlas account [This will open www.mongodb.com on your browser]?",
		Default: true,
	}
}

func newProfileDocQuestionOpenBrowser() survey.Prompt {
	return &survey.Confirm{
		Message: "Do you want more information to set up your profile [This will open www.mongodb.com on your browser]?",
		Default: true,
	}
}
