package shellmd

import (
	"os/exec"
	"strings"
)

func Cmdex(c string, bandera bool, inicial string) {
	if bandera == true {
		splxit := strings.Fields(c)
		_, _ = exec.Command(splxit[0], splxit[1:]...).CombinedOutput()
	} else {
		splxit := strings.Fields(c)
		_, _ = exec.Command(inicial, splxit[0:]...).CombinedOutput()
	}
}
