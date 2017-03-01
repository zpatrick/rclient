package rclient

func Do(s RequestSender, r ResponseReader) error {
	resp, err := s.Send()
	if err != nil {
		return err
	}

	return r.Read(resp)
}
