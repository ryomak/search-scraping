package search

type Articles []Article
type Article struct {
	Title  string
	IsMine bool
	URL    string
	Rank   uint
}

type Result struct {
	MyURL         string
	SearchResults []SearchResult
}

type SearchResult struct {
	Keyword  string
	Articles []Article
}
