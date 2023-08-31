package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Wahyu",
			request:  "Wahyu",
			expected: "Hello Wahyu",
		},
		{
			name:     "Adi",
			request:  "Adi",
			expected: "Hello Adi",
		},
		{
			name:     "Kusuma",
			request:  "Kusuma",
			expected: "Hello Kusuma",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

/**
SUB TEST :
- Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah :
go test -run TestNamaFunction
- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah :
go test -run TestNamaFunction/NamaSubTest
- Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah :
go test -run /NamaSubTest
*/

func TestSubTest(t *testing.T) {
	t.Run("SubTestWahyu", func(t *testing.T) {
		result := HelloWorld("Wahyu")
		require.Equal(t, "Hello Wahyu", result, "Result must be 'Hello Wahyu'")
	})
	t.Run("SubTestAdi", func(t *testing.T) {
		result := HelloWorld("Adi")
		require.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	})
}

// before and after test
func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

// skip test (jika berada di kondisi tertentu)
func TestSkip(t *testing.T) {
	// kalo run di OS windows, test akan di skip
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}
	result := HelloWorld("Wahyu")
	require.Equal(t, "Hello Wahyu", result, "Result must be 'Hello Wahyu'")
	// Package require implements the same assertions as the `assert` package but stops test execution when a test fails.
}

// menggunakan testify (library from github)
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Wahyuu")
	assert.Equal(t, "Hello Wahyuu", result, "Result must be 'Hello Wahyuu'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Wahyu")

	if result != "Hello Wahyu" {
		// t.FailNow() // kalo FailNow langsung gagal tanpa menjalankan program sampai akhir
		t.Fatal("Must be 'Hello Wahyu") //kalo fatal, habis nampilin error log dia langsung manggil failNow
	}
	fmt.Println("Wahyu")
}

func TestHelloWorldWahyu(t *testing.T) {
	result := HelloWorld("Mas Wahuy")

	if result != "Hello Mas Wahuy" {
		// t.Fail() // kalo Fail, tetap menjalankan program sampai akhir
		t.Error("Result must be 'Hello Mas Wahuy'")
	}
	fmt.Println("mantap")
}
