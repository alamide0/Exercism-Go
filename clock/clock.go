package clock
import (
	"fmt"
)

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	return deal(Clock{hour, minute})
}

func (c Clock) Add(minute int) Clock {
	c.m += minute
	return deal(c)
}

func (c Clock) Subtract(minute int) Clock {
	c.m -= minute
	return deal(c)
}

func deal(c Clock) Clock {

	sumMin := c.h * 60 + c.m

	if sumMin < 0 {
		c.m = (sumMin % 60 + 60) % 60// deal c.m = 60 
		if c.m == 0 {
			c.h = ((sumMin / 60 % 24) + 24) % 24// deal c.h = 24
		}else {
			c.h = (((sumMin / 60 - 1) % 24) + 24) % 24// deal c.h = 24
		}
	} else {
		c.m = sumMin % 60
		c.h = (sumMin / 60) % 24
	}

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
