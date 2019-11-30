package research

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidCategory indicates an inalid lookup.
	ErrInvalidCategory = errors.New("invalid category")

	// ErrInvalidMailPackage indicates an inalid lookup.
	ErrInvalidMailPackage = errors.New("invalid mail package")
)

// Marketplace has many categories. It is defined by a country code and URL
type Marketplace struct {
	Categories   []Category    `yaml:"categories"`
	Code         string        `yaml:"code"`
	MailPackages []MailPackage `yaml:"mail_packages"`
	Region       string        `yaml:"region"`
	URL          string        `yaml:"url"`
}

// FindCategoryByName in a marketplace, otherwise ErrInvalidCategory
func (m *Marketplace) FindCategoryByName(name string) (*Category, error) {
	for _, category := range m.Categories {
		if category.Name == name {
			return &category, nil
		}
	}

	return nil, ErrInvalidCategory
}

// ProductURL for a marketplace
func (m *Marketplace) ProductURL(asin string) string {
	// &psc=1 ensures that the size dropdown is selected
	return fmt.Sprintf("%s/dp/%s?psc=1", m.URL, asin)
}

// FindMailPackageByName finds a mail package based in a given name, otherwise ErrInvalidMailPackage
func (m *Marketplace) FindMailPackageByName(name string) (*MailPackage, error) {
	return m.findMailPackageByPredicate(func(m MailPackage) bool { return m.Name == name })
}

// FindMailPackageByID finds a mail package based in a given id, otherwise ErrInvalidMailPackage
func (m *Marketplace) FindMailPackageByID(id int) (*MailPackage, error) {
	return m.findMailPackageByPredicate(func(m MailPackage) bool { return m.ID == id })
}

func (m *Marketplace) findMailPackageByPredicate(predicate func(MailPackage) bool) (*MailPackage, error) {
	for _, mailPackage := range m.MailPackages {
		if predicate(mailPackage) {
			return &mailPackage, nil
		}
	}

	return nil, ErrInvalidMailPackage
}
