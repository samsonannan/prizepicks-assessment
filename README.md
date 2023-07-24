## Jurassic Park API

### Overview

It's 1993 and you're the lead software developer for the new Jurassic Park! Park operations
needs a system to keep track of the different cages around the park and the different dinosaurs
in each one. You'll need to develop a JSON formatted RESTful API to allow the builders to create
new cages. It will also allow doctors and scientists the ability to edit/retrieve the statuses of
dinosaurs and cages.

### Business Requirements

- All requests should respond with the correct HTTP status codes and a response, if necessary,
representing either the success or error conditions.
- Data should be persisted using some flavor of SQL.
- Each dinosaur must have a name.
- Each dinosaur is considered an herbivore or a carnivore, depending on its species.
- Carnivores can only be in a cage with other dinosaurs of the same species.
- Each dinosaur must have a species (See enumerated list below, feel free to add others).
- Herbivores cannot be in the same cage as carnivores.
- Use Carnivore dinosaurs like Tyrannosaurus, Velociraptor, Spinosaurus and Megalosaurus.
- Use Herbivores like Brachiosaurus, Stegosaurus, Ankylosaurus and Triceratops.