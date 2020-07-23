package proc

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

type Asset struct {
	SourcePath string
	TargetPath string
}

func CollectOtherContent(templatesDir string) []WritableContent {
	return []WritableContent{
		assemblePage(
			[]string{
				filepath.Join(templatesDir, `about`, "_about.html"),
				filepath.Join(templatesDir, "footer.html"),
				filepath.Join(templatesDir, "navigation.html"),
				filepath.Join(templatesDir, "head.html"),
			},
			"about.html"),
		assemblePage(
			[]string{
				filepath.Join(templatesDir, `index`, "_index.html"),
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

func CollectAssets(workDir, targetDir string) map[string]Asset {
	assetFolder := filepath.Join(workDir, `templates`, `assets`)

	allFiles, err := ioutil.ReadDir(assetFolder)
	util.PanicOnError(err)

	var resultMap = make(map[string]Asset)
	for _, fileInfo := range allFiles {
		resultMap[fileInfo.Name()] = Asset{
			SourcePath: filepath.Join(assetFolder, fileInfo.Name()),
			TargetPath: filepath.Join(targetDir, fileInfo.Name())}
	}
	return resultMap
}
