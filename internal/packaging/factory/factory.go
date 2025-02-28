package factory

import (
	"github.com/mehgokalp/re-partners-challenge/internal/packaging/domain"
)

func NewDefaultHandler() domain.Handler {
	return domain.Handler{
		// You can modify package sizes on your own
		// As mentioned in the challenge, it could be fetched from query string
		// But it's better to keep it hard coded for being more understandable
		PackageSizes: domain.DefaultPackageSizes,
		Packers: []domain.Packer{
			&domain.SmallPacker{Sizes: domain.DefaultPackageSizes, SmallestSize: domain.Small},
			// Big packer is going to create packs with big sizes but fewer packages
			// BigPacker can be used for different behaviour. We made a little hack
			// here to keep more functionality without any more complexity. That's why
			// function args are looking a little bit weird.
			&domain.BigPacker{Sizes: []domain.PackageSize{domain.XXXLarge}, SmallestSize: domain.XXXLarge},
			&domain.BigPacker{Sizes: []domain.PackageSize{domain.XLarge}, SmallestSize: domain.XLarge},
			&domain.BigPacker{Sizes: []domain.PackageSize{domain.Large}, SmallestSize: domain.Large},
			&domain.BigPacker{Sizes: []domain.PackageSize{domain.Medium}, SmallestSize: domain.Medium},
			&domain.BigPacker{Sizes: []domain.PackageSize{domain.Small}, SmallestSize: domain.Small},
		},
	}
}
