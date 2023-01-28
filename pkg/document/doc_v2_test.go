package document

import "testing"

func TestDocV2(t *testing.T) {
	doc := NewDocV2()

	txt := NewTextV2(doc, "Hello World")
	doc.Add(txt)

	doc.Generate()
}
