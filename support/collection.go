package support

type Collection struct {
	Collection []float64
}

func (collection *Collection) Item() []float64 {
	return collection.Collection
}

func (collection *Collection) Each(callback func(collection *Collection) Collection) Collection {
	return callback(collection)
}

func (collection *Collection) SetItem(items []float64) *Collection {
	collection.Collection = items
	return collection
}

func (collection *Collection) Append(item float64) *Collection {
	collection.Collection = append(collection.Collection, item)
	return collection
}
