import axios from 'axios';

export interface Traffic {
  total: number;
  used: number;
  remaining: number;
}

export interface Subscription {
  id: number;
  user_id: number;
  plan_name: string;
  remna_user_id: string;
  remna_sub_link: string;
  short_id: string;
  auth_key: string;
  traffic_total: number;
  traffic_used: number;
  expires_at: string;
  status: string;
}

export interface CreateSubscriptionRequest {
  email: string;
  username: string;
}

export interface CreateSubscriptionResponse {
  message: string;
  auth_key: string;
  short_id: string;
  subscription_link: string;
}

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
});

export const subscriptionApi = {
  create: async (data: CreateSubscriptionRequest) => {
    const response = await api.post<CreateSubscriptionResponse>('/subscriptions', data);
    return response.data;
  },
  getById: async (key: string) => {
    const response = await api.get<Subscription>(`/subscriptions/auth/${key}`);
    return response.data;
  },
  checkHealth: async () => {
    const response = await api.get('/health');
    return response.data;
  }
};
