Erstelle sinnvolle Aliase:

```sh
cat <<EOF >> ~/.bash_aliases
alias k=kubectl
alias e='ETCDCTL_API=3 etcdctl'
alias kns='kubectl config set-context --current --namespace'
alias cns='kubectl config view --minify | grep namespace'
EOF
```


Exportiere häufig verwendete env Variablen und installiere kubectl Bash-Completion:

```sh
cat <<EOF >> ~/.bashrc
export dy='--dry-run=client -o yaml'
export force='--grace-period=0 --force'
source <(kubectl completion bash)
complete -F __start_kubectl k
EOF
source ~/.bashrc
```


Konfiguriere vim (falls genutzt) für den reibungslosen Umgang mit yaml Dateien:

```sh
cat <<EOF > ~/.vimrc
set tabstop=2
set expandtab
set shiftwidth=2
EOF
```
