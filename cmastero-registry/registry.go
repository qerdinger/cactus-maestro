package cmastero_registry

import (
	"encoding/json"
	"errors"
	"fmt"
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

func NewRegistry(address string, username *string, password *string) *Registry {
	r := Registry{config: &RegistryConfig{address, username, password}, client: &http.Client{}}

	return &r
}

func (r *Registry) GetCatalog() error {
	if r.config == nil || r.client == nil {
		return errors.New("missing configuration/client")
	}

	registryV2 := fmt.Sprintf("%s/v2/", r.config.Address)

	catalogUrl := fmt.Sprintf("%s/_catalog", registryV2)
	req, err := http.NewRequest("GET", catalogUrl, nil)

	// Add authentification if specified
	if r.config.Username != nil || r.config.Password != nil {
		req.SetBasicAuth(*r.config.Username, *r.config.Password)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}

	var catalog CatalogResponse
	json.NewDecoder(resp.Body).Decode(&catalog)

	for _, repo := range catalog.Repositories {
		url := fmt.Sprintf("%s/%s/tags/list", registryV2, repo)

		req, _ := http.NewRequest("GET", url, nil)
		// Add authentification if specified
		if r.config.Username != nil || r.config.Password != nil {
			req.SetBasicAuth(*r.config.Username, *r.config.Password)
		}

		resp, err := r.client.Do(req)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		var tags TagsResponse
		json.NewDecoder(resp.Body).Decode(&tags)
		resp.Body.Close()

		fmt.Printf("---> Repository: %s\n", repo)
		fmt.Printf("     |--> Tags: %v\n\n", tags.Tags)
	}
	return nil
}
