package types

// Batch represents a batch of data items for a specific instrument
type Batch[T any] struct {
	InstrumentId InstrumentId
	Data         []T
}

// BatchManager manages batching of items for a specific instrument
type BatchManager[T any] struct {
	instrumentId InstrumentId
	data         []T
	sender       chan<- Batch[T]
	batchSize    int
}

// NewBatchManager creates a new BatchManager
func NewBatchManager[T any](instrumentId InstrumentId, sender chan<- Batch[T], batchSize int) BatchManager[T] {
	return BatchManager[T]{
		instrumentId: instrumentId,
		data:         make([]T, 0, batchSize),
		sender:       sender,
		batchSize:    batchSize,
	}
}

// Add adds items to the batch, flushing when the batch size is reached
func (self *BatchManager[T]) Add(item T) {
	self.data = append(self.data, item)
	if len(self.data) >= self.batchSize {
		self.flush()
	}

}

// Flush sends any remaining data in the batch, even if it's not full
func (self *BatchManager[T]) Flush() {
	if len(self.data) > 0 {
		self.flush()
	}
}

// flush sends the current batch and resets the buffer
func (self *BatchManager[T]) flush() {
	batch := Batch[T]{
		InstrumentId: self.instrumentId,
		Data:         make([]T, len(self.data)),
	}
	copy(batch.Data, self.data)

	self.sender <- batch
	self.data = self.data[:0]
}

// Because you often need to transform the data before you add it into a batch
// This uses a given map function to transform the data for you.
// type BatchManagerWithMap[T0,T any] struct {
// 	batch	BatchManager[T]
// 	mapFn	func(T0) T
// }
// func NewBatchManager[T0,T any](instrumentId InstrumentId, sender chan<- Batch[T], batchSize int, mapFn func(T0) T) BatchManagerWithMap[T0,T] {
// 	return BatchManagerWithMap[T0,T]{
// 		batch:NewBatchManager(instrumentId,sender,batchSize),
// 		mapFn: mapFn,
// 	}
// }

// func (self BatchManagerWithMap[T0, T]) MapAndAdd(items []T0) {
// 	for _,t0 := range items {
// 		t := self.mapFn(t0)
// 		self.batch.Add(t)
// 	}
// }
