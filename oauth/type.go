package oauth

import "encoding/json"

func UnmarshalGithubUser(data []byte) (GithubUser, error) {
	var r GithubUser
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GithubUser) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// GithubUser type
type GithubUser struct {
	// Login ID
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	// User Name
	Name     string      `json:"name"`
	Company  interface{} `json:"company"`
	Blog     string      `json:"blog"`
	Location interface{} `json:"location"`
	// User EMail
	Email       string      `json:"email"`
	Hireable    interface{} `json:"hireable"`
	Bio         interface{} `json:"bio"`
	PublicRepos int64       `json:"public_repos"`
	PublicGists int64       `json:"public_gists"`
	Followers   int64       `json:"followers"`
	Following   int64       `json:"following"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
	SuspendedAt interface{} `json:"suspended_at"`
}
