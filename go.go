//Package try ...
//Ref :https://dzone.com/articles/try-and-catch-in-golang
//Learning the GoLang
package try

type (
	//Exception ...
	Exception interface{}
	//Block ... try catch block
	Block struct {
		Try     func()
		Catch   func(Exception)
		Finally func()
	}
)

//Throw ...
func Throw(up Exception) {
	panic(up)
}

//Do ...
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
