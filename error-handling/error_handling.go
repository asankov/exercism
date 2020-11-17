package erratum

func Use(o ResourceOpener, input string) (err error) {
	var r Resource
	defer func() {
		rec := recover()
		if recErr, ok := rec.(error); ok {
			if defrobErr, ok := rec.(FrobError); ok {
				if r != nil {
					r.Defrob(defrobErr.defrobTag)
				}
			}
			err = recErr
		}
		
		if r != nil {
			r.Close()
		}
	}()

	r, err = open(o)
	if err != nil {
		return err
	}
	r.Frob(input)

	return nil
}

func open(o ResourceOpener) (Resource, error) {
	var (
		r   Resource
		err error
	)
	r, err = o()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return open(o)
		}
		return nil, err
	}
	return r, nil
}
