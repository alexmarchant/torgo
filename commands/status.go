package commands

func Status(cmd *Command) error {
	if len(cmd.Options) < 1 {
		return BadCommand
	}

	return nil
}
