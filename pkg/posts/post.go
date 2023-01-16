package posts

type Post struct {
	Title      string // E.g. title for a reddit/hackernews post
	Body       string // E.g. text post for reddit, tweet text, etc.
	MainLink   string // E.g. reddit link, main link on tweet, etc.
	ImagePaths []string
	VideoPaths []string
	AudioPaths []string
}
