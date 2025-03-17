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
    683 # Fix issue with ping ponging between the two entrances to the same area
    684 # Enable/Repair game.ShiftKey ( shift + left click )
    713 # Walkable logic ( door handling and monsters) 
    716 # Companion Mode + New Flow
    736 # Frostnova orb sorceress
    738 # Hydra/Ball Sorceress
    733 # better gambling circlets/coronets
    731 # potions to belt after death
    703 # add configurable Kill Ghosts? to Countess run
    690 # update Terror_Zone
    685 # update to Sorceress_Leveling
    674 # update to level_tools to improve allocation of stats/skills
    646 # Supervisor names in Discord Integration
    642 # Summoner red portal exit
    742 # automatically changes runs based and gives warnings when Leveling class selected
    
    # 729 # Optimize attack.go and move.go  -- this one has conflicts, not sure about it yet
)

for pr in ${prs[@]}

do
    merge_pr $pr
done