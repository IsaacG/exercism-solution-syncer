def is_armstrong_number(number):
    num_str = str(number)
    power = len(num_str)
    total = sum(int(char)**power for char in num_str)
    return number == total
