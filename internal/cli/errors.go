package cli

func handleErr(err error) error {
	if err == nil {
		return nil
	}

	return err
}
