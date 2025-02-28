package domain

import "sort"

type SmallPacker struct {
	Sizes        []PackageSize
	SmallestSize PackageSize
}

// Pack create pack stacks for given quantity with small packages
func (p *SmallPacker) Pack(quantity int) PackStack {
	var sizes []int
	for _, size := range p.Sizes {
		sizes = append(sizes, int(size))
	}

	// sort DESC
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	var packs []Pack
	remaining := quantity
	for _, size := range sizes {
		if remaining == 0 {
			break
		}
		count := remaining / size
		if count > 0 {
			packs = append(packs, Pack{Quantity: count, Size: PackageSize(size)})
			remaining -= count * size
		}

	}

	if remaining > 0 {
		// if there is a remaining quantity, add the smallest package
		packs = append(packs, Pack{Quantity: 1, Size: p.SmallestSize})
	}

	return packs
}

type BigPacker struct {
	Sizes        []PackageSize
	SmallestSize PackageSize
}

// Pack method packs the given quantity with the given package sizes but with maximum package size first
func (p *BigPacker) Pack(quantity int) PackStack {
	var sizes []int
	for _, size := range p.Sizes {
		sizes = append(sizes, int(size))
	}

	// sort ASC
	sort.Sort(sort.IntSlice(sizes))

	var packs []Pack
	remaining := quantity
	for _, size := range sizes {
		if remaining == 0 {
			break
		}
		count := remaining / size
		if count > 0 {
			packs = append(packs, Pack{Quantity: count, Size: PackageSize(size)})

			remaining -= count * size
		}
	}

	if remaining > 0 {
		// if there is a remaining quantity, add the smallest package
		packs = append(packs, Pack{Quantity: 1, Size: p.SmallestSize})
	}

	return packs
}
