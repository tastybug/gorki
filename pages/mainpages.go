package pages

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
)

func CollectMainPages(siteDir string) []WritableContent {
	templatesDir := filepath.Join(siteDir, `templates`)
	return []WritableContent{
		assemblePage(
			[]string{
				filepath.Join(templatesDir, "about.html"),
				filepath.Join(templatesDir, "footer.html"),
				filepath.Join(templatesDir, "navigation.html"),
				filepath.Join(templatesDir, "head.html"),
			},
			"about.html",
			siteDir),
		assemblePage(
			[]string{
				filepath.Join(templatesDir, "index.html"),
				filepath.Join(templatesDir, "footer.html"),
				filepath.Join(templatesDir, "navigation.html"),
				filepath.Join(templatesDir, "head.html"),
			},
			"index.html",
			siteDir),
	}
}

func assemblePage(paths []string, fileName string, workFolder string) WritableContent {
	tmpl := template.Must(template.ParseFiles(paths...))
	var buffer bytes.Buffer
	articles := CreateOrderListOfPreviewItems(workFolder)
	tmpl.Execute(&buffer, articles)
	return WritableContent{HtmlContent: buffer.String(), Path: fileName, assets: collectAssets(workFolder)}
}

func collectAssets(siteDir string) []Asset {
	assetFolder := filepath.Join(siteDir, `templates`, `assets`)

	allFiles, err := ioutil.ReadDir(assetFolder)
	util.PanicOnError(err)

	var resultMap []Asset
	for _, fileInfo := range allFiles {
		resultMap = append(resultMap, Asset{Filename: fileInfo.Name(), CopyFromPath: filepath.Join(assetFolder, fileInfo.Name())})
	}
	return resultMap
}
