package domain

import (
	"github.com/rotisserie/eris"
	"sort"
	"sync"
)

func (h *Handler) Pack(orderSize int) ([]Pack, error) {
	if orderSize <= 0 {
		return nil, eris.New("orderSize must be greater than zero")
	}

	var stacks []PackStack
	wg := sync.WaitGroup{}
	wg.Add(len(h.Packers))

	for _, packer := range h.Packers {
		go func() {
			defer wg.Done()
			stack := packer.Pack(orderSize)

			if stack != nil {
				stacks = append(stacks, stack)
			}
		}()
	}

	wg.Wait()

	if len(stacks) == 0 {
		return nil, eris.New("no packer found")
	}

	return h.choiceBest(stacks, orderSize), nil
}

func (h *Handler) choiceBest(stacks []PackStack, orderSize int) []Pack {
	weights := make(map[int]PackStack)
	for _, stack := range stacks {
		// calculate weight for each stack
		weights[stack.Weight(orderSize)] = stack
	}

	keys := make([]int, 0, len(weights))
	for k := range weights {
		keys = append(keys, k)
	}

	// sort ASC
	sort.Ints(keys)

	// return most efficient stack
	return weights[keys[0]]
}
