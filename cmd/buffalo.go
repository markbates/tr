package cmd

func BuffaloBuilder(args []string) *Cmd {
	args = append([]string{"test"}, args...)
	return New("buffalo", args...)
}
