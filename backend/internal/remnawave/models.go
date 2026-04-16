package remnawave

type RemnaUser struct {
	ShortUuid           string `json:"shortUuid"`
	Username            string `json:"username"`
	ExpiresAt           string `json:"expiresAt"`
	IsActive            bool   `json:"isActive"`
	UserStatus          string `json:"userStatus"`
	TrafficUsedBytes    int64  `json:"trafficUsedBytes"`
	TrafficLimitBytes   int64  `json:"trafficLimitBytes"`
}

type RemnaSubscription struct {
	User            RemnaUser `json:"user"`
	Links           []string  `json:"links"`
	SubscriptionUrl string    `json:"subscriptionUrl"`
}

type SubscriptionsResponse struct {
	Response struct {
		Subscriptions []RemnaSubscription `json:"subscriptions"`
	} `json:"response"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateSubscriptionRequest struct {
	UserID string `json:"user_id"`
}
