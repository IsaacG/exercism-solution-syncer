@export var position : Vector2i
@export var direction : String

var rot = {
	"L": {
		"north": "west",
		"west": "south",
		"south": "east",
		"east": "north",
	},
	"R": {
		"north": "east",
		"east": "south",
		"south": "west",
		"west": "north",
	},
}
var dir = {"north": Vector2i.DOWN, "east": Vector2i.RIGHT, "south": Vector2i.UP, "west": Vector2i.LEFT}


func move(instructions: String):
	for instruction in instructions:
		if instruction == "A":
			position += dir[direction]
		else:
			direction = rot[instruction][direction]

