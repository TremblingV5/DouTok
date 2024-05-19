package comment_count

import "errors"

type EntityCheckOption func(*Entity) error

func IsIdValid() EntityCheckOption {
	return func(c *Entity) error {
		if c.Id == 0 {
			return errors.New("video id of comment count can't be 0")
		}

		return nil
	}
}
