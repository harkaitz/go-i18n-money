package i18nm

func (c *Cents) MarshalCSV() (string, error) {
	return c.String(), nil
}

func (c *Cents) UnmarshalCSV(csv string) (err error) {
	*c, _, err = ParseMoney(csv)
	return
}
