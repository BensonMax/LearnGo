package mock

type Retriever struct {
	Content string
}

func (r Retriever) Get(url string) string {
	return r.Content
}
