package common

import "errors"

// rpc service
type Math struct{}

type Args struct {
	A, B int
}

func (math *Math) Add(args Args, ans *int) error {
	*ans = args.A + args.B
	return nil
}

func (math *Math) Substract(args Args, ans *int) error {
	*ans = args.A - args.B
	return nil
}

func (math *Math) Multiply(args Args, ans *int) error {
	*ans = args.A * args.B
	return nil
}

func (math *Math) Divide(args Args, ans *float64) error {
	if args.B == 0 {
		return errors.New("can't deivid by zero")
	}

	*ans = (float64(args.A) / float64(args.B))
	return nil
}
