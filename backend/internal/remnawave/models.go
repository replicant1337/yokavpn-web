package remnawave

type RemnaUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

type RemnaSubscription struct {
	ID        string `json:"id"`
	ShortID   string `json:"short_id"` // Добавляем ShortID
	UserID    string `json:"user_id"`
	SubLink   string `json:"subscription_url"`
	ExpiresAt string `json:"expires_at"`
	Traffic   struct {
		Total    int64 `json:"total"`
		Used     int64 `json:"used"`
		Remaining int64 `json:"remaining"`
	} `json:"traffic"`
}

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
