#!/bin/bash


yum install zsh

sleep 3

wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | sh

sleep 3


cp ./robbyrussell.zsh-theme /root/.oh-my-zsh/themes -rf

sleep 3

chsh -s /bin/zsh