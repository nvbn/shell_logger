#!/bin/sh

GITHUB_RELEASES="https://api.github.com/repos/nvbn/shell_logger/releases/latest"
RELEASE=$(curl -sL $GITHUB_RELEASES|grep "/download/" | grep $(uname -s | tr '[:upper:]' '[:lower:]')|awk '{ print $2 }'|cut -d '"' -f2)

sudo sh -c "curl -sL $RELEASE > /usr/local/bin/shell_logger && chmod +x /usr/local/bin/shell_logger"

CONFIGURATION='eval $(shell_logger --configure)'

if [ "$(basename $SHELL)" = "zsh" ]; then
    grep -q -F "$CONFIGURATION" ~/.zshrc || echo "$CONFIGURATION" >> ~/.zshrc

    echo 'Run `source ~/.zshrc` to apply changes'
else
    test -f ~/.bash-preexec.sh || curl https://raw.githubusercontent.com/rcaloras/bash-preexec/master/bash-preexec.sh -o ~/.bash-preexec.sh

    grep -q -F "source ~/.bash-preexec.sh" ~/.bashrc || echo '[[ -f ~/.bash-preexec.sh ]] && source ~/.bash-preexec.sh' >> ~/.bashrc

    grep -q -F "$CONFIGURATION" ~/.bashrc || echo "$CONFIGURATION" >> ~/.bashrc

    echo 'Run `source ~/.bashrc` to apply changes'
fi;
