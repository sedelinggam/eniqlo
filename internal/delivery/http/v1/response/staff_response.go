package response

type UserAccessToken struct {
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
