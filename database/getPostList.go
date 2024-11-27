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
	testPost3 := models.Posts{
		Date:  "2023-05-07",
		ID:    3,
		Title: "This is the third title of the post",
	}
	testPost4 := models.Posts{
		Date:  "2023-05-08",
		ID:    4,
		Title: "This is the fourth title of the post",
	}
	testlist := models.List{
		testPost1,
		testPost2,
		testPost3,
		testPost4,
	}
	return testlist, nil
}
