export SHELL=/usr/bin/zsh
export HISTFILE=~/.zsh_history
echo > $HISTFILE
export SAVEHIST=100
export HISTSIZE=100
eval $(shell_logger --configure)
setopt INC_APPEND_HISTORY
