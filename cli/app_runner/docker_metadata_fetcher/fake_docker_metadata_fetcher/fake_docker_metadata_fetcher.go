// This file was generated by counterfeiter
package fake_docker_metadata_fetcher

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/cli/app_runner/docker_metadata_fetcher"
)

type FakeDockerMetadataFetcher struct {
	FetchMetadataStub        func(repoName string, tag string) (*docker_metadata_fetcher.ImageMetadata, error)
	fetchMetadataMutex       sync.RWMutex
	fetchMetadataArgsForCall []struct {
		repoName string
		tag      string
	}
	fetchMetadataReturns struct {
		result1 *docker_metadata_fetcher.ImageMetadata
		result2 error
	}
}

func (fake *FakeDockerMetadataFetcher) FetchMetadata(repoName string, tag string) (*docker_metadata_fetcher.ImageMetadata, error) {
	fake.fetchMetadataMutex.Lock()
	fake.fetchMetadataArgsForCall = append(fake.fetchMetadataArgsForCall, struct {
		repoName string
		tag      string
	}{repoName, tag})
	fake.fetchMetadataMutex.Unlock()
	if fake.FetchMetadataStub != nil {
		return fake.FetchMetadataStub(repoName, tag)
	} else {
		return fake.fetchMetadataReturns.result1, fake.fetchMetadataReturns.result2
	}
}

func (fake *FakeDockerMetadataFetcher) FetchMetadataCallCount() int {
	fake.fetchMetadataMutex.RLock()
	defer fake.fetchMetadataMutex.RUnlock()
	return len(fake.fetchMetadataArgsForCall)
}

func (fake *FakeDockerMetadataFetcher) FetchMetadataArgsForCall(i int) (string, string) {
	fake.fetchMetadataMutex.RLock()
	defer fake.fetchMetadataMutex.RUnlock()
	return fake.fetchMetadataArgsForCall[i].repoName, fake.fetchMetadataArgsForCall[i].tag
}

func (fake *FakeDockerMetadataFetcher) FetchMetadataReturns(result1 *docker_metadata_fetcher.ImageMetadata, result2 error) {
	fake.FetchMetadataStub = nil
	fake.fetchMetadataReturns = struct {
		result1 *docker_metadata_fetcher.ImageMetadata
		result2 error
	}{result1, result2}
}

var _ docker_metadata_fetcher.DockerMetadataFetcher = new(FakeDockerMetadataFetcher)
