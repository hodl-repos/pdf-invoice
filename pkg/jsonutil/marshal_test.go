package jsonutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestMarshalResponse(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()

	toSave := map[string]string{
		"name": "John",
	}

	MarshalResponse(w, http.StatusOK, toSave)

	if w.Code != http.StatusOK {
		t.Errorf("wrong response code, want: %v got: %v", http.StatusOK, w.Code)
	}

	got := w.Body.String()
	want := `{"name":"John"}`
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unmarshal mismatch (-want +got):\n%v", diff)
	}
}

func TestMarshalResponseWithError(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()

	toSave := map[string]string{
		"name": "John",
	}

	err := MarshalResponseWithError(w, http.StatusOK, toSave)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code, "wrong response code")

	assert.Equal(t, `{"name":"John"}`, w.Body.String(), "unmarshal mismatch")
}

func TestMarshalResponseError(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()

	type Circular struct {
		Name string    `json:"name"`
		Next *Circular `json:"next"`
	}

	badInput := &Circular{
		Name: "Bob",
	}
	badInput.Next = badInput

	MarshalResponse(w, http.StatusInternalServerError, badInput)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong response code, want: %v got: %v", http.StatusOK, w.Code)
	}

	got := w.Body.String()
	want := `` //no body
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unmarshal mismatch (-want +got):\n%v", diff)
	}
}

func TestMarshalResponseErrorWithError(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()

	type Circular struct {
		Name string    `json:"name"`
		Next *Circular `json:"next"`
	}

	badInput := &Circular{
		Name: "Bob",
	}
	badInput.Next = badInput

	err := MarshalResponseWithError(w, http.StatusOK, badInput)

	assert.NotNil(t, err)
	assert.Equal(t, "", w.Body.String())
}
