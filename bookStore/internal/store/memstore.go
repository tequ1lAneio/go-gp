package store

import (
	myStore "bookStore/store"
	"bookStore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*myStore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*myStore.Book
}

func (ms *MemStore) Create(book *myStore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.books[book.Id]; ok {
		return myStore.ErrExist
	}

	nBook := *book
	ms.books[book.Id] = &nBook

	return nil
}

func (ms *MemStore) Update(book *myStore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	oldBook, ok := ms.books[book.Id]
	if !ok {
		return myStore.ErrNotFound
	}

	nBook := *oldBook
	if book.Name != "" {
		nBook.Name = book.Name
	}

	if book.Authors != nil {
		nBook.Authors = book.Authors
	}

	if book.Press != "" {
		nBook.Press = book.Press
	}

	ms.books[book.Id] = &nBook

	return nil
}

func (ms *MemStore) Get(id string) (myStore.Book, error) {
	ms.RLock()
	defer ms.RUnlock()

	t, ok := ms.books[id]
	if ok {
		return *t, nil
	}
	return myStore.Book{}, myStore.ErrNotFound
}

func (ms *MemStore) Delete(id string) error {
	ms.Lock()

	if _, ok := ms.books[id]; !ok {
		return myStore.ErrNotFound
	}

	delete(ms.books, id)
	return nil
}

func (ms *MemStore) GetAll() ([]myStore.Book, error) {
	ms.RLock()
	defer ms.RUnlock()

	allBooks := make([]myStore.Book, 0, len(ms.books))
	for _, book := range ms.books {
		allBooks = append(allBooks, *book)
	}
	return allBooks, nil
}
