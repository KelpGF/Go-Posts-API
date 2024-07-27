package dto

type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	AuthorName  string `json:"author_name"`
	PublishedAt string `json:"published_at"`
}

type ListPostsInput struct {
	AuthorName    string
	PublishedSort string
	Paginate
}

type ListPostsOutput struct {
	Posts []Post
}
