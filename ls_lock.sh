#!/bin/bash
set -o errexit

if [ $# -ne 1 ]; then
  me=`basename $0`
  echo "No directory name given."
  echo "Usage $me: my_dicrectory_with_locks"
  exit 1
fi

target_dir=$1

index_flock_inodes=()

# credit goes to http://stackoverflow.com/a/14550606/410823
make_flock_index () {
  local -a value_array=("$@")
  local i
  for i in "${!value_array[@]}"; do
    eval index_flock_inodes["${value_array[$i]}"]=$i
  done
}

inode_file_pairs=( $(find ${target_dir} -name "*" | xargs -I {} stat --printf="%i,{}\n" {} | sort) )
inode_of_flocks=( $(cat /proc/locks  | grep FLOCK | cut -d' ' -f8 | cut -d':' -f3 | sort) )

make_flock_index "${inode_of_flocks[@]}"

printf "%10s %s\n" "PID" "Lockedfile";

for inode_file in "${inode_file_pairs[@]}"
do
    inode=${inode_file%,*}  
    lock_file=${inode_file#*.}
    # inode="12848"  
    if [ "${index_flock_inodes[$inode]}" ]; then 
        # IMPROVEMENT: create a IDX from lsof output
        pid=$(lsof  | grep  ${lock_file}  | cut -d' ' -f2)
        printf "%10s %s\n" ${pid} ${lock_file}
    fi     
done
