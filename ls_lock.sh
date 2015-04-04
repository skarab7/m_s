#!/bin/bash
set -o errexit

index_flock_inodes=()

# credit goes to http://stackoverflow.com/a/14550606/410823
make_index () {
  local -a value_array=("$@")
  local i
  for i in "${!value_array[@]}"; do
    eval index_flock_inodes["${value_array[$i]}"]=$i
    echo "${value_array[$i]}"
  done
}


inode_file_pairs=( $(find m_s -name "*" | xargs -I {} stat --printf="%i,{}\n" {} | sort) )
inode_of_flocks=( $(cat /proc/locks  | grep FLOCK | cut -d' ' -f8 | cut -d':' -f3 | sort) )

make_index "${inode_of_flocks[@]}"

for inode_file in "${inode_file_pairs[@]}"
do
    inode=${inode_file%,*}  
    lockfile=${inode_file#*.}
    inode="12848"  
    if [ "${index_flock_inodes[$inode]}" ]; then 
        echo "FOUND" 
    fi     
done

#echo $inode_file_pairs
#echo $inode_of_flocks



