package erratum

// Use uses a resource with an input.
func Use(o ResourceOpener, input string) error {
	var res Resource
	var err error
	// Try to open.
	for res, err = o(); err != nil; res, err = o() {
		if _, ok := err.(TransientError); ok {
			// Retry on TransientError.
			continue
		}
		return err
	}
	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			if f, ok := r.(FrobError); ok {
				res.Defrob(f.defrobTag)
				err = f.inner
			}
		}
	}()
	res.Frob(input)
	return err
}
