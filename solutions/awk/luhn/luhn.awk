BEGIN { FS = "" }
#
# Check for invalid chars
/[^0-9 ]/ {
    print "false"
    next
}

# Valid chars only
{
    # Drop any spaces
    gsub(/ /, "")
    double = !(length % 2)
    sum = 0
    for ( i = 1; i <= NF; i++ ) {
        val = $i
        if (double) {
            val = val * 2
            if (val > 9) {
                val = val - 9
            }
        }
        sum += val
        double = !double
    }
    print (sum % 10 || length < 2) ? "false" : "true"
}
