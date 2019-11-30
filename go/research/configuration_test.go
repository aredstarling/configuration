// +build spec

package research

import (
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configuration", func() {
	Context("Marketplaces", func() {
		Specify("Has All", func() {
			codes := os.Getenv("MARKETPLACES")
			Expect(codes).ShouldNot(BeEmpty(), "MARKETPLACES env var should be provided for test")

			for _, code := range strings.Split(codes, ",") {
				Expect(FindMarketplaceByCode(code)).ShouldNot(BeNil(), fmt.Sprintf("%s marketplace should exist", code))
			}
		})

		Specify("Valid", func() {
			for _, marketplace := range Marketplaces {
				Expect(len(strings.TrimSpace(marketplace.Code))).ShouldNot(BeZero(), "Marketplace should have Code")
				Expect(len(strings.TrimSpace(marketplace.URL))).ShouldNot(BeZero(), "Marketplace should have URL")
				Expect(len(marketplace.Categories)).ShouldNot(BeZero(), "Marketplace should have Categories")
			}
		})
	})

	Context("Categories", func() {
		Specify("Valid", func() {
			for _, marketplace := range Marketplaces {
				Expect(len(marketplace.Categories)).ShouldNot(BeZero(), "Categories should be present")

				for _, category := range marketplace.Categories {
					Expect(len(strings.TrimSpace(category.Name))).ShouldNot(BeZero(), "Category should have Name")
					Expect(len(strings.TrimSpace(category.Type))).ShouldNot(BeZero(), "Category should have Type")
					Expect(len(category.URLs)).ShouldNot(BeZero())
				}
			}
		})

		Specify("Finds specific category", func() {
			marketplaceCategory := os.Getenv("FIND_CATEGORY")

			Expect(marketplaceCategory).ShouldNot(BeEmpty(), "FIND_CATEGORY env var should be provided for test")

			parts := strings.Split(marketplaceCategory, "|")
			code, name := parts[0], parts[1]
			marketplace, _ := FindMarketplaceByCode(code)
			category, _ := marketplace.FindCategoryByName(name)

			Expect(category).ShouldNot(BeNil())
		})
	})

	Context("MailPackage", func() {
		Specify("Valid", func() {
			for _, marketplace := range Marketplaces {
				Expect(len(marketplace.MailPackages)).ShouldNot(BeZero(), "MailPackages should be present")

				for _, mailPackage := range marketplace.MailPackages {
					Expect(mailPackage.ID).ShouldNot(BeZero(), "MailPackage should have ID")
					Expect(len(strings.TrimSpace(mailPackage.Name))).ShouldNot(BeZero(), "MailPackage should have Name")
				}
			}
		})

		Specify("Finds ID 1", func() {
			marketplace, _ := FindMarketplaceByCode("US")
			mailPackage, _ := marketplace.FindMailPackageByID(1)

			Expect(mailPackage).ShouldNot(BeNil())
		})

		Specify("Finds Standard (Small)", func() {
			marketplace, _ := FindMarketplaceByCode("US")
			mailPackage, _ := marketplace.FindMailPackageByName("Standard (Small)")

			Expect(mailPackage).ShouldNot(BeNil())
		})
	})
})
