package postable

func CollectPostables(workDir string) map[string]Postable {

	postsDir := workDir + "/posts"
	files := listPostableFiles(postsDir)

	var postableMap = make(map[string]Postable)
	for _, file := range files {
		content := readFileContent(postsDir, file.Name())
		postableMap[file.Name()] = CreatePostableFromFile(content)
	}
	return postableMap
}
