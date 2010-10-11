
//                                Duck Typing
//                                (real world)

type Vector []interface{}

// add some useful methods
func (a Vector) First() interface{} {
	return a[0]
}

// ... Last, Push, Pop, etc.

// implement Sortable interface
func (a Vector) Len() int { return len(a) }

func (a Vector) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Vector) Less(i, j int) bool {
	return a[i].(Comparable).LessThan(a[j])
}

type Comparable interface {
	LessThan(y interface{}) bool
}
