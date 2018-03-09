package decorator

import (
	"sort"

	"github.com/ufoscout/go-up/reader"
)

// PriorityQueueDecoratorReader A Reader that wraps a prioritized list of other Readers.
// Each Reader in the list has a priority that is used to resolve conflicts
// when a key is defined more than once.
// If two or more Readers have the same priority, the one added by last has the highest priority
type PriorityQueueDecoratorReader struct {
	ReadersMap map[int][]reader.Reader
}

func (f *PriorityQueueDecoratorReader) Read() (map[string]reader.Property, error) {

	output := map[string]reader.Property{}

	// To store the keys in slice in sorted order
	var keys []int
	for k := range f.ReadersMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
		value := f.ReadersMap[key]
		for _, readers := range value {

			properties, err := readers.Read()

			if err != nil {
				return nil, err
			}

			for k, v := range properties {
				output[k] = v
			}
		}
	}
	return output, nil

}

// Add adds a new Reader
func (f *PriorityQueueDecoratorReader) Add(newReader reader.Reader, priority int) *PriorityQueueDecoratorReader {
	readers, found := f.ReadersMap[priority]
	if !found {
		readers = []reader.Reader{}
		f.ReadersMap[priority] = readers
	}
	readers = append(readers, newReader)
	f.ReadersMap[priority] = readers
	return f
}

// NewPriorityQueueDecoratorReader creats a new PriorityQueueDecoratorReader instance
func NewPriorityQueueDecoratorReader() *PriorityQueueDecoratorReader {
	return &PriorityQueueDecoratorReader{map[int][]reader.Reader{}}
}
