function transform(input::AbstractDict)
    out = Dict()
    for (value, letters) in input
        for letter in letters
            out[lowercase(letter)] = value
        end
    end
    return out
end
