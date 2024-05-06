package main

type Repository struct {
	Index int
	Posts []Post
}

func NewRepository() Repository {
	return Repository{
		Index: 0,
		Posts: []Post{},
	}
}

func (s *Repository) CreatePost(post Post) (Post, error) {
	s.Index += 1
	post.Id = s.Index
	s.Posts = append(s.Posts, post)
	return post, nil
}

func (s *Repository) ReadPosts() ([]Post, error) {
	return s.Posts, nil
}

// func (s *Repository) UpdatePostById(id int, title string, body string) error {
// 	s.Posts[id].Title = title
// 	s.Posts[id].Body = body
// 	return nil
// }

// func (s *Repository) DeletePostById(id int) error {
// 	s.Posts = append(s.Posts[:id], s.Posts[id+1:]...)
// 	return nil
// }
