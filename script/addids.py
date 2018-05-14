import sys
import re

pattern = re.compile(r"_:((topic|link)_\w+)")
pattern_owned = re.compile(r"_:organization_tyrell <owns> _:((topic|link)_\w+)")
resources = set()
owned = set()

for line in sys.stdin:
    m = pattern_owned.search(line)
    if m:
        owned.add(m.group(1))
    else:
        m = pattern.search(line)
        if m:
            resources.add(m.group(1))

for free in resources - owned:
    print('_:organization_tyrell <owns> _:{} .'.format(free))
