package category

type Category struct {
	Id int
}

func AddCategory(post_id int, category string) {
	postCategory(post_id, category)
}
