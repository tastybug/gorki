package postable

func CollectPostables(workDir string) map[string]Postable {

	postsDir := workDir + "/posts"
	var postableMap = make(map[string]Postable)

	for _, mdFile := range listMarkdownFiles(postsDir) {
		postableMap[mdFile.Name()] = CreatePostable(readFileContent(postsDir, mdFile.Name()))
	}
	return postableMap
}
