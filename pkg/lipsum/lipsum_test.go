package lipsum_test

import (
	"testing"

	lipsum "github.com/alwian/go-lipsum/pkg"
)

// generate

func TestGeneratePassesCorrectParams(t *testing.T) {
	_, err := lipsum.Bytes(0)
	if err == nil || err.Error() != "invalid amount '0' < 1" {
		t.Fail()
	}
}

// Bytes

func TestBytesPassesCorrecParams(t *testing.T) {
	_, err := lipsum.Bytes(0)
	if err == nil || err.Error() != "invalid amount '0' < 1" {
		t.Fail()
	}
}

// Words

func TestWordsPassesCorrecParams(t *testing.T) {
	_, err := lipsum.Words(0)
	if err == nil || err.Error() != "invalid amount '0' < 1" {
		t.Fail()
	}
}

// Paragraphs

func TestParagraphsPassesCorrecParams(t *testing.T) {
	_, err := lipsum.Paragraphs(0)
	if err == nil || err.Error() != "invalid amount '0' < 1" {
		t.Fail()
	}
}
