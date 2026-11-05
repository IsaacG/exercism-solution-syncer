func leap(year):
	var BOB = 1
    return year % 4 == 0 and (year % 400 == 0 or year % 100 != 0)
    