export SHELL=/bin/bash
export PS1="$ "
echo > $HISTFILE

[[ -f ~/.bash-preexec.sh ]] && source ~/.bash-preexec.sh

eval $(shell_logger --configure)
