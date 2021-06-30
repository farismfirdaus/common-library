package recover

//Exception for containing exception
type Exception interface{}

// OnError is handler function when error occured
type OnError func(Exception)

// Catch error
func Catch(onerror OnError) {
	if r := recover(); r != nil {
		onerror(r)
	}
}
