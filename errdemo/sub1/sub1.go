package sub1

import (
	"./sub2"
	"errors"
)

func Diff(foo int, bar int) error {
	if foo < 0 {
		return errors.New("diff error")
	}

	if err := sub2.Diff(foo, bar); err != nil {
		return err
	}

	return nil
}
