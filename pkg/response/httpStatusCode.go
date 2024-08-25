package response

const (
	ErrCodeSuccess 		= 20001 // Sucess
	ErrCodeParamInvalid = 20003 //Email invalid
)

// message

var msg = map[int]string {
	ErrCodeSuccess: "Sucess",
	ErrCodeParamInvalid: "Email is invalid",
	
}