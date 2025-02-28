package order

import (
	"github.com/mehgokalp/re-partners-challenge/internal/packaging/domain"
)

func MapPacks(stack domain.PackStack) []Pack {
	var packs []Pack
	for _, pack := range stack {
		packs = append(packs, Pack{
			Quantity: pack.Quantity,
			Size:     int(pack.Size),
		})
	}

	return packs
}
