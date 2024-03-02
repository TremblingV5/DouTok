package ParamsValidator

import "github.com/TremblingV5/DouTok/pkg/errno"

type Rule struct {
	Result  bool
	Message *errno.ErrNo
	Ops     []func() (bool, *errno.ErrNo)
	Index   int
}

func New(message *errno.ErrNo) *Rule {
	return &Rule{
		Result:  true,
		Message: message,
		Ops:     []func() (bool, *errno.ErrNo){},
		Index:   0,
	}
}

func (r *Rule) Set(f func() (bool, *errno.ErrNo)) *Rule {
	r.Ops = append(r.Ops, f)
	return r
}

func (r *Rule) SetMore(f ...func() (bool, *errno.ErrNo)) *Rule {
	r.Ops = append(r.Ops, f...)
	return r
}

func (r *Rule) Next() bool {
	if r.Index < len(r.Ops) {
		result, errNo := r.Ops[r.Index]()
		r.Result = r.Result && result
		r.Index++
		if r.Result {
			return r.Next()
		} else {
			r.Message = errNo
			return false
		}
	} else {
		return r.Result
	}
}

func (r *Rule) Validate() (bool, *errno.ErrNo) {
	return r.Next(), r.Message
}
