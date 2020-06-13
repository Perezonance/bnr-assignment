package models

type (
	User struct {
		Id			float64		`json:"id"`
		Name 		string		`json:"name"`
		Email 		string		`json:"email"`
		Expertise 	string		`json:"expertise"`
	}

	Post struct {
		Id 			float64		`json:"id"`
		UserId 		float64 	`json:"userId"`
		Title 		string 		`json:"title"`
		Body 		string 		`json:"body"`
	}
	//Used for Input validation of the request body
	RequestUsersPayload struct {
		Payload 	[]User		`json:"users"`
	}

	RequestPostsPayload struct {
		Payload 	[]Post	 	`json:"posts"`
	}

	RequestUserPayload struct {
		User 		User 		`json:"user"`
	}

	RequestPostPayload struct {
		Post 		Post 		`json:"post"`
	}

	RequestUserById struct {
		Id 			float64		`json:"id"`
	}

	RequestUsersById struct {
		Users 		[]float64 	`json:"users"`
	}

	RequestPostById struct {
		Id 			float64 	`json:"id"`
	}

	RequestPostsById struct {
		Posts 		[]float64 	`json:"posts"`
	}
)


