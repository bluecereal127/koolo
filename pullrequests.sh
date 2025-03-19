#!/bin/bash

if ! git config remote.upstream.url > /dev/null; then
    git remote add upstream https://github.com/hectorgimenez/koolo.git
fi
git fetch upstream

merge_pr () {
    echo merging $1
    git fetch upstream pull/$pr/head:pr-$1
    git merge pr-$1 -m pr-$1
}

prs=(
    751 # support for blood ring + blood amulet crafting
    748 # improves companion script, can follow manually played player as well
    749 # generic character + other characters added to playable classes
    # 729 # Optimize attack.go and move.go  -- this one has conflicts, not sure about it yet
)

for pr in ${prs[@]}

do
    merge_pr $pr
done