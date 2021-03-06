package container

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/archive"
)

// BuildResponse is the object that is returned by the docker api in json
type BuildResponse struct {
	Stream string `json:"stream"`
	Error  string `json:"error"`
}

// Build builds a docker image.
func (c *Container) Build(path string) (tag string, err error) {
	excludeFiles := []string{}
	// TODO: remove .mesgignore for a future release
	for _, file := range []string{".mesgignore", ".dockerignore"} {
		excludeFilesBytes, _ := ioutil.ReadFile(filepath.Join(path, file))
		excludeFiles = append(excludeFiles, strings.Fields(string(excludeFilesBytes))...)
	}
	buildContext, err := archive.TarWithOptions(path, &archive.TarOptions{
		Compression:     archive.Gzip,
		ExcludePatterns: excludeFiles,
	})
	if err != nil {
		return "", err
	}
	defer buildContext.Close()
	response, err := c.client.ImageBuild(context.Background(), buildContext, types.ImageBuildOptions{
		Remove:         true,
		ForceRemove:    true,
		SuppressOutput: true,
	})
	if err != nil {
		return "", err
	}
	return parseBuildResponse(response)
}

func parseBuildResponse(response types.ImageBuildResponse) (tag string, err error) {
	lastOutput, err := extractLastOutputFromBuildResponse(response)
	if err != nil {
		return "", err
	}
	var buildResponse BuildResponse

	if err := json.Unmarshal([]byte(lastOutput), &buildResponse); err != nil {
		return "", fmt.Errorf("Could not parse container build response. %s", err)
	}
	if buildResponse.Error != "" {
		return "", fmt.Errorf("Image build failed. %s", buildResponse.Error)
	}
	return strings.TrimSuffix(buildResponse.Stream, "\n"), nil
}

func extractLastOutputFromBuildResponse(response types.ImageBuildResponse) (lastOutput string, err error) {
	defer response.Body.Close()
	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	lastOutput = ""
	rs := strings.Split(string(r), "\n")
	i := len(rs) - 1
	for lastOutput == "" && i >= 0 {
		lastOutput = rs[i]
		i--
	}
	if lastOutput == "" {
		return "", errors.New("Could not parse container build response")
	}
	return lastOutput, nil
}
