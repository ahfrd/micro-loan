package helpers

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}
