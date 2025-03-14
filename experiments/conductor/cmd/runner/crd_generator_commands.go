// Copyright 2025 Google LLC
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

package runner

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// NOT USED
/*
const GENERATOR_SCRIPT_TEMPLATE = `#!/bin/bash
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

go run . generate-types \
    --service <PROTO_PACKAGE> \
    --api-version <CRD_GROUP>/<CRD_VERSION> \
    --resource <CRD_KIND>:<PROTO_RESOURCE>

go run . generate-mapper \
    --service <PROTO_PACKAGE> \
    --api-version <CRD_GROUP>/<CRD_VERSION>

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w pkg/controller/direct/<SERVICE>/
`

// NOT USED
const UPDATE_GENERATE_SCRIPT_PROMPT = `
Please update the apis/<SERVICE>/v1alpha1/generate.sh script for the <SERVICE> API to generate the CRD for the <CRD_KIND> resource.

The generate.sh script is located at apis/<SERVICE>/v1alpha1/generate.sh.

Add the parameter <TICK> --resource <CRD_KIND>:<PROTO_RESOURCE><TICK> to the <TICK>go run . generate-types --api-version <CRD_GROUP>/<CRD_VERSION>  <TICK> command.

At the end of the script, ensure the following lines are present:

cd ${REPO_ROOT}
dev/tasks/generate-crds

go run -mod=readonly golang.org/x/tools/cmd/goimports@latest -w pkg/controller/direct/<SERVICE>/
`

// NOT USED
func generateCRDFromScripts(opts *RunnerOptions, branch Branch) {
	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := opts.branchRepoDir

	var out strings.Builder
	checkoutBranch(branch, workDir, &out)

	// Create the apis/<service>/<version> directory
	serviceDir := filepath.Join(workDir, "apis", branch.Group, "v1alpha1")
	if err := os.MkdirAll(serviceDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Create or update generate.sh
	scriptPath := filepath.Join(serviceDir, "generate.sh")
	// Check if generate.sh already exists.
	if _, err := os.Stat(scriptPath); errors.Is(err, os.ErrNotExist) {
		// File doesn't exist, use template approach
		log.Printf("Creating new generate.sh at %s", scriptPath)

		// Replace template markers with actual values and write to file
		writeTemplateToFile(branch, scriptPath, GENERATOR_SCRIPT_TEMPLATE)
	} else {
		// File exists, use codebot to update it
		log.Printf("Updating existing generate.sh at %s", scriptPath)

		// Delete then write the prompt file.
		promptPath := filepath.Join(workDir, "mockgcp", "crdgen_prompt.txt")
		writeTemplateToFile(branch, promptPath, UPDATE_GENERATE_SCRIPT_PROMPT)

		cfg := CommandConfig{
			Name:    "CODEBOT GENERATE",
			Cmd:     "codebot",
			Args:    []string{"--ui-type=prompt", "--prompt=mockgcp/crdgen_prompt.txt"},
			WorkDir: workDir,
		}
		if err := executeCommand(cfg, &out); err != nil {
			log.Fatal(err)
		}
	}

	// Stage the changed files
	scriptRelativePath := filepath.Join("apis", branch.Group, "v1alpha1", "generate.sh")
	gitAdd(workDir, &out, scriptRelativePath)

	// Commit the changes
	gitCommit(workDir, &out, fmt.Sprintf("add/update crd generation script for %s", branch.Group))

	// Run the generator script
	cfg := CommandConfig{
		Name:    "Generator script",
		Cmd:     scriptPath,
		WorkDir: workDir,
	}
	if err := executeCommand(cfg, &out); err != nil {
		log.Fatal(err)
	}

	// Stage the changed files
	gitAdd(workDir, &out,
		fmt.Sprintf("apis/%s/v1alpha1/", branch.Group),
		fmt.Sprintf("pkg/controller/direct/%s/", branch.Group),
		"config/crds/resources/")

	// Commit the changes
	gitCommit(workDir, &out, fmt.Sprintf("autogenerated types and CRDs using %s", scriptRelativePath))
}
*/

