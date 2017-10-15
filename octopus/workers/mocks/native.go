package mocks

type Native struct {

}

func (*Native) Request() []byte {
	panic("implement me")
}

func (*Native) IsExtValid() bool {
	panic("implement me")
}

func (*Native) AdLength() int {
	panic("implement me")
}

func (*Native) Attributes() map[string]interface{} {
	panic("implement me")
}
