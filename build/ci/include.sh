#!/usr/bin/env bash

NORMAL_FOREGROUND="\033[0m"
GREEN_FOREGROUND="\033[32m"
YELLOW_FOREGROUND="\033[33m"
BLUE_FOREGROUND="\033[36m"
RED_FOREGROUND="\033[31m"
CHECK_MARK="✅"
ARROW="➜"
CROSS_MARK="❌"

# Prints a message based on the given code
# $1 - Exit code of a command
# $2 - Message on success
# $3 - Message on error
function f_process_command {
    if [[ $1 == 0 ]]
    then
        echo -e "$GREEN_FOREGROUND$CHECK_MARK $2"
        echo -e "$NORMAL_FOREGROUND"
    else
        echo -e "$RED_FOREGROUND$CROSS_MARK $3"
        echo -e "$NORMAL_FOREGROUND"
        exit 1
    fi
}

# Prints an info message
# $1 - The message
function f_info_log {
    echo -e "$BLUE_FOREGROUND$ARROW $1"
    echo -e "$NORMAL_FOREGROUND"
}

# Prints an error message
# $1 - The message
function f_error_log {
    echo -e "$RED_FOREGROUND$CROSS_MARK $1"
    echo -e "$NORMAL_FOREGROUND"
}

# Prints a success message
# $1 - The message
function f_success_log {
    echo -e "$GREEN_FOREGROUND$CHECK_MARK $1"
    echo -e "$NORMAL_FOREGROUND"
}
