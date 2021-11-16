#!/usr/bin/env bash

# export IMPLEMENTATION_ESTIMATE="1d"
# export CODEMOD_ESTIMATE="2d"
# export LINT_RULE_ESTIMATE="0.5d"
# export MANUAL_MIGRATION_ESTIMATE="1d"
# ./create-wildcard-issue.sh "<Link />" "link"

export IMPLEMENTATION_ESTIMATE="0.5d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE="0.5d"
export MANUAL_MIGRATION_ESTIMATE="1d"
./create-wildcard-issue.sh "<Icon />" "icon"

export IMPLEMENTATION_ESTIMATE="1d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="2d"
./create-wildcard-issue.sh "<IconButton />" "icon-button"

export IMPLEMENTATION_ESTIMATE="2d"
export CODEMOD_ESTIMATE="3d"
export LINT_RULE_ESTIMATE="1d"
export MANUAL_MIGRATION_ESTIMATE="0.5d"
./create-wildcard-issue.sh "<Typography />" "typography"

export IMPLEMENTATION_ESTIMATE="2d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE="0.5d"
export MANUAL_MIGRATION_ESTIMATE="2d"
./create-wildcard-issue.sh "<Card />" "card"

export IMPLEMENTATION_ESTIMATE="0.5d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="0.5d"
./create-wildcard-issue.sh "<MonacoInput />" "monaco-input"

export IMPLEMENTATION_ESTIMATE="1d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE="0.5d"
export MANUAL_MIGRATION_ESTIMATE="0.5d"
./create-wildcard-issue.sh "<Alert />" "alert"

export IMPLEMENTATION_ESTIMATE="0.5d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="0.5d"
./create-wildcard-issue.sh "<Tooltip />" "tooltip"

export IMPLEMENTATION_ESTIMATE="3d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="4d"
./create-wildcard-issue.sh "<Popover />" "popover"

export IMPLEMENTATION_ESTIMATE="2d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="2d"
./create-wildcard-issue.sh "<Prompt />" "prompt"

export IMPLEMENTATION_ESTIMATE="1d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="3d"
./create-wildcard-issue.sh "<Menu />, <MenuItem />" "menu"

export IMPLEMENTATION_ESTIMATE="2d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="2d"
./create-wildcard-issue.sh "<NavMenu />" "nav-menu"

export IMPLEMENTATION_ESTIMATE="4d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="5d"
./create-wildcard-issue.sh "<Combobox />" "combobox"

export IMPLEMENTATION_ESTIMATE="3d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="3d"
./create-wildcard-issue.sh "<Panel />" "panel"

export IMPLEMENTATION_ESTIMATE="0.5d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="3d"
./create-wildcard-issue.sh "<FeedbackWidget />" "feedback-widget"

export IMPLEMENTATION_ESTIMATE="2d"
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="1d"
./create-wildcard-issue.sh "<FileBreadcrumbs />" "file-breadcrumbs"

export IMPLEMENTATION_ESTIMATE=""
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE="0.5d"
export MANUAL_MIGRATION_ESTIMATE="1d"
./create-wildcard-issue.sh "<Checkbox />" "checkbox"

export IMPLEMENTATION_ESTIMATE=""
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE="0.5d"
export MANUAL_MIGRATION_ESTIMATE="1d"
./create-wildcard-issue.sh "<Input />" "input"

export IMPLEMENTATION_ESTIMATE=""
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE="0.5d"
export MANUAL_MIGRATION_ESTIMATE="1d"
./create-wildcard-issue.sh "<RadioButton />" "radio-button"

export IMPLEMENTATION_ESTIMATE=""
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE="0.5d"
export MANUAL_MIGRATION_ESTIMATE="1d"
./create-wildcard-issue.sh "<Select />" "select"

export IMPLEMENTATION_ESTIMATE=""
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE="0.5d"
export MANUAL_MIGRATION_ESTIMATE="1d"
./create-wildcard-issue.sh "<Textarea />" "textarea"

export IMPLEMENTATION_ESTIMATE=""
export CODEMOD_ESTIMATE=""
export LINT_RULE_ESTIMATE=""
export MANUAL_MIGRATION_ESTIMATE="3d"
./create-wildcard-issue.sh "<Tabs />" "tabs"
