#!/bin/sh

GITHUB_RELEASES="https://api.github.com/repos/nvbn/shell_logger/releases/latest"
RELEASE=$(curl -sL $GITHUB_RELEASES|grep "/download/" | grep $(uname -s | tr '[:upper:]' '[:lower:]')|awk '{ print $2 }'|cut -d '"' -f2)

sudo sh -c "curl -sL $RELEASE > /usr/local/bin/shell_logger && chmod +x /usr/local/bin/shell_logger"

CONFIGURATION='eval $(shell_logger --configure)'
grep -q -F "$CONFIGURATION" ~/.zshrc || echo "$CONFIGURATION" >> ~/.zshrc

echo 'Run `source ~/.zshrc to apply changes`'
