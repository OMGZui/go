#!/bin/zsh

deploy_time=$(date "+%Y-%m-%d %H:%M:%S")
echo "---开始构建---\n"
git add .
git commit -m "$deploy_time"
git push origin master
echo "\n"
echo "---构建完成---"