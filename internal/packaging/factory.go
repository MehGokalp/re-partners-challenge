package packaging

func NewDefaultHandler() Handler {
	return Handler{
		// You can modify package sizes on your own
		// As mentioned in the challenge, it could be fetched from query string
		// But it's better to keep it hard coded for being more understandable
		PackageSizes: DefaultPackageSizes,
		Packers: []Packer{
			&SmallPacker{Sizes: DefaultPackageSizes, SmallestSize: Small},
			// Big packer is going to create packs with big sizes but fewer packages
			// BigPacker can be used for different behaviour. We made a little hack
			// here to keep more functionality without any more complexity. That's why
			// function args are looking a little bit weird.
			&BigPacker{Sizes: []PackageSize{XXXLarge}, SmallestSize: XXXLarge},
			&BigPacker{Sizes: []PackageSize{XLarge}, SmallestSize: XLarge},
			&BigPacker{Sizes: []PackageSize{Large}, SmallestSize: Large},
			&BigPacker{Sizes: []PackageSize{Medium}, SmallestSize: Medium},
			&BigPacker{Sizes: []PackageSize{Small}, SmallestSize: Small},
		},
	}
}
