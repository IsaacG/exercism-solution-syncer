var COLORS = ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"]
var TOLERANCES = {
	"grey": 0.05,
	"violet": 0.1,
	"blue": 0.25,
	"green": 0.5,
	"brown": 1,
	"red": 2,
	"gold": 5,
	"silver": 10,
}
var UNITS = ["ohms", "kiloohms", "megaohms"]


func color_code(colors):
	if len(colors) == 1:
		return "0 ohms"
	var values = colors.slice(0, len(colors) - 2)
	var multiplier = colors[-2]
	var tolerance = colors[-1]

	var resistance = 0.0
	for value in values:
		resistance = resistance * 10 + COLORS.find(value)
	resistance *= 10 ** COLORS.find(multiplier)

	var power = 0
	while resistance >= 1000:
		resistance /= 1000
		power += 1
	
	var r = "%f" % resistance
	return "%s %s ±%s%%" % [r.rstrip("0").rstrip("."), UNITS[power], TOLERANCES[tolerance]]
