package pmy

type commandGetter interface {
	// GetCommand returns the command to be executed
	// given the
	GetCommand(
		bufferLeft string,
		bufferRight string,
	) string
}

type commandGetterMock struct {
}

func (cg *commandGetterMock) GetCommand(
	bufferLeft string,
	bufferRight string,
) string {
	return "ls"
}

type commandGetterImpl struct {
}

func (cg *commandGetterImpl) GetCommand(
	bufferLeft string,
	bufferRight string,
) string {
	return ""
}
