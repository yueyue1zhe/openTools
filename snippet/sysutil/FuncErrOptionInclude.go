package optutil

type FuncErrAddOption func() error

func FuncErrOptionInclude(opts ...FuncErrAddOption) []FuncErrAddOption {
	var opt []FuncErrAddOption
	return append(opt, opts...)
}
