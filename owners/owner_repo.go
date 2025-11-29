package owners

type OwnerRepo interface {
	Add(o Owner) (Owner, error)
	Get(id int) (Owner, bool, error)
	All() ([]Owner, error)
	Delete(id int) (bool, error)
}
