package comment

import "errors"

type EntityCheckOption func(*Entity) error

func IsIdValid() EntityCheckOption {
	return func(c *Entity) error {
		if c.Id == 0 {
			return errors.New("comment id can't be 0")
		}

		return nil
	}
}

func IsVideoIdValid() EntityCheckOption {
	return func(c *Entity) error {
		if c.VideoId == 0 {
			return errors.New("comment video id can't be 0")
		}

		return nil
	}
}

func IsUserIdValid() EntityCheckOption {
	return func(c *Entity) error {
		if c.UserId == 0 {
			return errors.New("comment user id can't be 0")
		}

		return nil
	}
}

func IsBelongToUser(userId int64) EntityCheckOption {
	return func(c *Entity) error {
		if c.UserId != userId {
			return errors.New("comment is not belong to user")
		}

		return nil
	}
}
