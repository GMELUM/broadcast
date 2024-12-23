package cfg

func ConfigWriting() error {

	err := token()
	if err != nil {
		return err
	}

	err = rps()
	if err != nil {
		return err
	}

	return nil

}
