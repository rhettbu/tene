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

package builder

import (
	"github.com/tofuutils/tenv/config"
	"github.com/tofuutils/tenv/versionmanager"
	terraformretriever "github.com/tofuutils/tenv/versionmanager/retriever/terraform"
	terragruntretriever "github.com/tofuutils/tenv/versionmanager/retriever/terragrunt"
	tofuretriever "github.com/tofuutils/tenv/versionmanager/retriever/tofu"
	"github.com/tofuutils/tenv/versionmanager/semantic"
	flatparser "github.com/tofuutils/tenv/versionmanager/semantic/parser/flat"
	terragruntparser "github.com/tofuutils/tenv/versionmanager/semantic/parser/terragrunt"
	tomlparser "github.com/tofuutils/tenv/versionmanager/semantic/parser/toml"
	"github.com/tofuutils/tenv/versionmanager/semantic/parser/types"
)

func BuildTfManager(conf *config.Config, gruntParser terragruntparser.TerragruntParser) versionmanager.VersionManager {
	tfRetriever := terraformretriever.NewTerraformRetriever(conf)
	versionFiles := []types.VersionFile{
		{Name: ".terraform-version", Parser: flatparser.RetrieveVersion},
		{Name: ".tfswitchrc", Parser: flatparser.RetrieveVersion},
		{Name: terragruntparser.HCLName, Parser: gruntParser.RetrieveTerraformVersionConstraintFromHCL},
		{Name: terragruntparser.JSONName, Parser: gruntParser.RetrieveTerraformVersionConstraintFromJSON},
	}

	return versionmanager.MakeVersionManager(conf, "Terraform", semantic.TfPredicateReaders, tfRetriever, config.TfVersionEnvName, config.TfDefaultVersionEnvName, versionFiles)
}

func BuildTgManager(conf *config.Config, gruntParser terragruntparser.TerragruntParser) versionmanager.VersionManager {
	tgRetriever := terragruntretriever.NewTerragruntRetriever(conf)
	versionFiles := []types.VersionFile{
		{Name: ".terragrunt-version", Parser: flatparser.RetrieveVersion},
		{Name: ".tgswitchrc", Parser: flatparser.RetrieveVersion},
		{Name: ".tgswitch.toml", Parser: tomlparser.RetrieveVersion},
		{Name: terragruntparser.HCLName, Parser: gruntParser.RetrieveTerragruntVersionConstraintFromHCL},
		{Name: terragruntparser.JSONName, Parser: gruntParser.RetrieveTerragruntVersionConstraintFromJSON},
	}

	return versionmanager.MakeVersionManager(conf, "Terragrunt", nil, tgRetriever, config.TgVersionEnvName, config.TgDefaultVersionEnvName, versionFiles)
}

func BuildTofuManager(conf *config.Config, gruntParser terragruntparser.TerragruntParser) versionmanager.VersionManager {
	tofuRetriever := tofuretriever.NewTofuRetriever(conf)
	versionFiles := []types.VersionFile{
		{Name: ".opentofu-version", Parser: flatparser.RetrieveVersion},
		{Name: terragruntparser.HCLName, Parser: gruntParser.RetrieveTerraformVersionConstraintFromHCL},
		{Name: terragruntparser.JSONName, Parser: gruntParser.RetrieveTerraformVersionConstraintFromJSON},
	}

	return versionmanager.MakeVersionManager(conf, "OpenTofu", semantic.TfPredicateReaders, tofuRetriever, config.TofuVersionEnvName, config.TofuDefaultVersionEnvName, versionFiles)
}
