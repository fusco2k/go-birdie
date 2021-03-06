package models

//StatusUpdate struct to send a status update
type StatusUpdate struct {
	Status             string   `url:"status,omitempty"`
	InReplyToStatusID  int64    `url:"in_reply_to_status_id,omitempty"`
	PossiblySensitive  *bool    `url:"possibly_sensitive,omitempty"`
	Lat                *float64 `url:"lat,omitempty"`
	Long               *float64 `url:"long,omitempty"`
	PlaceID            string   `url:"place_id,omitempty"`
	DisplayCoordinates *bool    `url:"display_coordinates,omitempty"`
	TrimUser           *bool    `url:"trim_user,omitempty"`
	MediaIds           []int64  `url:"media_ids,omitempty,comma"`
	TweetMode          string   `url:"tweet_mode,omitempty"`
}

//StatusResponse struct to parse the response from the status update
type StatusResponse struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	IDStr     string `json:"id_str"`
	Text      string `json:"text"`
	Truncated bool   `json:"truncated"`
	Entities  struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []interface{} `json:"user_mentions"`
		Urls         []interface{} `json:"urls"`
	} `json:"entities"`
	Source               string      `json:"source"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	User                 struct {
		ID          int    `json:"id"`
		IDStr       string `json:"id_str"`
		Name        string `json:"name"`
		ScreenName  string `json:"screen_name"`
		Location    string `json:"location"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Entities    struct {
			URL struct {
				Urls []struct {
					URL         string `json:"url"`
					ExpandedURL string `json:"expanded_url"`
					DisplayURL  string `json:"display_url"`
					Indices     []int  `json:"indices"`
				} `json:"urls"`
			} `json:"url"`
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		Protected                      bool        `json:"protected"`
		FollowersCount                 int         `json:"followers_count"`
		FriendsCount                   int         `json:"friends_count"`
		ListedCount                    int         `json:"listed_count"`
		CreatedAt                      string      `json:"created_at"`
		FavouritesCount                int         `json:"favourites_count"`
		UtcOffset                      interface{} `json:"utc_offset"`
		TimeZone                       interface{} `json:"time_zone"`
		GeoEnabled                     bool        `json:"geo_enabled"`
		Verified                       bool        `json:"verified"`
		StatusesCount                  int         `json:"statuses_count"`
		Lang                           string      `json:"lang"`
		ContributorsEnabled            bool        `json:"contributors_enabled"`
		IsTranslator                   bool        `json:"is_translator"`
		IsTranslationEnabled           bool        `json:"is_translation_enabled"`
		ProfileBackgroundColor         string      `json:"profile_background_color"`
		ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool        `json:"profile_background_tile"`
		ProfileImageURL                string      `json:"profile_image_url"`
		ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
		ProfileBannerURL               string      `json:"profile_banner_url"`
		ProfileLinkColor               string      `json:"profile_link_color"`
		ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string      `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
		HasExtendedProfile             bool        `json:"has_extended_profile"`
		DefaultProfile                 bool        `json:"default_profile"`
		DefaultProfileImage            bool        `json:"default_profile_image"`
		Following                      bool        `json:"following"`
		FollowRequestSent              bool        `json:"follow_request_sent"`
		Notifications                  bool        `json:"notifications"`
		TranslatorType                 string      `json:"translator_type"`
	} `json:"user"`
	Geo           interface{} `json:"geo"`
	Coordinates   interface{} `json:"coordinates"`
	Place         interface{} `json:"place"`
	Contributors  interface{} `json:"contributors"`
	IsQuoteStatus bool        `json:"is_quote_status"`
	RetweetCount  int         `json:"retweet_count"`
	FavoriteCount int         `json:"favorite_count"`
	Favorited     bool        `json:"favorited"`
	Retweeted     bool        `json:"retweeted"`
	Lang          string      `json:"lang"`
}
