package passwordgen

/*
GenerateRequest Defines something
*/
type GenerateRequest struct {
	AddSymbol       bool `json:"addSymbol"`
	LeetMode        bool `json:"leetMode"`
	NumPasswords    int  `json:"numPasswords"`
	RandomUpperCase bool `json:"randomUpperCase"`
}
