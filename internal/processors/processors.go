package processors

type Type uint8

type Processors struct {
	m map[Type]*list
}

func New() Processors {
	return Processors{m: make(map[Type]*list)}
}

func (ps *Processors) getOrCreate(pType Type) *list {
	if ps.m == nil {
		ps.m = make(map[Type]*list)
	}
	if ps.m[pType] == nil {
		ps.m[pType] = &list{}
	}
	return ps.m[pType]
}

func (ps *Processors) Using(pType Type) *list {
	return ps.getOrCreate(pType)
}

func (ps *Processors) Copy() Processors {
	mCopy := make(map[Type]*list)
	for k, v := range ps.m {
		l := v.copy()
		mCopy[k] = &l
	}
	return Processors{
		m: mCopy,
	}
}
