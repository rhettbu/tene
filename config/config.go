/*
 *
 * Copyright 2024 gotofuenv authors.
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

package config

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

const (
	LatestAllowedKey = "latest-allowed"
	LatestKey        = "latest"
	MinRequiredKey   = "min-required"

	VersionFileName = ".opentofu-version"
)

const (
	defaultRemoteUrl = "https://api.github.com/repos/opentofu/opentofu/releases"

	envPrefix = "GOTOFUENV_"

	autoInstallEnvName = envPrefix + "AUTO_INSTALL"
	remoteUrlEnvName   = envPrefix + "REMOTE"
	rootPathEnvName    = envPrefix + "ROOT"
	tokenEnvName       = envPrefix + "GITHUB_TOKEN"
	verboseEnvName     = envPrefix + "VERBOSE"
	versionEnvName     = envPrefix + "TOFU_VERSION"
)

type Config struct {
	NoInstall    bool
	RemoteUrl    string
	RootPath     string
	Token        string
	UserHomeFile string
	Verbose      bool
	Version      string
	WorkingDir   bool
}

func InitConfigFromEnv() (Config, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	autoInstall := true
	autoInstallStr := os.Getenv(autoInstallEnvName)
	if autoInstallStr != "" {
		var err error
		autoInstall, err = strconv.ParseBool(autoInstallStr)
		if err != nil {
			return Config{}, err
		}
	}

	remoteUrl := os.Getenv(remoteUrlEnvName)
	if remoteUrl == "" {
		remoteUrl = defaultRemoteUrl
	}

	rootPath := os.Getenv(rootPathEnvName)
	if rootPath == "" {
		rootPath = path.Join(userHome, ".gotofuenv")
	}

	verbose := false
	verboseStr := os.Getenv(verboseEnvName)
	if verboseStr != "" {
		verbose, err = strconv.ParseBool(verboseStr)
		if err != nil {
			return Config{}, err
		}
	}

	return Config{
		NoInstall:    !autoInstall,
		RemoteUrl:    remoteUrl,
		RootPath:     rootPath,
		Token:        os.Getenv(tokenEnvName),
		UserHomeFile: path.Join(userHome, VersionFileName),
		Verbose:      verbose,
		Version:      os.Getenv(versionEnvName),
	}, nil
}

// (made lazy method : not always useful and allows flag override)
func (c *Config) ResolveVersion(defaultVersion string) string {
	if c.Version != "" {
		return c.Version
	}

	data, err := os.ReadFile(VersionFileName)
	if err == nil {
		return string(data)
	}

	data, err = os.ReadFile(c.UserHomeFile)
	if err == nil {
		return string(data)
	}

	data, err = os.ReadFile(c.RootVersionFilePath())
	if err == nil {
		return string(data)
	}
	return defaultVersion
}

// (made lazy method : not always useful and allows flag override)
func (c *Config) RootVersionFilePath() string {
	return path.Join(c.RootPath, VersionFileName)
}

// try to ensure the directory exists with a MkdirAll call.
// (made lazy method : not always useful and allows flag override)
func (c *Config) InstallPath() string {
	dir := path.Join(c.RootPath, "OpenTofu")
	if err := os.MkdirAll(dir, 0755); err != nil && c.Verbose {
		fmt.Println("Can not create installation directory :", err)
	}
	return dir
}
