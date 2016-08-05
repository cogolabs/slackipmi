package main

import "os/exec"

func ipmipower(host, user, pass, action string) ([]byte, error) {
	args := []string{"-H", host, "-U", user, "-P", pass, "power", action}
	cmd := exec.Command("ipmitool", args...)
	return cmd.CombinedOutput()
}
