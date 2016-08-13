package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"gopher", "rehpog"},
		{"reverse", "esrever"},
		{"I'am Gopher", "rehpoG ma'I"},
	}

	for _, test := range tests {
		got := Reverse(test.input)
		if got != test.expected {
			t.Errorf("Reverse(%q)=%q. but expected %q", test.input, got, test.expected)
		}
	}
}

func TestReverseUnicode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"gopher😬😁", "😁😬rehpog"},
		{"helloété", "étéolleh"},
		{"Hello, 世界", "界世 ,olleH"},
		{"The quick brown 狐 jumped over the lazy 犬", "犬 yzal eht revo depmuj 狐 nworb kciuq ehT"},
	}

	for _, test := range tests {
		got := Reverse(test.input)
		if got != test.expected {
			t.Errorf("Reverse(%q)=%q. but expected %q", test.input, got, test.expected)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Reverse("The quick brown 狐 jumped over the lazy 犬")
	}
}

func ExampleReverse() {
	fmt.Println(Reverse("The quick brown 狐 jumped over the lazy 犬"))
	fmt.Println(Reverse("Gophers!"))

	// Output:
	// 犬 yzal eht revo depmuj 狐 nworb kciuq ehT
	// !srehpoG
}
