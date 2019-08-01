package errno

//TODO: Error Code check out
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
	ErrACLAdd  = &Errno{Code: 20206, Message: "Error ACL Add"}
	ErrACLDel  = &Errno{Code: 20207, Message: "Error ACL Del"}
	ErrACLList = &Errno{Code: 20207, Message: "Fail to get ACL list"}

	//Address Set Error
	ErrASAdd    = &Errno{Code: 20208, Message: "Address Set Add Fail."}
	ErrASDel    = &Errno{Code: 20209, Message: "Address Set Delete Fail."}
	ErrASUpdate = &Errno{Code: 20210, Message: "Address Set Update Fail."}
	ErrASList   = &Errno{Code: 20211, Message: "Get List of Address set Fail"}
	ErrASGet    = &Errno{Code: 20212, Message: "Fail to get Address Set by this name"}

	//Logical Router Error
	ErrLRAdd  = &Errno{Code: 20213, Message: "Add Logical router Fail , already exist or connect fail"}
	ErrLRDel  = &Errno{Code: 20214, Message: "Delete Logical router Fail , object not found or connect fail"}
	ErrLRGet  = &Errno{Code: 20215, Message: "Get Logical router Fail, Object not found or connect fail"}
	ErrLRList = &Errno{Code: 20216, Message: "Fail to get router list, Object not found or connect fail"}

	//LoadBlancer
	ErrLBAdd    = &Errno{Code: 20217, Message: "Add LoadBlancer Fail ,May already Exist or connect fail"}
	ErrLBUpdate = &Errno{Code: 20218, Message: "LoadBlancer Update Fail. Object doesn't exist or connect fail"}
	ErrLBDel    = &Errno{Code: 20219, Message: "LoadBlancer Delete Fail. Object doesn't exist or connect fail"}
	ErrLBList   = &Errno{Code: 20220, Message: "Can't List LoadBlancer. Object doesn't exist or connect fail"}

	ErrLSLBAdd = &Errno{Code: 20221, Message: "Logical Switch add LoadBlancer Fail , some of those doesn't exist or connect fail"}
	ErrLSLBDel = &Errno{Code: 20222, Message: "Logical Switch Delete LoadBlancer Fail , some of those doesn't exist or connect fail"}

	ErrDHCPOptionAdd  = &Errno{Code: 20223, Message: "DHCP OPTION add fail. Already Exist, schema error or connect fail"}
	ErrDHCPOptionDel  = &Errno{Code: 20224, Message: "DHCP OPTION del fail. Object doesn't exist or connect Fail"}
	ErrDHCPOptionSet  = &Errno{Code: 20225, Message: "DHCP OPTION set fail. Object doesn't exist or connect Fail"}
	ErrDHCPOptionList = &Errno{Code: 20226, Message: "Get DHCP OPTION List fail. Object doesn't exist or connect Fail'"}

	ErrDHCPV4Set = &Errno{Code: 20227, Message: "ErrDHCPV4 Set FAILED . Object doesn't exist or connect fail"}
	ErrDHCPV4Get = &Errno{Code: 20228, Message: "ErrDHCPV4 Get FAILED . Object doesn't exist or connect fail"}

	ErrDHCPV6Set = &Errno{Code: 20229, Message: "ErrDHCPV6 Set FAILED . Object doesn't exist or connect fail"}
	ErrDHCPV6Get = &Errno{Code: 20230, Message: "ErrDHCPV6 Get FAILED . Object doesn't exist or connect fail"}

	ErrNatAdd = &Errno{Code: 20231, Message: "NAT Add FAILED , ERROR SCHEMA or connect fail"}
	ErrNatDel = &Errno{Code: 20231, Message: "NAT Delete FAILED , Object not found or connect fail"}
)
