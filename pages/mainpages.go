package pages

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

func CollectMainPages(templatesDir string) []WritableContent {
	return []WritableContent{
		assemblePage(
			[]string{
				filepath.Join(templatesDir, "_about.html"),
				filepath.Join(templatesDir, "footer.html"),
				filepath.Join(templatesDir, "navigation.html"),
				filepath.Join(templatesDir, "head.html"),
			},
			"about.html"),
		assemblePage(
			[]string{
				filepath.Join(templatesDir, "_index.html"),
				filepath.Join(templatesDir, "footer.html"),
				filepath.Join(templatesDir, "navigation.html"),
				filepath.Join(templatesDir, "head.html"),
			},
			"index.html"),
	}
}

func assemblePage(paths []string, fileName string) WritableContent {
	tmpl := template.Must(template.ParseFiles(paths...))
	var b bytes.Buffer
	tmpl.Execute(&b, nil)
	return WritableContent{HtmlContent: b.String(), Path: fileName}
}

func CollectAssets(workDir string) map[string]Asset {
	assetFolder := filepath.Join(workDir, `templates`, `assets`)

	allFiles, err := ioutil.ReadDir(assetFolder)
	util.PanicOnError(err)

	var resultMap = make(map[string]Asset)
	for _, fileInfo := range allFiles {
		resultMap[fileInfo.Name()] = Asset{
			Filename:     fileInfo.Name(),
			CopyFromPath: filepath.Join(assetFolder, fileInfo.Name())}
	}
	return resultMap
}
