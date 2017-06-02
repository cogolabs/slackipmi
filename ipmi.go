package main

import "os/exec"

var argv0 = "ipmitool"

func ipmipower(host, user, pass, action string) ([]byte, error) {
	args := []string{"-H", host, "-U", user, "-P", pass, "power", action}
	cmd := exec.Command(argv0, args...)
	return cmd.CombinedOutput()
}
