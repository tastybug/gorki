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
			append([]string{filepath.Join(templatesDir, "about.html")},
				getPartialsPaths(templatesDir)...),
			"about",
			siteDir),
		assemblePage(
			append([]string{filepath.Join(templatesDir, "index.html")},
				getPartialsPaths(templatesDir)...),
			"index",
			siteDir),
	}
}

func getPartialsPaths(templatesDir string) []string {
	return []string{
		filepath.Join(templatesDir, "footer.html"),
		filepath.Join(templatesDir, "navigation.html"),
		filepath.Join(templatesDir, "head.html"),
	}
}

func assemblePage(paths []string, canonicalName string, siteDir string) WritableContent {
	tmpl := template.Must(template.ParseFiles(paths...))
	var buffer bytes.Buffer
	articles := CreateOrderListOfPreviewItems(siteDir)
	tmpl.Execute(&buffer, articles) // just in case the template wants to render article list
	return WritableContent{HtmlContent: buffer.String(), FileName: canonicalName + ".html", assets: collectContentSpecificAssets(siteDir, canonicalName)}
}

func collectContentSpecificAssets(siteDir, canonicalName string) []Asset {
	assetFolder := filepath.Join(siteDir, `templates`, canonicalName)

	if !util.Exists(assetFolder) {
		return nil
	}

	allFiles, err := ioutil.ReadDir(assetFolder)
	util.PanicOnError(err)

	var resultMap []Asset
	for _, fileInfo := range allFiles {
		resultMap = append(resultMap, Asset{Filename: fileInfo.Name(), Context: canonicalName, CopyFromPath: filepath.Join(assetFolder, fileInfo.Name())})
	}
	return resultMap
}

func CollectGlobalAssets(siteDir string) []Asset {
	assetFolder := filepath.Join(siteDir, `templates`, `global-assets`)

	allFiles, err := ioutil.ReadDir(assetFolder)
	util.PanicOnError(err)

	var resultMap []Asset
	for _, fileInfo := range allFiles {
		resultMap = append(resultMap, Asset{Filename: fileInfo.Name(), CopyFromPath: filepath.Join(assetFolder, fileInfo.Name())})
	}
	return resultMap
}
