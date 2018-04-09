package errors

import "fmt"

// AggregateError is error which composes multiple errors into one
type AggregateError struct {
	errs []error
}

//Add adds new error
func (c *AggregateError) Add(e error) {
	if c.errs == nil {
		c.errs = make([]error, 0, 1)
	}
	c.errs = append(c.errs, e)
}

func (c *AggregateError) Error() (err string) {
	err = ""
	for i, e := range c.errs {
		err += fmt.Sprintf("error %d: %s\n", i, e.Error())
	}
	return err
}
