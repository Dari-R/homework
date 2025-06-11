package main

type Bin struct{
	id string
	private bool
	createdAt string //дата и время
	name string
}
type BinList struct{
	Bin
}
func (Bin) createBin() Bin{
	return Bin{
		id: "",
		private: true,
		createdAt: "",
		name: "",
	}
}
func main(){

}