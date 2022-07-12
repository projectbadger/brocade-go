package test_data

var (
	SessionImplAuthHeader = []testData{
		{
			Name: "Sessionless username:password",
			InputStr: []string{
				"username",
				"password",
				"", //authHeader
				"", // authHeader after login
				"", // authHeader after logout
			},
			InputInt: []int{
				0,   // sessionType sessionless
				200, // login response
				204, // logout response
			},
			ExpectedStr: []string{
				"Basic dXNlcm5hbWU6cGFzc3dvcmQ", // client response headers
				"Basic dXNlcm5hbWU6cGFzc3dvcmQ", // client response headers
				"Basic dXNlcm5hbWU6cGFzc3dvcmQ", // client response headers
			},
			ExpectedErr: []string{
				"",
				"",
			},
			ExpectedBool: []bool{
				true, // Is sessionless
			},
		},
		{
			Name: "Sessionless username:",
			InputStr: []string{
				"username",
				"",
				"", // authHeader
				"", // authHeader after login
				"", // authHeader after logout
			},
			InputInt: []int{
				0,   // sessionType sessionless
				200, // login response
				204, // logout response
			},
			ExpectedStr: []string{
				"Basic dXNlcm5hbWU6", // client response headers
				"Basic dXNlcm5hbWU6", // client response headers
				"Basic dXNlcm5hbWU6", // client response headers
			},
			ExpectedErr: []string{
				"",
				"",
			},
			ExpectedBool: []bool{
				true, // Is sessionless
			},
		},
		{
			Name: "Sessionless :password",
			InputStr: []string{
				"",
				"password",
				"", // authHeader
				"", // authHeader after login
				"", // authHeader after logout
			},
			InputInt: []int{
				0,   // sessionType sessionless
				200, // login response
				204, // logout response
			},
			ExpectedStr: []string{
				"Basic OnBhc3N3b3Jk", // client response headers
				"Basic OnBhc3N3b3Jk", // client response headers
				"Basic OnBhc3N3b3Jk", // client response headers
			},
			ExpectedErr: []string{
				"",
				"",
			},
			ExpectedBool: []bool{
				true, // Is sessionless
			},
		},
		{
			Name: "Sessionless username:password existing login",
			InputStr: []string{
				"username",
				"password",
				"Custom neki", // authHeader
				"",            // authHeader after login
				"",            // authHeader after logout
			},
			InputInt: []int{
				0,   // sessionType sessionless
				200, // login response
				204, // logout response
			},
			ExpectedStr: []string{
				"Basic dXNlcm5hbWU6cGFzc3dvcmQ", // client response headers
				"Basic dXNlcm5hbWU6cGFzc3dvcmQ", // client response headers
				"Basic dXNlcm5hbWU6cGFzc3dvcmQ", // client response headers
			},
			ExpectedErr: []string{
				"",
				"",
			},
			ExpectedBool: []bool{
				true, // Is sessionless
			},
		},
		{
			Name: "Session username:password",
			InputStr: []string{
				"username",
				"password",
				"", // authHeader
				"", // authHeader after login
				"", // authHeader after login
			},
			InputInt: []int{
				1,   // sessionType session
				200, // login response
				204, // logout response
			},
			ExpectedStr: []string{
				"Custom_Basic dXNlcm5hbWU6cGFzc3dvcmQ", // client response headers
				"Custom_Basic dXNlcm5hbWU6cGFzc3dvcmQ", // client response headers
			},
			ExpectedErr: []string{
				"",
				"",
			},
			ExpectedBool: []bool{
				false, // Is sessionless
			},
		},
		{
			Name: "Session username:password existing login",
			InputStr: []string{
				"username",
				"password",
				"Custom neki",                          // authHeader
				"Custom_Basic dXNlcm5hbWU6cGFzc3dvcmQ", // authHeader after logout
				"",                                     // authHeader after logout
			},
			InputInt: []int{
				1,   // sessionType session
				200, // login response
				204, // logout response
			},
			ExpectedStr: []string{
				"Custom neki", // client response headers
				"Custom neki", // client response headers
				"Custom_Basic dXNlcm5hbWU6cGFzc3dvcmQ",
			},
			ExpectedErr: []string{
				"",
				"",
			},
			ExpectedBool: []bool{
				false, // Is sessionless
			},
		},
		{
			Name: "Session username:password failed login",
			InputStr: []string{
				"username",
				"password",
				"",
				"", // authHeader after login
				"", // authHeader after logout
			},
			InputInt: []int{
				1,   // sessionType session
				401, // login response
				401, // logout response
			},
			ExpectedStr: []string{
				"Custom_Basic dXNlcm5hbWU6cGFzc3dvcmQ", // auth header before login
				"",                                     // auth header after login
			},
			ExpectedErr: []string{
				"login error 401",
				"logout error 401",
			},
			ExpectedBool: []bool{
				false, // Is sessionless
			},
		},
	}
)
