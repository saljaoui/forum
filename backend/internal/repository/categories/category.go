package category

func AddCategory(post_id int, category string) error {
	err:=postCategory(post_id, category)
	if err!=nil{
		return err
	}
	return nil
}
