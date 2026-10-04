var COLORS = ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"]

func color_code(colors):
	return COLORS.find(colors[0]) * 10 + COLORS.find(colors[1])
