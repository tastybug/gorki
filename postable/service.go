package postable

import "path/filepath"

func CollectPostables(workDir string) map[string]Postable {

	postsDir := filepath.Join(workDir, "posts")
	var postableMap = make(map[string]Postable)

	for _, mdFile := range listMarkdownFiles(postsDir) {
		postableMap[mdFile.Name()] = CreatePostableFromRawString(readFileContent(postsDir, mdFile.Name()))
	}
	return postableMap
}
