package domain

import (
	"testing"

	"github.com/rotisserie/eris"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Pack(t *testing.T) {
	packers := []Packer{
		&SmallPacker{Sizes: DefaultPackageSizes, SmallestSize: Small},
		&BigPacker{Sizes: []PackageSize{XXXLarge}, SmallestSize: XXXLarge},
		&BigPacker{Sizes: []PackageSize{XLarge}, SmallestSize: XLarge},
		&BigPacker{Sizes: []PackageSize{Large}, SmallestSize: Large},
		&BigPacker{Sizes: []PackageSize{Medium}, SmallestSize: Medium},
		&BigPacker{Sizes: []PackageSize{Small}, SmallestSize: Small},
	}

	tests := []struct {
		name      string
		orderSize int
		packers   []Packer
		wantErr   error
		wantPacks []Pack
	}{
		{
			name:      "orderSize less than or equal to zero",
			orderSize: 0,
			wantErr:   eris.New("orderSize must be greater than zero"),
		},
		{
			name:      "no packer found",
			orderSize: 10,
			packers:   []Packer{},
			wantErr:   eris.New("no packer found"),
		},
		{
			name:      "test order size: 1",
			orderSize: 1,
			packers:   packers,
			wantErr:   nil,
			wantPacks: []Pack{
				{
					Size:     Small,
					Quantity: 1,
				},
			},
		},
		{
			name:      "test order size: 250",
			orderSize: 250,
			packers:   packers,
			wantErr:   nil,
			wantPacks: []Pack{
				{
					Size:     Small,
					Quantity: 1,
				},
			},
		},
		{
			name:      "test order size: 251",
			orderSize: 251,
			packers:   packers,
			wantErr:   nil,
			wantPacks: []Pack{
				{
					Size:     Medium,
					Quantity: 1,
				},
			},
		},
		{
			name:      "test order size: 501",
			orderSize: 501,
			packers:   packers,
			wantErr:   nil,
			wantPacks: []Pack{
				{
					Size:     Medium,
					Quantity: 1,
				},
				{
					Size:     Small,
					Quantity: 1,
				},
			},
		},
		{
			name:      "test order size: 12001",
			orderSize: 12001,
			packers:   packers,
			wantErr:   nil,
			wantPacks: []Pack{
				{
					Size:     XXXLarge,
					Quantity: 2,
				},
				{
					Size:     XLarge,
					Quantity: 1,
				},
				{
					Size:     Small,
					Quantity: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Packers:      tt.packers,
				PackageSizes: DefaultPackageSizes,
			}
			gotPacks, err := h.Pack(tt.orderSize)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantPacks, gotPacks)
			}
		})
	}
}
