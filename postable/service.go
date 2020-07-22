package postable

import (
	"bloggo/util"
	"path/filepath"
)

func CollectPostables(workDir string) map[string]Postable {

	postsDir := filepath.Join(workDir, "posts")
	var postableMap = make(map[string]Postable)

	for _, mdFile := range util.ListFilesWithSuffix(postsDir, ".md") {
		postableMap[mdFile.Name()] = CreatePostableFromRawString(util.ReadFileContent(postsDir, mdFile.Name()))
	}
	return postableMap
}
