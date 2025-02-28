package domain

type PackageSize int

const (
	Small    PackageSize = 250
	Medium               = 500
	Large                = 1000
	XLarge               = 2000
	XXXLarge             = 5000
)

var DefaultPackageSizes = []PackageSize{Small, Medium, Large, XLarge, XXXLarge}

type Handler struct {
	PackageSizes []PackageSize
	Packers      []Packer
}

type Pack struct {
	Quantity int
	Size     PackageSize
}

type Packer interface {
	Pack(quantity int) PackStack
}

type PackStack []Pack

func (s PackStack) TotalQuantity() int {
	sum := 0
	for _, pack := range s {
		sum += pack.Quantity * int(pack.Size)
	}
	return sum
}

func (s PackStack) QuantityDelta(orderSize int) int {
	return s.TotalQuantity() - orderSize
}

func (s PackStack) PackSize() int {
	sum := 0
	for _, pack := range s {
		sum += pack.Quantity
	}

	return sum
}

// Weight calculates the weight of the pack stack
// The weight is calculated by multiplying the quantity delta with the pack size
// Sample:
// orderSize: 501, 1x500 + 1x250 -> weight: 498
// orderSize: 501, 1x1000 -> weight: 499
// orderSize: 501, 3x250 -> weight: 747
// lowest weight is the best option
func (s PackStack) Weight(orderSize int) int {
	return s.QuantityDelta(orderSize) * s.PackSize()
}
