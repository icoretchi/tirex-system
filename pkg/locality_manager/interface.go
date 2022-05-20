package locality_manager

//Reader interface
type Reader interface {
	Get(code LocalityCode) (*Locality, error)
	Search(query string) ([]*Locality, error)
	List() ([]*Locality, error)
}

//Writer interface
type Writer interface {
	Create(e *Locality) (LocalityCode, error)
	Update(e *Locality) error
	Delete(id LocalityCode) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetLocality(id LocalityCode) (*Locality, error)
	SearchLocalities(query string) ([]*Locality, error)
	ListLocalities() ([]*Locality, error)
	CreateLocality(code LocalityCode, statisticalCode LocalityStatisticalCode, name LocalityName, status LocalityStatus) (LocalityCode, error)
	UpdateLocality(e *Locality) error
	DeleteLocality(code LocalityCode) error
}
