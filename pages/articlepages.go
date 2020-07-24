package pages

import (
	"bloggo/util"
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func TurnArticlesIntoContentPack(articles Articles, siteDir string) []ContentPack {
	templatesFolder := filepath.Join(siteDir, `templates`)
	postsDir := filepath.Join(siteDir, "posts")

	var contentPacks []ContentPack
	for _, postable := range articles.Articles {
		contentPacks = append(contentPacks, toContentPack(postable, postsDir, templatesFolder))
	}
	return contentPacks
}

func toContentPack(article Article, postsDir string, templatesFolder string) ContentPack {
	var htmlString bytes.Buffer

	contentTemplate := createContentTemplate(article.ContentAsHtml)
	t, _ := template.ParseFiles(
		append([]string{
			filepath.Join(templatesFolder, `blogpost`, `blogpost.html`),
			filepath.Join(contentTemplate.Name()),
		},
			getPartialsPaths(templatesFolder)...)...)
	err := t.Execute(&htmlString, article)
	util.PanicOnError(err)

	defer os.Remove(contentTemplate.Name())

	return ContentPack{
		Folders:     article.BucketName,
		HtmlContent: htmlString.String(),
		FileName:    "article.html",
		assets:      collectAssetsForArticle(postsDir, article),
	}
}

func createContentTemplate(content string) *os.File {
	return util.WriteToTempFile("{{define \"content\"}}" + content + "{{end}}")
}

func collectAssetsForArticle(bucketsDir string, article Article) []Asset {
	assetFolderPath := filepath.Join(bucketsDir, article.BucketName)
	var assets []Asset
	if util.Exists(assetFolderPath) {
		for _, assetFile := range util.ListFilesWithoutSuffix(assetFolderPath, `.md`) {
			log.Println("Article %s has asset %s "+article.BucketName, assetFile)
			assets = append(assets,
				Asset{
					Context:      article.BucketName,
					Filename:     assetFile.Name(),
					CopyFromPath: filepath.Join(assetFolderPath, assetFile.Name())})
		}
	}
	if len(assets) > 0 {
		return assets
	} else {
		return nil
	}
}
