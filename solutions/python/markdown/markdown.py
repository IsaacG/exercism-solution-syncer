import re
import enum


class Tag(enum.Enum):
    """Represents various tags that can wrap text."""
    LIST = 'ul'
    BULLET = 'li'
    PARAGRAPH = 'p'
    EMPHASIS = 'em'
    STRONG = 'strong'
    HEADER1 = 'h1'
    HEADER2 = 'h2'
    HEADER3 = 'h3'
    HEADER4 = 'h4'
    HEADER5 = 'h5'
    HEADER6 = 'h6'
    HEADER7 = 'h7'


# Indexed headers, index matching header number.
HEADERS = (None, Tag.HEADER1, Tag.HEADER2, Tag.HEADER3, Tag.HEADER4, Tag.HEADER5, Tag.HEADER6, Tag.HEADER7)
# Midline wraps, sorted by processing order.
MIDLINE_WRAP = (('__', Tag.STRONG), ('_', Tag.EMPHASIS))


def _Wrap(line, tag):
    """Wrap text in a tag."""
    return f'<{tag.value}>{line}</{tag.value}>'


def _ParseMidLineWraps(line: str) -> str:
    """Parse all mid-line text wrappers.

    Transforms e.g. '__a__' into '<strong>a</strong>'.
    Operates on a whole line, looping until no more wraps found.
    """
    while True:
        for md, token in MIDLINE_WRAP:
            if re.match(f'(.*){md}(.*){md}(.*)', line):
                pre, text, post = line.split(md, 2)
                line = pre + _Wrap(text, token) + post
                break
        else:
            return line


def parse(markdown):
    lines = markdown.split('\n')
    res = ''
    in_list = False
    for i in lines:
        # Close out a list if needed.
        if in_list and not i.startswith('*'):
            in_list = False
            res += '</ul>'

        # Default line wrapper, if none others found.
        line_wrap = Tag.PARAGRAPH
        # Determine how to wrap this line.
        if i.startswith('#'):
            header, i = i.split(' ', 1)
            line_wrap = HEADERS[len(header)]
        elif i.startswith('*'):
            line_wrap = Tag.BULLET
            i = i.split(' ', 1)[1]
            if not in_list:
                res += '<ul>'
            in_list = True
        # Parse the line's contents.
        i = _ParseMidLineWraps(i)

        res += _Wrap(i, line_wrap)

    # Close out any list we may be in.
    if in_list:
        res += '</ul>'
    return res


# vim:ts=4:sw=4:expandtab
