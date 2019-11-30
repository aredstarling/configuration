package research

import "sort"

// MailPackage structure
type MailPackage struct {
	ID             int       `yaml:"id"`
	Name           string    `yaml:"name"`
	MaxDimensions  []float32 `yaml:"max_dimensions"`
	MaxWeight      float32   `yaml:"max_weight"`
	MaxLengthGirth float32   `yaml:"max_length_girth"`
	MaxPrice       float32   `yaml:"max_price"`
}

// Fits checks if the given dimensions fits into the mail package
func (pack *MailPackage) Fits(height, length, width *float32, weight, price *float32) bool {
	dimensions := pack.descSortedDimensions(height, length, width)
	fitsDimensions := pack.fitsDimensions(dimensions)
	fitsGirth := pack.fitsGirth(dimensions)
	fitsWeight := pack.fitsWeight(weight)

	fitsPrice := pack.fitsPrice(price)

	return fitsDimensions && fitsWeight && fitsGirth && fitsPrice
}

func (pack *MailPackage) fitsDimensions(dimensions []float32) bool {
	// Goes from longest to shortest side, checking if it fits
	for i, side := range pack.MaxDimensions {
		if dimensions[i] > side {
			return false
		}
	}
	return true
}

func (pack *MailPackage) fitsWeight(weight *float32) bool {
	var w float32
	if weight != nil {
		w = *weight
	}

	return pack.MaxWeight == 0 || w <= pack.MaxWeight
}

func (pack *MailPackage) fitsGirth(dimensions []float32) bool {
	girth := (dimensions[1] + dimensions[2]) * 2
	return pack.MaxLengthGirth == 0 || girth+dimensions[0] <= pack.MaxLengthGirth
}

func (pack *MailPackage) fitsPrice(price *float32) bool {
	return pack.MaxPrice == 0 || price != nil && *price <= pack.MaxPrice
}

func (pack *MailPackage) descSortedDimensions(height, length, width *float32) (dimensions []float32) {
	dimensions = []float32{*height, *length, *width}
	sort.Slice(dimensions, func(i, j int) bool { return dimensions[i] > dimensions[j] })
	return
}
