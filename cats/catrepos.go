package cats

type CatRepo interface {
	Add(cat Cat) (Cat, error)
	Get(id int) (Cat, bool, error)
	All() ([]Cat, error)
	Delete(id int) (bool, error)
	Filter(breed, owner string) ([]Cat, error)
}
