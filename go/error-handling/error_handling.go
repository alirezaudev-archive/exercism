package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	var r Resource
	var transientError TransientError

	for {
		res, e := opener()
		if e != nil {
			if errors.As(e, &transientError) {
				continue
			}
			return e
		}
		r = res
		break
	}

	defer r.Close()

	defer func() {
		if rec := recover(); rec != nil {
			if fe, ok := rec.(FrobError); ok {
				r.Defrob(fe.defrobTag)
				err = fe
			} else {
				err = rec.(error)
			}
		}
	}()

	r.Frob(input)
	return nil
}
