#!/bin/bash

if [ "$1" = "" ];then
  echo "[ERROR] Required argument <page-file-name>"
  exit 1
fi

echo "generate page:$1 related files …"

# 生成するファイルを配置するディレクトリを事前に定義
s_dir="./app/screen/"
c_dir="./app/containers/"
t_dir="./test/screen/"
dirs=("$s_dir" "$c_dir" "$t_dir")

for d in ${dirs[@]}; do
  dir=$(dirname $1)
  file=$(basename $1)
  mkdir -p "$d"/"$dir"
  touch "$d"/"$dir"/"$file".tsx
done
