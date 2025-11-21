package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var res Resource
	for {
		res, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case FrobError:
				res.Defrob(v.defrobTag)
				err = v.inner
			case error:
				err = v
			}
		}
		_ = res.Close()
	}()
	res.Frob("hello")
	return err
}