func generateTypesAndMapper(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := opts.branchRepoDir

	if branch.Kind == "" || branch.Proto == "" || branch.Group == "" {
		if branch.Kind == "" {
			log.Printf("SKIPPING %s, missing Kind", branch.Name)
		}
		if branch.Proto == "" {
			log.Printf("SKIPPING %s, missing Proto", branch.Name)
		}
		if branch.Group == "" {
			log.Printf("SKIPPING %s, missing Group", branch.Name)
		}
		return
	}

	checkoutBranch(ctx, branch, workDir)

	// Change to controllerbuilder directory
	controllerBuilderDir := filepath.Join(workDir, "dev", "tools", "controllerbuilder")
	hasChange := false

	// Generate types
	apisDir := filepath.Join(opts.branchRepoDir, "apis", branch.Group, "v1alpha1", string(filepath.Separator))
	if _, err := os.Stat(apisDir); errors.Is(err, os.ErrNotExist) || opts.force {
		cfg := CommandConfig{
			Name: "Generate types",
			Cmd:  "go",
			Args: []string{
				"run", ".",
				"generate-types",
				"--service", branch.Package,
				"--api-version", fmt.Sprintf("%s.cnrm.cloud.google.com/v1alpha1", branch.Group),
				"--resource", fmt.Sprintf("%s:%s", branch.Kind, branch.Proto),
			},
			WorkDir:    controllerBuilderDir,
			MaxRetries: 1,
		}
		_, _, err := executeCommand(opts, cfg)
		if err != nil {
			log.Fatal(err)
		}
		gitAdd(ctx, workDir, apisDir)
		hasChange = true
	} else {
		log.Printf("SKIPPING generating apis, %s already exists", apisDir)
	}

	// Generate mapper
	mapperDir := filepath.Join(opts.branchRepoDir, "pkg", "controller", "direct", branch.Group, string(filepath.Separator))
	if _, err := os.Stat(mapperDir); errors.Is(err, os.ErrNotExist) || opts.force {
		cfg := CommandConfig{
			Name: "Generate mapper",
			Cmd:  "go",
			Args: []string{
				"run", ".",
				"generate-mapper",
				"--service", branch.Package,
				"--api-version", fmt.Sprintf("%s.cnrm.cloud.google.com/v1alpha1", branch.Group),
			},
			WorkDir:    controllerBuilderDir,
			MaxRetries: 2,
		}
		_, _, err := executeCommand(opts, cfg)
		if err != nil {
			log.Fatal(err)
		}

		// Run goimports
		cfg = CommandConfig{
			Name: "Format generated code",
			Cmd:  "go",
			Args: []string{
				"run", "-mod=readonly",
				"golang.org/x/tools/cmd/goimports@latest",
				"-w",
				mapperDir,
			},
			WorkDir:    workDir,
			MaxRetries: 2,
		}
		_, _, err = executeCommand(opts, cfg)
		if err != nil {
			log.Fatal(err)
		}
		gitAdd(ctx, workDir, mapperDir)
		hasChange = true
	} else {
		log.Printf("SKIPPING generating mappers, %s already exists", mapperDir)
	}

	// Commit the changes
	if hasChange {
		gitCommit(ctx, workDir, fmt.Sprintf("Generated types and mapper for %s", branch.Kind))
	} else {
		log.Printf("SKIPPING git commit, no new changes for %s", branch.Name)
	}
}

func generateCRD(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := opts.branchRepoDir

	checkoutBranch(ctx, branch, workDir)

	// Generate CRDs
	cfg := CommandConfig{
		Name:       "Generate CRDs",
		Cmd:        filepath.Join(workDir, "dev", "tasks", "generate-crds"),
		WorkDir:    workDir,
		MaxRetries: 1,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Stage the changed files
	gitAdd(ctx, workDir, "config/crds/resources/")

	// Commit the changes
	gitCommit(ctx, workDir, fmt.Sprintf("Generated CRD for %s", branch.Kind))
}

func generateSpecStatus(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := opts.branchRepoDir

	checkoutBranch(ctx, branch, workDir)

	// Run controllerbuilder to generate spec and status
	log.Printf("Generating spec and status for %s", branch.Name)
	stdinInput := fmt.Sprintf("// +kcc:proto=%s.%s\n", branch.ProtoSvc, branch.Proto)

	cfg := CommandConfig{
		Name:         "Spec/Status generation",
		Cmd:          "controllerbuilder",
		Args:         []string{"prompt", "--src-dir", workDir, "--proto-dir", filepath.Join(workDir, ".build", "third_party", "googleapis")},
		WorkDir:      workDir,
		Stdin:        strings.NewReader(stdinInput),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	_, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Stage the changed files
	gitAdd(ctx, workDir, fmt.Sprintf("apis/%s/v1alpha1/%s_types.go", branch.Group, strings.ToLower(branch.Resource)))

	// Commit the changes
	gitCommit(ctx, workDir, fmt.Sprintf("%s: Update types from generated", branch.Kind))
}

func generateFuzzer(opts *RunnerOptions, branch Branch) {
	ctx := context.TODO()

	close := setLoggingWriter(opts, branch)
	defer close()
	workDir := opts.branchRepoDir

	checkoutBranch(ctx, branch, workDir)

	// Generate fuzzer file
	fuzzerDir := filepath.Join(workDir, "pkg", "controller", "direct", branch.Group)
	if err := os.MkdirAll(fuzzerDir, 0755); err != nil {
		log.Fatal(err)
	}

	fuzzerPath := filepath.Join(fuzzerDir, fmt.Sprintf("%s_fuzzer.go", strings.ToLower(branch.Resource)))
	stdinInput := fmt.Sprintf(`// +tool:fuzz-gen
// proto.message: %s
`, branch.ProtoMsg)

	cfg := CommandConfig{
		Name:         "Fuzzer generation",
		Cmd:          "controllerbuilder",
		Args:         []string{"prompt", "--src-dir", workDir, "--proto-dir", filepath.Join(workDir, ".build", "third_party", "googleapis")},
		WorkDir:      workDir,
		Stdin:        strings.NewReader(stdinInput),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	output, _, err := executeCommand(opts, cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(fuzzerPath, []byte(output), 0644); err != nil {
		log.Fatal(err)
	}

	// Update register.go to import the new package
	registerPath := filepath.Join(workDir, "pkg", "controller", "direct", "register", "register.go")
	importLine := fmt.Sprintf(`_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/%s"`, branch.Group)
	stdinInput = fmt.Sprintf("Add an unnamed (_) go import for %s to the imports in %s", importLine, registerPath)

	cfg = CommandConfig{
		Name:         "Import addition",
		Cmd:          "codebot",
		Args:         []string{"--prompt=/dev/stdin"},
		WorkDir:      workDir,
		Stdin:        strings.NewReader(stdinInput),
		RetryBackoff: GenerativeCommandRetryBackoff,
	}
	_, _, err = executeCommand(opts, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Stage the changed files
	gitAdd(ctx, workDir,
		fmt.Sprintf("pkg/controller/direct/%s/%s_fuzzer.go", branch.Group, strings.ToLower(branch.Resource)),
		"pkg/controller/direct/register/register.go")

	// Commit the changes
	gitCommit(ctx, workDir, fmt.Sprintf("%s: Create fuzz test", branch.Kind))
}
