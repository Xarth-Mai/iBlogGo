package database

import "github.com/Xarth-Mai/iBlogGo/models"

// GetTestPostList returns a test post list
func GetTestPostList() (models.List, error) {
	testPost1 := models.Posts{
		Date:  "2023-05-05",
		ID:    1,
		Title: "This is the title of the post",
	}
	testPost2 := models.Posts{
		Date:  "2023-05-06",
		ID:    2,
		Title: "This is the second title of the post",
	}
	testlist := models.List{
		testPost1,
		testPost2,
	}
	return testlist, nil
}
