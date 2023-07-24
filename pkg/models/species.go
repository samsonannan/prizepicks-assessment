package models

import (
	"errors"
	"strings"

	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
)

// DinosaurSpecies is a custom type representing the various dinosaur species.
type DinosaurSpecies string

// Constants for different dinosaur species using the DinosaurSpecies custom type.
const (
	Tyrannosaurus DinosaurSpecies = "TYRANNOSAURUS"
	Velociraptor  DinosaurSpecies = "VELOCIRAPTOR"
	Spinosaurus   DinosaurSpecies = "SPINOSAURUS"
	Megalosaurus  DinosaurSpecies = "MEGALOSAURUS"

	Brachiosaurus DinosaurSpecies = "BRACHIOSAURUS"
	Stegosaurus   DinosaurSpecies = "STEGOSAURUS"
	Ankylosaurus  DinosaurSpecies = "ANKYLOSAURUS"
	Triceratops   DinosaurSpecies = "TRICERATOPS"
)

// Groups is a map that associates each dinosaur species with its corresponding group.
var Groups map[DinosaurSpecies]dinosaur.Group

// Initialize the Groups map with the dinosaur species and their corresponding groups.
func init() {
	Groups = map[DinosaurSpecies]dinosaur.Group{
		Tyrannosaurus: dinosaur.GroupCARNIVORE,
		Velociraptor:  dinosaur.GroupCARNIVORE,
		Spinosaurus:   dinosaur.GroupCARNIVORE,
		Megalosaurus:  dinosaur.GroupCARNIVORE,
		Brachiosaurus: dinosaur.GroupHERBIVORE,
		Stegosaurus:   dinosaur.GroupHERBIVORE,
		Ankylosaurus:  dinosaur.GroupHERBIVORE,
		Triceratops:   dinosaur.GroupHERBIVORE,
	}
}

// GetGroup returns the group of the dinosaur species provided as input.
// It takes a string representing the dinosaur species and returns the corresponding dinosaur.Group.
// If the provided species is not found in the Groups map, it returns an error.
func GetGroup(species string) (dinosaur.Group, error) {
	// Convert the species string to DinosaurSpecies custom type.
	dinoSpecies := DinosaurSpecies(strings.ToUpper(species))

	// Check if the species exists in the Groups map.
	group, ok := Groups[dinoSpecies]
	if !ok {
		// If the species is not found in the Groups map, return an error.
		return "", errors.New("failed to identify dinosaur species")
	}

	// Return the corresponding group for the species.
	return group, nil
}
