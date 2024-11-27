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
	testPost3 := models.Post{
		Content: "This is the content of the third post",
		Date:    "2023-05-07",
		ID:      3,
		Title:   "This is the third title of the post",
	}
	testPost4 := models.Post{
		Content: "This is the content of the fourth post",
		Date:    "2023-05-08",
		ID:      4,
		Title:   "This is the fourth title of the post",
	}
	switch id {
	case "1":
		return testPost1, nil
	case "2":
		return testPost2, nil
	case "3":
		return testPost3, nil
	case "4":
		return testPost4, nil
	default:
		return testPost1, nil
	}
}
