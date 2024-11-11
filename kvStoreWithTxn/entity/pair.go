package entity

type Pair struct {
	Key     string
	Value   string
	Version int
}

func (p *Pair) Copy() Pair {
	return Pair{Key: p.Key, Value: p.Value, Version: p.Version}
}
