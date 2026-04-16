import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

export interface CreateSubscriptionRequest {
  email: string;
  username: string;
}

export interface CreateSubscriptionResponse {
  message: string;
  remna_user_id: string;
  subscription_link: string;
}

export const createSubscription = async (data: CreateSubscriptionRequest): Promise<CreateSubscriptionResponse> => {
  const response = await axios.post<CreateSubscriptionResponse>(`${API_BASE_URL}/subscriptions`, data);
  return response.data;
};

export const checkHealth = async () => {
  const response = await axios.get(`${API_BASE_URL}/health`);
  return response.data;
};
