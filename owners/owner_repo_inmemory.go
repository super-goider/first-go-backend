package owners

type InMemoryOwnerRepo struct {
	owners []Owner
}

func NewInMemory() *InMemoryOwnerRepo {
	return &InMemoryOwnerRepo{owners: make([]Owner, 0)}
}

func (mem *InMemoryOwnerRepo) Add(own Owner) (Owner, error) {
	own.ID = len(mem.owners) + 1
	mem.owners = append(mem.owners, own)
	return own, nil
}

func (mem *InMemoryOwnerRepo) Get(id int) (Owner, bool, error) {
	for _, o := range mem.owners {
		if o.ID == id {
			return o, true, nil
		}
	}
	return Owner{}, false, nil
}

func (mem *InMemoryOwnerRepo) All() ([]Owner, error) {
	ownersCopy := make([]Owner, len(mem.owners))
	copy(ownersCopy, mem.owners)
	return ownersCopy, nil
}

func (mem *InMemoryOwnerRepo) Delete(id int) (bool, error) {
	for i, o := range mem.owners {
		if o.ID == id {
			mem.owners = append(mem.owners[:i], mem.owners[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}
