package comment

import "time"

type EntityOption func(*Entity)

func WithId(id int64) EntityOption {
	return func(c *Entity) {
		c.Id = id
	}
}

func WithVideoId(videoId int64) EntityOption {
	return func(c *Entity) {
		c.VideoId = videoId
	}
}

func WithUserId(userId int64) EntityOption {
	return func(c *Entity) {
		c.UserId = userId
	}
}

func WithConversationId(conversationId int64) EntityOption {
	return func(c *Entity) {
		c.ConversationId = conversationId
	}
}

func WithLastId(lastId int64) EntityOption {
	return func(c *Entity) {
		c.LastId = lastId
	}
}

func WithToUserId(toUserId int64) EntityOption {
	return func(c *Entity) {
		c.ToUserId = toUserId
	}
}

func WithContent(content string) EntityOption {
	return func(c *Entity) {
		c.Content = content
	}
}

func WithTimestamp(timestamp string) EntityOption {
	return func(c *Entity) {
		c.Timestamp = timestamp
	}
}

func WithStatus(status bool) EntityOption {
	return func(c *Entity) {
		c.Status = status
	}
}

func WithCreatedAt(createdAt time.Time) EntityOption {
	return func(c *Entity) {
		c.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt time.Time) EntityOption {
	return func(c *Entity) {
		c.UpdatedAt = updatedAt
	}
}
