package models

//TestsStruct for ut case
type TestsStruct struct {
	Name    string
	Args    interface{}
	Want    interface{}
	WantErr error
	Mock    func()
}
