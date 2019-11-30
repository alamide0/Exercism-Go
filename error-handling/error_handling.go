package erratum
import (
	"fmt"
)

func Use(f func() (Resource, error), input string) (err error) {
	r, err := f()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(f, input)
		}
		return
	}

	defer func(){
		if p := recover(); p != nil {
			if v, ok := p.(FrobError); ok {
				r.Defrob(v.defrobTag)
			} 
			err = fmt.Errorf("%v", p)
		}
		r.Close()
	}()

	r.Frob(input)
	return
}