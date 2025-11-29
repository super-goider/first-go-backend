package cats

import "strconv"

type InMemoryCatRepo struct {
	cats []Cat
}

func NewInMemory() *InMemoryCatRepo {
	return &InMemoryCatRepo{cats: make([]Cat, 0)}
}

func (mem *InMemoryCatRepo) Add(cat Cat) (Cat, error) {
	cat.ID = len(mem.cats) + 1
	mem.cats = append(mem.cats, cat)
	return cat, nil
}

func (mem *InMemoryCatRepo) Get(id int) (Cat, bool, error) {
	for _, k := range mem.cats {
		if k.ID == id {
			return k, true, nil
		}
	}
	return Cat{}, false, nil
}

func (mem *InMemoryCatRepo) All() ([]Cat, error) {
	catsCopy := make([]Cat, len(mem.cats))
	copy(catsCopy, mem.cats)
	return catsCopy, nil
}

func (mem *InMemoryCatRepo) Delete(id int) (bool, error) {
	for i, c := range mem.cats {
		if c.ID == id {
			mem.cats = append(mem.cats[:i], mem.cats[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (mem *InMemoryCatRepo) Filter(breed, owner string) ([]Cat, error) {
	result := make([]Cat, 0)

	for _, c := range mem.cats {
		if breed != "" && breed != c.Breed {
			continue
		}

		s, _ := strconv.Atoi(owner)

		if owner != "" && c.OwnerId != s {
			continue
		}
		result = append(result, c)
	}
	return result, nil
}
