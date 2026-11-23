package main

type Warehouse struct{ ID int64 }

type Cache interface {
	Put(id int64, value any)
	Get(id int64) (any, bool)
}

type Storage struct {
	cache Cache
}

func (s *Storage) Set(warehouse *Warehouse) {
	s.cache.Put(warehouse.ID, warehouse)
}

func (s *Storage) Get(id int64) *Warehouse {
	item, ok := s.cache.Get(id)
	if !ok {
		return nil
	}

	warehouse, ok := item.(*Warehouse)
	if !ok {
		return nil
	}
	return warehouse
}
