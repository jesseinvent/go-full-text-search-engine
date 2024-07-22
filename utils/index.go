package utils

type Index map[string][]int

// Add documents to index
func (index Index) Add(docs []Document) {
	for _, doc := range docs {
		// Tokenize & Analyze each text in doc
		for _, token := range Analyze(doc.Text) {
			indexToken := index[token]

			// Check in token with doc id has already been indexed
			if indexToken != nil && indexToken[len(indexToken)-1] == doc.ID {
				// Don't add same ID twice
				continue
			}

			index[token] = append(indexToken, doc.ID)
		}
	}
}

// a n b
func Intersection(a []int, b []int) []int {
	maxLen := len(a)

	if len(b) > maxLen {
		maxLen = len(b)
	}

	r := make([]int, 0, maxLen)

	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}

	return r
}

func (index Index) Search(text string) []int {
	var r []int
	for _, token := range Analyze(text) {
		if indexToken, ok := index[token]; ok {
			if r == nil {
				r = indexToken
			} else {
				r = Intersection(r, indexToken)
			}
		} else {
			// Tpken does not exist
			return nil
		}
	}

	return r
}
