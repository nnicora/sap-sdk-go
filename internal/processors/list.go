package processors

// A item represents an entry in the list which
// is being run.
type item struct {
	index     int
	processor Processor
	arg       interface{}
}

// A list manages zero or more processors in a list.
type list struct {
	list []Processor

	// Called after each request processor in the list is called. If set
	// and the func returns true the list will continue to iterate
	// over the request processors. If false is returned the list
	// will stop iterating.
	//
	// Should be used if extra logic to be performed between each processor
	// in the list. This can be used to terminate a list's iteration
	// based on a condition such as error like, processorListStopOnError.
	// Or for logging like processorListLogItem.
	afterEachFn func(i item) bool
}

// copy creates a copy of the processor list.
func (l *list) copy() list {
	n := list{
		afterEachFn: l.afterEachFn,
	}
	if len(l.list) == 0 {
		return n
	}

	n.list = append(make([]Processor, 0, len(l.list)), l.list...)
	return n
}

// Clear clears the processor list.
func (l *list) Clear() *list {
	l.list = l.list[0:0]
	return l
}

// Len returns the number of processors in the list.
func (l *list) Len() int {
	return len(l.list)
}

// PushBackHandler pushes processor f to the back of the processor list.
func (l *list) PushBackHandler(f func(arg interface{})) *list {
	l.PushBack(&DefaultProcessor{"-", f})
	return l
}

// PushBack pushes named processor f to the back of the processor list.
func (l *list) PushBack(n Processor) *list {
	if cap(l.list) == 0 {
		l.list = make([]Processor, 0, 5)
	}
	l.list = append(l.list, n)
	return l
}

// PushFrontHandler pushes processor f to the front of the processor list.
func (l *list) PushFrontHandler(f func(arg interface{})) *list {
	l.PushFront(&DefaultProcessor{"-", f})
	return l
}

// PushFront pushes named processor f to the front of the processor list.
func (l *list) PushFront(n Processor) *list {
	if cap(l.list) == len(l.list) {
		// Allocating new list required
		l.list = append([]Processor{n}, l.list...)
	} else {
		// Enough room to prepend into list.
		l.list = append(l.list, &DefaultProcessor{})
		copy(l.list[1:], l.list)
		l.list[0] = n
	}
	return l
}

// Remove removes a DefaultProcessor n
func (l *list) Remove(n Processor) *list {
	l.RemoveByName(n.Label())
	return l
}

// RemoveByName removes a DefaultProcessor by name.
func (l *list) RemoveByName(name string) *list {
	for i := 0; i < len(l.list); i++ {
		m := l.list[i]
		if m.Label() == name {
			// Shift array preventing creating new arrays
			copy(l.list[i:], l.list[i+1:])
			l.list[len(l.list)-1] = &DefaultProcessor{}
			l.list = l.list[:len(l.list)-1]

			// decrement list so next check to length is correct
			i--
		}
	}
	return l
}

// SwapNamed will swap out any existing processors with the same name as the
// passed in DefaultProcessor returning true if processors were swapped. False is
// returned otherwise.
func (l *list) SwapNamed(n Processor) (swapped bool) {
	for i := 0; i < len(l.list); i++ {
		if l.list[i].Label() == n.Label() {
			l.list[i] = n
			swapped = true
		}
	}

	return swapped
}

// Swap will swap out all processors matching the name passed in. The matched
// processors will be swapped in. True is returned if the processors were swapped.
func (l *list) Swap(name string, replace Processor) bool {
	var swapped bool

	for i := 0; i < len(l.list); i++ {
		if l.list[i].Label() == name {
			l.list[i] = replace
			swapped = true
		}
	}

	return swapped
}

// SetBackNamed will replace the named processor if it exists in the processor list.
// If the processor does not exist the processor will be added to the end of the list.
func (l *list) SetBackNamed(n Processor) *list {
	if !l.SwapNamed(n) {
		l.PushBack(n)
	}
	return l
}

// SetFrontNamed will replace the named processor if it exists in the processor list.
// If the processor does not exist the processor will be added to the beginning of
// the list.
func (l *list) SetFrontNamed(n Processor) *list {
	if !l.SwapNamed(n) {
		l.PushFront(n)
	}
	return l
}

func (l *list) StopOnError() {
	l.afterEachFn = func(i item) bool {
		if errorChecker, ok := i.arg.(ErrorChecker); ok {
			return errorChecker.HasError()
		}
		return false
	}
}

func (l *list) Exec(arg interface{}) {
	for i, h := range l.list {
		h.Execute(arg)
		it := item{
			index: i, processor: h, arg: arg,
		}
		if l.afterEachFn != nil && l.afterEachFn(it) {
			return
		}
	}
}
