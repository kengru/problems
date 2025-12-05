#!/bin/bash

# Script to create a new Advent of Code day
# Usage: ./scripts/new-advent-day.sh YEAR DAY
# Example: ./scripts/new-advent-day.sh 2025 6

set -e

if [ $# -ne 2 ]; then
    echo "Usage: $0 YEAR DAY"
    echo "Example: $0 2025 6"
    exit 1
fi

YEAR=$1
DAY=$(printf "%02d" $2)
DAY_NUM=$2

PROJECT_ROOT=$(cd "$(dirname "$0")/.." && pwd)
DAY_DIR="$PROJECT_ROOT/advent/$YEAR/$DAY"
REGISTRY_FILE="$PROJECT_ROOT/registry.go"

# Check if day already exists
if [ -d "$DAY_DIR" ]; then
    echo "Error: Day $DAY for year $YEAR already exists at $DAY_DIR"
    exit 1
fi

# Create the day directory
mkdir -p "$DAY_DIR"

# Create part1.go
cat > "$DAY_DIR/part1.go" << EOF
package advent

import (
	"fmt"

	"github.com/kengru/problems/advent"
)

func Year${YEAR}${DAY}01() {
	lines := advent.GetLinesFromFile("advent/$YEAR/$DAY/ie.txt")
	// lines := advent.GetLinesFromFile("advent/$YEAR/$DAY/input.txt")

	fmt.Println(len(lines))
}
EOF

# Create part2.go
cat > "$DAY_DIR/part2.go" << EOF
package advent

import (
	"fmt"

	"github.com/kengru/problems/advent"
)

func Year${YEAR}${DAY}02() {
	lines := advent.GetLinesFromFile("advent/$YEAR/$DAY/ie.txt")
	// lines := advent.GetLinesFromFile("advent/$YEAR/$DAY/input.txt")

	fmt.Println(len(lines))
}
EOF

# Create empty input files
touch "$DAY_DIR/input.txt"
touch "$DAY_DIR/ie.txt"

# Update registry.go
# First, add the import
IMPORT_LINE="d${YEAR}${DAY} \"github.com/kengru/problems/advent/$YEAR/$DAY\""

# Find the right place to insert the import (after the last import for this year, or create new year section)
# We'll use a simple approach: insert before the closing paren of imports

# Check if import already exists
if grep -q "d${YEAR}${DAY}" "$REGISTRY_FILE"; then
    echo "Warning: Import for d${YEAR}${DAY} already exists in registry"
else
    # Find the line with the closing paren of imports and insert before it
    # We need to find the right year section or create one

    # Check if this year has any imports already
    if grep -q "advent/$YEAR/" "$REGISTRY_FILE"; then
        # Find the last import for this year and add after it
        LAST_IMPORT_LINE=$(grep -n "advent/$YEAR/" "$REGISTRY_FILE" | tail -1 | cut -d: -f1)
        sed -i '' "${LAST_IMPORT_LINE}a\\
	${IMPORT_LINE}
" "$REGISTRY_FILE"
    else
        # New year - add before the closing paren with a comment
        IMPORT_CLOSE_LINE=$(grep -n "^)" "$REGISTRY_FILE" | head -1 | cut -d: -f1)
        sed -i '' "${IMPORT_CLOSE_LINE}i\\
\\
	${IMPORT_LINE}
" "$REGISTRY_FILE"
    fi
fi

# Add registry entries
REGISTRY_ENTRY_1="\"${YEAR}${DAY}01\": d${YEAR}${DAY}.Year${YEAR}${DAY}01,"
REGISTRY_ENTRY_2="\"${YEAR}${DAY}02\": d${YEAR}${DAY}.Year${YEAR}${DAY}02,"

# Check if this year has entries already
if grep -q "\"${YEAR}.*\":" "$REGISTRY_FILE"; then
    # Find the last entry for this year and add after it
    LAST_ENTRY_LINE=$(grep -n "\"${YEAR}[0-9]*\":" "$REGISTRY_FILE" | tail -1 | cut -d: -f1)
    sed -i '' "${LAST_ENTRY_LINE}a\\
	${REGISTRY_ENTRY_1}\\
	${REGISTRY_ENTRY_2}
" "$REGISTRY_FILE"
else
    # New year - add before the closing brace with a comment
    REGISTRY_CLOSE_LINE=$(grep -n "^}" "$REGISTRY_FILE" | head -1 | cut -d: -f1)
    sed -i '' "${REGISTRY_CLOSE_LINE}i\\
\\
	// ${YEAR}\\
	${REGISTRY_ENTRY_1}\\
	${REGISTRY_ENTRY_2}
" "$REGISTRY_FILE"
fi

echo "Created new Advent of Code day:"
echo "  - $DAY_DIR/part1.go"
echo "  - $DAY_DIR/part2.go"
echo "  - $DAY_DIR/input.txt"
echo "  - $DAY_DIR/ie.txt"
echo "  - Updated registry.go"
echo ""
echo "Run with: go run . --year $YEAR --day $DAY_NUM --part 1"
