package post

const POSTS_FILENAME_PATTERN string = ".*\\.md"

func GetSitePosts(workDir string) []BlogPost {
	var first BlogPost
	first.Title = "The Bible"
	first.Description = "God has spoken"
	first.Filename = "bible.md"

	var second BlogPost
	second.Title = "The Hobbit"
	second.Description = "JRR Tolkien about little people"
	second.Filename = "the_hobbit.md"

	return []BlogPost{first, second}
}

type BlogPost struct {
	Title       string
	Description string
	Filename    string
}
