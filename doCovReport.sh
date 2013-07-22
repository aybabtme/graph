#!/bin/sh

cmd="gocov test -deps -exclude-goroot ./... | gocov report >> coverage.md"

echo "# Coverage" > coverage.md
echo "Result of running \`$cmd\` on this package." >> coverage.md
echo "\`\`\`" >> coverage.md 
gocov test -deps -exclude-goroot ./... | gocov report >> coverage.md
echo "\`\`\`" >> coverage.md
