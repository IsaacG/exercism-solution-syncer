def is_valid(isbn):
    isbn = isbn.replace("-", "")
    if len(isbn) != 10:
        return False

    nums = list(isbn)
    # The last character only can be "X". Map "X" => 10
    if nums[-1] == "X":
        nums[-1] = "10"
    if not all(num.isdigit() for num in nums):
        return False

    count = sum((10 - i) * int(num) for i, num in enumerate(nums))
    return count % 11 == 0
