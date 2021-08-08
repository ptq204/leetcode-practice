#!/bin/bash/

untracked_files=$(git status -s -b | awk '{ if ($1 == "??") { print $2 } }')

commit_msg="solved:"

for ufile in $untracked_files
do
    git add "$ufile"
    IFS='-' read -a ufile_name <<< "$ufile"
    commit_msg+=" ${ufile_name[0]}"
done

echo "$commit_msg"
git commit -m "$commit_msg"
