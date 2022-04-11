#!/usr/bin/env bash

set -e

w6p install kind -v 0.12.0
w6p install kubectl
w6p install krew
w6p install k9s -v 0.25.18
w6p install helm -v 3.8.1

./kind version >/dev/null 2>&1
./kubectl version --short --client >/dev/null 2>&1
./kubectl-krew version >/dev/null 2>&1
./k9s version --short >/dev/null 2>&1
./helm version --short >/dev/null 2>&1

mkdir -p ~/bin
mv kind kubectl k9s helm ~/bin/
echo "export PATH=$PATH:~/bin:~/.krew/bin" >> ~/.bashrc

cat <<EOF >> ~/.bash_aliases
alias k=kubectl
alias kns='kubectl config set-context --current --namespace'
alias cns='kubectl config view --minify | grep namespace'
EOF

cat <<EOF >> ~/.bashrc
export do='--dry-run=client -o yaml'
export force='--grace-period=0 --force'
source <(kubectl completion bash)
complete -F __start_kubectl k
EOF

cat <<EOF > ~/.vimrc
set tabstop=2
set expandtab
set shiftwidth=2
EOF

source ~/.bashrc

echo "Setup succeeded!"
