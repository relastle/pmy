package anypm

type sourceGetter interface {
	getSource(
		bufferLeft string,
		bufferRight string,
		regexpLeft string,
		regexpRight string,
	) []string
}

type sourceGetterMock struct {
}
