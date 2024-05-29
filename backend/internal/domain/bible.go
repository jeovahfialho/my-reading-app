package domain

// BibleVerse represents a single verse in the Bible.
type BibleVerse struct {
	Livro      string   `json:"livro"`
	Capitulo   int      `json:"capitulo"`
	Versiculos []string `json:"versiculos"`
}

type Bible struct {
	Verses []BibleVerse `json:"verses"`
}
