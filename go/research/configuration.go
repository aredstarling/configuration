package research

import (
	"errors"
	"fmt"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type configuration struct {
	Marketplaces []Marketplace `yaml:"marketplaces"`
}

var (
	environmentConfiguration *configuration

	// ErrInvalidMarketplace indicates an inalid lookup.
	ErrInvalidMarketplace = errors.New("invalid marketplace")

	// Marketplaces for the current configuration
	Marketplaces []Marketplace
)

func init() {
	var err error
	environmentConfiguration, err = loadConfiguration()
	if err != nil {
		panic(err)
	}

	Marketplaces = environmentConfiguration.Marketplaces
}

func loadConfiguration() (*configuration, error) {
	data, err := Asset(fmt.Sprintf("../product_research/%s.yml", os.Getenv("APP_ENV")))
	if err != nil {
		return nil, err
	}

	var configuration configuration

	if err := yaml.Unmarshal(data, &configuration); err != nil {
		return nil, err
	}

	return &configuration, nil
}

func (configuration *configuration) findMarketplacByCode(code string) (*Marketplace, error) {
	for _, marketplace := range configuration.Marketplaces {
		if strings.Contains(marketplace.Code, code) {
			return &marketplace, nil
		}
	}

	return nil, ErrInvalidMarketplace
}

// FindMarketplaceByCode from the configuration, otherwise ErrInvalidMarketplace
func FindMarketplaceByCode(code string) (*Marketplace, error) {
	return environmentConfiguration.findMarketplacByCode(code)
}
