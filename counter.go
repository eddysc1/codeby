package main

import (
	"os/exec"
	"fmt"
)

func addedRemovedLoc(author string) string {
	output := ""
	isFirst := true

	for _, l := range languages {
		checkLang := fmt.Sprintf("git log --author='%s' --pretty=tformat: --numstat | grep -v '^-' | grep \\\\.%s$", author, l.extension)

		_, checkLangErr := exec.Command("bash", "-c", checkLang).Output()

		if checkLangErr != nil {
			continue
		}

		cmd := fmt.Sprintf("git log --author='%s' --pretty=tformat: --numstat | grep -v '^-' | grep \\\\.%s$ | awk '{ add+=$1; remove+=$2 } END { print \"--- %s ---\\nadded LOC:\", add, \"\\nremoved LOC:\", remove }'", author, l.extension, l.name)

		out, err := exec.Command("bash", "-c", cmd).Output()

		if err != nil {
			return fmt.Sprintf("Failed to execute command: %s", cmd)
		}

		if (isFirst) {
			output += string(out)
			isFirst = false
		} else {
			output += "\n" + string(out)
		}
	}

	return output
}
