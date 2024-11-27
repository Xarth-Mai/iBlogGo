package database

import "github.com/Xarth-Mai/iBlogGo/models"

// GetTestPost returns a test post
func GetTestPost(id string) (models.Post, error) {
	testPost1 := models.Post{
		Content: "This is the content of the post",
		Date:    "2023-05-05",
		ID:      1,
		Title:   "This is the title of the post",
	}
	testPost2 := models.Post{
		Content: "This is the content of the second post",
		Date:    "2023-05-06",
		ID:      2,
		Title:   "This is the second title of the post",
	}
	if id == "1" {
		return testPost1, nil
	} else {
		return testPost2, nil
	}
}
