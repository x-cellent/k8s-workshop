Installiere die folgenden Tools:

```sh
w6p install kind -v 0.12.0
w6p install kubectl
w6p install krew
w6p install k9s -v 0.25.18
w6p install helm -v 3.8.1
```

======================================================================================================
Die Tools sollten nun in deinem aktuellen Verzeichnis liegen.
Damit du sie bequem ausführen kannst, mache folgendes:

```sh
mkdir -p ~/bin
mv kind kubectl k9s helm ~/bin/
echo "export PATH=$PATH:~/bin:~/.krew/bin" >> ~/.bashrc
source ~/.bashrc
```


======================================================================================================
Vergewissere dich, dass alle installierten Tools funktionieren:

```sh
kind version >/dev/null 2>&1 && echo "kind works" || echo "kind does not work"
kubectl version --short --client >/dev/null 2>&1 && echo "kubectl works" || "kubectl does not work"
kubectl-krew version >/dev/null 2>&1 && echo "krew works" || "krew does not work"
k9s version --short >/dev/null 2>&1 && echo "k9s works" || "k9s does not work"
helm version --short >/dev/null 2>&1 && echo "helm works" || echo "helm does not work"
```


======================================================================================================
Erstelle sinnvolle Aliase:

```sh
cat <<EOF >> ~/.bash_aliases
alias k=kubectl
alias kns='kubectl config set-context --current --namespace'
alias cns='kubectl config view --minify | grep namespace'
EOF
```


======================================================================================================
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


======================================================================================================
Konfiguriere vim (falls genutzt) für den reibungslosen Umgang mit yaml Dateien:

```sh
cat <<EOF > ~/.vimrc
set tabstop=2
set expandtab
set shiftwidth=2
EOF
```
