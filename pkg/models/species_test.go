package models

import (
	"testing"

	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
)

func TestGetGroup(t *testing.T) {
	tests := []struct {
		species       string
		expectedGroup dinosaur.Group
		expectedErr   bool
	}{
		// Test cases for carnivore species.
		{"Tyrannosaurus", dinosaur.GroupCARNIVORE, false},
		{"Velociraptor", dinosaur.GroupCARNIVORE, false},
		{"Spinosaurus", dinosaur.GroupCARNIVORE, false},
		{"Megalosaurus", dinosaur.GroupCARNIVORE, false},

		// Test cases for herbivore species.
		{"Brachiosaurus", dinosaur.GroupHERBIVORE, false},
		{"Stegosaurus", dinosaur.GroupHERBIVORE, false},
		{"Ankylosaurus", dinosaur.GroupHERBIVORE, false},
		{"Triceratops", dinosaur.GroupHERBIVORE, false},

		// Test case for an unknown species.
		{"UnknownSpecies", "", true},
	}

	for _, test := range tests {
		group, err := GetGroup(test.species)
		if test.expectedErr {
			// Expecting an error.
			if err == nil {
				t.Errorf("Expected error for species %s, but got nil", test.species)
			}
		} else {
			// Expecting no error.
			if err != nil {
				t.Errorf("Unexpected error for species %s: %v", test.species, err)
			}
			// Check if the retrieved group matches the expected group.
			if group != test.expectedGroup {
				t.Errorf("Expected group %s for species %s, but got %s", test.expectedGroup, test.species, group)
			}
		}
	}
}
