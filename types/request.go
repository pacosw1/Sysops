package types

//Request used to send communicate commands and their arguments between components
type Request struct {
	Type    string   //Command Type
	Args    []int    //Arguments for that command
	Message []string //Message if avaliable
}

//NewRequest Initializes a new Request
func NewRequest(commandType string, args []int, message []string) *Request {
	return &Request{
		Type:    commandType,
		Args:    args,
		Message: message,
	}
}
