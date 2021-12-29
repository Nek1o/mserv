package main

type MangadexProvider struct {
}

func (p *MangadexProvider) GetChapter(id string) (Chapter, error) {
	panic("not implemented") // TODO: Implement
}

func (p *MangadexProvider) ListChapters(id string) ([]Chapter, error) {
	return []Chapter{
		{
			Name:  "The show begins!",
			Pages: []string{"a", "b", "c"},
		},
	}, nil
}

func (p *MangadexProvider) Search(name string) ([]Title, error) {
	panic("not implemented") // TODO: Implement
}

func (p *MangadexProvider) ToChapterModel() Chapter {
	panic("not implemented") // TODO: Implement
}

func (p *MangadexProvider) ToTitleModel() Title {
	panic("not implemented") // TODO: Implement
}
