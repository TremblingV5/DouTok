package comment_count

type Option func(*Entity)

func WithId(id int64) Option {
	return func(e *Entity) {
		e.Id = id
	}
}

func WithNumber(number int64) Option {
	return func(e *Entity) {
		e.Number = number
	}
}
