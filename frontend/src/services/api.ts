import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

export interface Traffic {
  total: number;
  used: number;
  remaining: number;
}

export interface RemnaSubscription {
  id: string;
  short_id: string;
  user_id: string;
  subscription_url: string;
  expires_at: string;
  traffic: Traffic;
}

export interface CreateSubscriptionRequest {
  email: string;
  username: string;
}

export interface CreateSubscriptionResponse {
  message: string;
  remna_user_id: string;
  subscription_link: string;
  short_id: string;
}

export const createSubscription = async (data: CreateSubscriptionRequest): Promise<CreateSubscriptionResponse> => {
  const response = await axios.post<CreateSubscriptionResponse>(`${API_BASE_URL}/subscriptions`, data);
  return response.data;
};

export const getSubscription = async (shortId: string): Promise<RemnaSubscription> => {
  const response = await axios.get<RemnaSubscription>(`${API_BASE_URL}/subscriptions/${shortId}`);
  return response.data;
};

export const checkHealth = async () => {
  const response = await axios.get(`${API_BASE_URL}/health`);
  return response.data;
};
