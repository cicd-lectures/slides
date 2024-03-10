package main

import (
	"testing"
)

// tag::basic[]
func TestToStringSlice(t *testing.T) {
	// Given.
	input := []int{1, 2, 3, 4}

	// When.
	got := ToStringSlice(input)

	// Then.
	want := []string{"1", "2", "3", "4"}

	// On vérifie que la longueur récupérée est conforme à celle attendue.
	if len(want) != len(got) {
		t.Fail("Mismatched length")
		return
	}

	// Pour chacune des entrées récupérées
	// On vérifie que l'entrée correspond à la valeur attendue.
	for i, gotValue := range got {
		if want[i] != gotValue {
			t.Fail()
		}
	}
}

// end::basic[]

// tag::table_driven[]

func TestToStringSlice(t *testing.T) {
	// On crée un tableau de cas de tests
	testCases := []struct {
		name  string
		input []int
		want  []string
	}{
		{
			name:   "normal case",
			input:  []int{1, 2, 3, 4},
			output: []string{"1", "2", "3", "4"},
		},
		{
			name:   "empty case",
			input:  nil,
			output: nil,
		},
	}

	// Pour chacun de ces cas de tests.
	for _, testCase := range testCases {
		// On execute un sous test qui utilises les attributs
		// du cas de test.
		t.Run(testCase.name, func(t *testing.T) {
			got := ToStringSlice(testCase.input)

			if len(testCase.want) != len(got) {
				t.Fail("Mismatched length")
				return
			}

			for i, gotValue := range got {
				if testCase.want[i] != gotValue {
					t.Fail()
				}
			}
		})
	}
}

// end::table_driven[]
