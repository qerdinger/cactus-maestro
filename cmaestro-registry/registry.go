package cmastero_registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Registry struct {
	config *RegistryConfig
	client *http.Client
}

type CatalogResponse struct {
	Repositories []string `json:"repositories"`
}

type TagsResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type Repository struct {
	Name string
	Tags []string
}

func (r *Registry) requestBuilder(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	// Authentification
	if r.config.Username != nil || r.config.Password != nil {
		req.SetBasicAuth(*r.config.Username, *r.config.Password)
	}

	return req, nil
}

func New(address string, username *string, password *string) *Registry {
	r := Registry{config: &RegistryConfig{address, username, password}, client: &http.Client{}}

	return &r
}

func (r *Registry) GetCatalog() (*[]Repository, error) {
	if r.config == nil || r.client == nil {
		return nil, errors.New("missing configuration/client")
	}

	registryV2 := fmt.Sprintf("%s/v2/", r.config.Address)

	catalogUrl := fmt.Sprintf("%s/_catalog", registryV2)
	req, err := r.requestBuilder("GET", catalogUrl, nil)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	var catalogResp CatalogResponse
	json.NewDecoder(resp.Body).Decode(&catalogResp)

	var catalog []Repository

	for _, repo := range catalogResp.Repositories {
		tags, _ := r.GetTags(repo)
		catalog = append(catalog, Repository{repo, tags})
	}

	fmt.Println(catalog)

	return &catalog, nil
}

func (r *Registry) GetTags(repo string) ([]string, error) {
	if r.config == nil || r.client == nil {
		return nil, errors.New("missing configuration/client")
	}

	registryV2 := fmt.Sprintf("%s/v2/", r.config.Address)

	url := fmt.Sprintf("%s/%s/tags/list", registryV2, repo)

	req, err := r.requestBuilder("GET", url, nil)

	resp, err := r.client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	}

	var tags TagsResponse
	json.NewDecoder(resp.Body).Decode(&tags)
	resp.Body.Close()

	fmt.Printf("---> Repository: %s\n", repo)
	fmt.Printf("     |--> Tags: %v\n\n", tags.Tags)
	return tags.Tags, nil
}

func (r *Registry) GetDigest(repo string, tag string) (string, error) {
	if r.config == nil || r.client == nil {
		return "", errors.New("missing configuration/client")
	}

	registryV2 := fmt.Sprintf("%s/v2/", r.config.Address)
	url := fmt.Sprintf("%s/%s/manifests/%s", registryV2, repo, tag)

	req, err := r.requestBuilder("GET", url, nil)
	resp, err := r.client.Do(req)

	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}

	return resp.Header.Get("Docker-Content-Digest"), nil
}

func (r *Registry) RemoveTag(repo string, tag string) (bool, error) {
	if r.config == nil || r.client == nil {
		return false, errors.New("missing configuration/client")
	}

	digest, err := r.GetDigest(repo, tag)

	if err != nil {
		return false, err
	}

	registryV2 := fmt.Sprintf("%s/v2", r.config.Address)
	fmt.Println("Removing --> Repository:", repo)
	fmt.Println("			    |--> Digest:", digest)
	url := fmt.Sprintf("%s/%s/manifests/%s", registryV2, repo, digest)

	fmt.Println("Deletion Url:", url)

	req, err := r.requestBuilder("DELETE", url, nil)
	resp, err := r.client.Do(req)

	if err != nil {
		fmt.Println("error:", err)
		return false, err
	}

	// assert whether return code is 202/StatusAccepted
	return resp.StatusCode == http.StatusAccepted, nil
}
