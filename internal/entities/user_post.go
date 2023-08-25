package entities

type UserPost struct {
	ID       int
	UserID   int
	ImageURL string
	VideoURL string
	Headline string
	Likes    Like
	Comments []Comment
}

type ObjectType string

const (
	PostType    ObjectType = "post"
	CommentType ObjectType = "comment"
)

type Like struct {
	ID         int
	UserID     int
	ObjectType ObjectType
	ObjectID   int
}

type Comment struct {
	ID     int
	UserID int
	Text   string
}
