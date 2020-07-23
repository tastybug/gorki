package pages

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

func CollectMainPages(workDir string) []WritableContent {
	templatesDir := filepath.Join(workDir, `templates`)
	return []WritableContent{
		assemblePage(
			[]string{
				filepath.Join(templatesDir, "about.html"),
				filepath.Join(templatesDir, "footer.html"),
				filepath.Join(templatesDir, "navigation.html"),
				filepath.Join(templatesDir, "head.html"),
			},
			"about.html",
			workDir),
		assemblePage(
			[]string{
				filepath.Join(templatesDir, "index.html"),
				filepath.Join(templatesDir, "footer.html"),
				filepath.Join(templatesDir, "navigation.html"),
				filepath.Join(templatesDir, "head.html"),
			},
			"index.html",
			workDir),
	}
}

func assemblePage(paths []string, fileName string, workFolder string) WritableContent {
	tmpl := template.Must(template.ParseFiles(paths...))
	var buffer bytes.Buffer
	articles := CreateOrderListOfPreviewItems(workFolder)
	tmpl.Execute(&buffer, articles)
	return WritableContent{HtmlContent: buffer.String(), Path: fileName, assets: collectAssets(workFolder)}
}

func collectAssets(workDir string) map[string]Asset {
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
