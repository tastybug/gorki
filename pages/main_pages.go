package pages

import "path/filepath"

func CollectMains(templatesFolderPath string) []ContentPage {
	return []ContentPage{
		{
			BucketName:    "",
			Title:         "",
			Description:   "",
			PublishedDate: "",
			templatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `index`),
				`index`,
				`index.html`,
				``,
				`index.html`},
		},
		{
			BucketName:    "",
			Title:         "",
			Description:   "",
			PublishedDate: "",
			templatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `about`),
				`about`,
				`about.html`,
				`about`,
				`about.html`},
		},
		{
			BucketName:    "",
			Title:         "",
			Description:   "",
			PublishedDate: "",
			templatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `privacy-imprint`),
				`privacy-imprint`,
				`privacy-imprint.html`,
				`privacy-imprint`,
				`privacy-imprint.html`},
		},
	}
}
