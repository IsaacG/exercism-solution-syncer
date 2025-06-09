function sum_of_multiples(limit, factors)
    found = Set{Int64}()
    for factor in factors
        if factor == 0
            continue
        end
        for i in range(factor, step=factor, stop=limit)
            if i < limit
                push!(found, i)
            end
        end
    end
    return sum(found)
end
