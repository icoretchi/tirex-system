package locality_manager

type LocalityCode uint
type LocalityStatisticalCode uint
type LocalityName string
type LocalityStatus uint8

const (
	Raion LocalityStatus = iota + 1
	Oras
	Sector
	Municipiu
	Localitate
	Comuna
	Sat
)

func (t LocalityStatus) ToString() string {
	switch t {
	case Raion:
		return "Raion"
	case Oras:
		return "Oras"
	case Sector:
		return "Sector"
	case Municipiu:
		return "Municipiu"
	case Localitate:
		return "Localitate"
	case Comuna:
		return "Comuna"
	case Sat:
		return "Sat"
	}
	return ""
}

// Locality Data
type Locality struct {
	Code            LocalityCode
	StatisticalCode LocalityStatisticalCode
	Name            LocalityName
	Status          LocalityStatus
}

//NewLocality create a new locality_manager
func NewLocality(code LocalityCode, statisticalCode LocalityStatisticalCode, name LocalityName, status LocalityStatus) (*Locality, error) {
	l := &Locality{
		Code:            code,
		StatisticalCode: statisticalCode,
		Name:            name,
		Status:          status,
	}
	err := l.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return l, nil
}

//Validate validate data
func (l *Locality) Validate() error {
	if l.Name == "" || l.Code == 0 || l.StatisticalCode == 0 || l.Status == 0 {
		return ErrInvalidEntity
	}

	return nil
}
