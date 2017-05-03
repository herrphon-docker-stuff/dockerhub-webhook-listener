package api

type HubMessage struct {
	Callback_url string
	Repository   struct {
		Status    string
		RepoUrl   string `json:"repo_url"`
		Owner     string
		IsPrivate bool `json:"is_private"`
		Name      string
		StarCount int    `json:"star_count"`
		RepoName  string `json:"repo_name"`
	}

	Push_data struct {
		PushedAt float64 `json:"pushed_at"`
		Images   []string
		Pusher   string
	}
}
