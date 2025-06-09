#!/usr/bin/env gawk -f

BEGIN {
    n = split(flags, f)
    for (i = 1; i <= n; i++)
        switch (f[i]) {
        case "n": line_num = 1; break
        case "i": IGNORECASE = 1; break
        case "l": filename = 1; break
        case "x": pattern = "^" pattern "$"; break
        case "v": invert = 1; break
    }
}

($0 ~ pattern) != invert {
    found = 1
    if (filename) {
        print FILENAME
        nextfile
    }
    print (ARGC > 2 ? FILENAME ":" : "") (line_num ? FNR ":" : "") $0
}

END { if (!found) exit 1 }
