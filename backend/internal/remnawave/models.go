package remnawave

type RemnaUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}

type RemnaSubscription struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	SubLink   string `json:"subscription_url"`
	ExpiresAt string `json:"expires_at"`
}

// Response models for the Remnawave API wrapper
type UserResponse struct {
	Data RemnaUser `json:"data"`
}

type SubscriptionResponse struct {
	Data RemnaSubscription `json:"data"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateSubscriptionRequest struct {
	UserID string `json:"user_id"`
}
