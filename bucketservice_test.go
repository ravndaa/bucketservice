package main

import "testing"

var bck = Init()

func TestCreate(t *testing.T) {

	expectedLength := 9
	actual, err := bck.Create()

	if len(actual) != expectedLength {
		t.Errorf("Expected the Create to be %d but instead got %d!", expectedLength, len(actual))
	} else if err != nil {
		t.Errorf("Expcted error to be nil, but got: %s", err)
	}

}

func TestDelete(t *testing.T) {

	bucket, _ := bck.Create()
	err := bck.Delete(bucket)

	if err != nil {
		t.Errorf("Expcted error to be nil, but got: %s", err)
	}

	if _, ok := bck.DB[bucket]; ok {
		t.Errorf("Not expecting to find bucket.")
	}

}
