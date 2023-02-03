package imap

import (
	"fmt"
	"testing"
)

func TestChangeKeyCase(t *testing.T) {

	var data = []struct {
		in  map[string]any
		out map[string]any
	}{
		{
			in: map[string]any{
				"hello": 1,
				"world": 2,
				"test":  "haha",
			},
			out: map[string]any{
				"HELLO": 1,
				"WORLD": 2,
				"TEST":  "haha",
			},
		},
		{
			in: map[string]any{
				"hello": []int{1, 2, 34, 45, 56},
				"world": 2,
				"test":  "haha",
			},
			out: map[string]any{
				"HELLO": []int{1, 2, 34, 45, 56},
				"WORLD": 2,
				"TEST":  "haha",
			},
		},
	}

	for _, val := range data {
		res := ChangeKeyCase(val.in, CaseUpper)
		resStr := fmt.Sprintf("%v", res)
		fmt.Printf("res: %v \n", res)
		fmt.Printf("expected: %v \n", val.out)
		expected := fmt.Sprintf("%v", val.out)
		if resStr != expected {
			t.Errorf("Execute Error : \n RES : %v, EXPECT : %v", res, expected)
		}
	}

}

func TestChangeKeyCaseLower(t *testing.T) {
	var data = []struct {
		in  string
		out string
	}{
		{"helloWorld", "helloworld"},
		{"AbchjsZbcs", "abchjszbcs"},
		{"HelloWoRlD", "helloworld"},
		{"HAelloWoRlDZ", "haelloworldz"},
		{"aAelloWoRlDz", "aaelloworldz"},
	}

	for _, v := range data {
		res := changeKeyCaseLower(v.in)
		if res != v.out {
			t.Errorf("Execute Error : \n RES : %s, EXPECT : %s", res, v.out)
		}
	}
}

func TestChangeKeyCaseUpper(t *testing.T) {

	var data = []struct {
		in  string
		out string
	}{
		{"helloWorld", "HELLOWORLD"},
		{"AbchjsZbcs", "ABCHJSZBCS"},
		{"HelloWoRlD", "HELLOWORLD"},
		{"HAelloWoRlDZ", "HAELLOWORLDZ"},
		{"aAelloWoRlDz", "AAELLOWORLDZ"},
	}

	for _, v := range data {
		res := changeKeyCaseUpper(v.in)
		if res != v.out {
			t.Errorf("Execute Error : \n RES : %s, EXPECT : %s", res, v.out)
		}
	}
}
