package post

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

const POSTS_FILENAME_PATTERN string = `.*\.md`

func GetSitePosts(workDir string) map[string]BlogPost {

	return listPosts(workDir + "/posts")
}

func listFiles(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func listPosts(postsDir string) map[string]BlogPost {
	var fileNamePattern = regexp.MustCompile(POSTS_FILENAME_PATTERN)
	files := listFiles(postsDir)

	var postMap map[string]BlogPost = make(map[string]BlogPost)

	for _, file := range files {
		if !fileNamePattern.MatchString(file.Name()) {
			continue
		}
		content := readContent(postsDir, file.Name())
		postMap[file.Name()] = CreateBlogPost(content)
	}
	return postMap
}

func readContent(dir string, fileName string) string {
	file, err := os.Open(dir + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += (scanner.Text() + "\n")
	}

	return content
}
