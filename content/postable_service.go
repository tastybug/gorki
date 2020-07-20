package content

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func CollectPostables(workDir string) map[string]Postable {

	return collectPostables(workDir + "/posts")
}

func collectPostables(postsDir string) map[string]Postable {
	var fileNamePattern = regexp.MustCompile(`.*\.md`)
	files := listPostableFiles(postsDir)

	var postMap = make(map[string]Postable)

	for _, file := range files {
		if !fileNamePattern.MatchString(file.Name()) {
			continue
		}
		content := readFileContent(postsDir, file.Name())
		postMap[file.Name()] = CreatePostableFromFile(content)
	}
	return postMap
}

func listPostableFiles(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func readFileContent(dir string, fileName string) string {
	file, err := os.Open(dir + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	return content
}
