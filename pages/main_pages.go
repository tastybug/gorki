package pages

import (
	"gorki/util"
	"path/filepath"
)

func CollectMains() []Page {
	templatesFolderPath := util.GetSettings().TemplatesRoot
	return []Page{
		{
			TemplatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `index`),
				`index`,
				`index.html`,
				``,
				`index.html`},
		},
		{
			TemplatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `about`),
				`about`,
				`about.html`,
				`about`,
				`about.html`},
		},
		{
			TemplatingConf: TemplatingConf{
				``,
				filepath.Join(templatesFolderPath, `privacy-imprint`),
				`privacy-imprint`,
				`privacy-imprint.html`,
				`privacy-imprint`,
				`privacy-imprint.html`},
		},
		{
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
