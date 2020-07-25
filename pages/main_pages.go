package pages

import "path/filepath"

func CollectMains() []ContentPage {
	templatesFolderPath := GetTemplatesRootDirectory()
	return []ContentPage{
		{
			BucketName:    ``,
			Title:         ``,
			Description:   ``,
			PublishedDate: ``,
			TemplatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `index`),
				`index`,
				`index.html`,
				``,
				`index.html`},
		},
		{
			BucketName:    ``,
			Title:         ``,
			Description:   ``,
			PublishedDate: ``,
			TemplatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `about`),
				`about`,
				`about.html`,
				`about`,
				`about.html`},
		},
		{
			BucketName:    ``,
			Title:         ``,
			Description:   ``,
			PublishedDate: ``,
			TemplatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `privacy-imprint`),
				`privacy-imprint`,
				`privacy-imprint.html`,
				`privacy-imprint`,
				`privacy-imprint.html`},
		},
		{
			BucketName:    ``,
			Title:         ``,
			Description:   ``,
			PublishedDate: ``,
			TemplatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `rss`),
				`rss`,
				`feed.tpl`,
				``,
				`feed.xml`},
		},
	}
}
