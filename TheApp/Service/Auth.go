package Service

type AuthInformation struct {
	Email string
	Permission []string
}

func StaticAuthService() []AuthInformation{
	var authlist = []AuthInformation{
		AuthInformation{
			Email:      "hakim@nganu.com",
			Permission: []string{"GET", "POST","PUT","DELETE"},
		},
		AuthInformation{
			Email:      "andreas@nganu.com",
			Permission: []string{"GET", "POST"},
		},
	}
	return authlist


}