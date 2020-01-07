package userdb

// db act as a dummy package level database.
var db map[string]bool

// init initialize a dummy db with some data
func init() {
	db = make(map[string]bool)
	db["ankuranand@dummy.org"] = true
	db["anand@example.com"] = true
}

// UserExist check if the User is registered with the provided email.
func UserExist(email string) bool {
	if _, ok := db[email]; !ok {
		return false
	}
	return true
}
