package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

	//ovn errors
	ErrLsGet        = &Errno{Code: 20201, Message: "Error Logical Switch Get"}
	ErrLsAdd        = &Errno{Code: 20202, Message: "Error Logical Switch Add"}
	ErrLsDel        = &Errno{Code: 20203, Message: "Error Logical Switch Delete"}
	ErrLsListGet    = &Errno{Code: 20204, Message: "Fail to get list"}
	ErrLsExidOprate = &Errno{Code: 20205, Message: "Error Logical Switch Extend id"}

	//if case ovn api error , the code will fix to 20200
	ErrACLAdd		= &Errno{Code:20206,Message: "Error ACL Add"}
	ErrACLDel		= &Errno{Code:20207,Message: "Error ACL Del"}
	ErrACLList		= &Errno{Code:20207,Message: "Fail to get ACL list"}

	//Address Set Error
	ErrASAdd		= &Errno{Code:20208,Message:"Address Set Add Fail."}
	ErrASDel		= &Errno{Code:20209,Message:"Address Set Delete Fail."}
	ErrASUpdate		= &Errno{Code:20210,Message:"Address Set Update Fail."}
	ErrASList		= &Errno{Code:20211,Message:"Get List of Address set Fail"}
	ErrASGet		= &Errno{Code:20212,Message:"Fail to get Address Set by this name"}
)
