package entity

type Store map[string]Pair

func (s *Store) SnapShot() Store {
	snapShot := make(Store)
	for key, value := range *s {
		snapShot[key] = value
	}
	return snapShot
}
