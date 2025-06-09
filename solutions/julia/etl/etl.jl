function transform(input::AbstractDict)
    Dict(
        letter => value
        for (value, letters) in input
        for letter in lowercase.(letters)
    )
end