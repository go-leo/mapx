package mapx

type Hash[K comparable] interface {
	Sum(k K) int
}

type Hash32Func[K comparable] func(k K) int

func (f Hash32Func[K]) Sum(k K) int {
	return f(k)
}
