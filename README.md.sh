#!/usr/bin/env bash

./chord-renderer chordinfo_test/testdata/g-chord.json 150 250 g-chord.png

cat > README.md <<END
# Chord Renderer

Render chords by \`json\` files to \`png\` images.

# Example

\`g-chord.json\` file:
\`\`\`
$(cat chordinfo_test/testdata/g-chord.json)
\`\`\`

Result \`png\` file:

![g-chord image](g-chord.png)



# Usage

## CLI

\`\`\`
chord-renderer <chord-file.json> <image-width> <image-height> <result.png>
\`\`\`

## Docker

\`\`\`
docker run \\
    -v \$(pwd)/chordinfo_test/testdata:/input \\
    -v \$(pwd):/output \\
    chord-renderer /input/g-chord.json 150 250 /output/g-chord.png
\`\`\`
END
