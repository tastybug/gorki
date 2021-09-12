package gorkiconcurrent

import (
	"fmt"
)

type Settings struct {
	SiteRoot      string
	TargetRoot    string
	TemplatesRoot string
	ArticlesRoot  string
}

func Gorkify(settings Settings) {

	discoveredChan := make(chan discovery)
	loadedChan := make(chan bundle)
	renderedChan := make(chan renderedBundle)
	writtenChan := make(chan renderedBundle)

	// discover -> load -> render -> write
	go discover(settings.ArticlesRoot, settings.TemplatesRoot, discoveredChan)
	go load(settings.ArticlesRoot, settings.TemplatesRoot, discoveredChan, loadedChan)
	go render(settings.TemplatesRoot, loadedChan, renderedChan)
	go write(settings.TargetRoot, renderedChan, writtenChan)

	writtenContent := []renderedBundle{}
	for written := range writtenChan {
		writtenContent = append(writtenContent, written)
	}
	fmt.Println("The following articles were written:")
	for _, written := range writtenContent {
		if written.bundle.kind == ARTICLE_BUNDLE {
			fmt.Printf("%s", written)
		}
	}
}
