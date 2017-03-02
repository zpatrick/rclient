package rclient

func Do(send Sender, read Reader) error {
	resp, err := send()
	if err != nil {
		return err
	}

	return read(resp)
}
