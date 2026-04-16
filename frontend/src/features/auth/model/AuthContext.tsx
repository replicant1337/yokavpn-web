import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import type { Subscription } from '@/shared/api.ts';

interface AuthContextType {
  subscription: Subscription | null;
  isAuthenticated: boolean;
  login: (sub: Subscription) => void;
  logout: () => void;
  isLoading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [subscription, setSubscription] = useState<Subscription | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const saved = localStorage.getItem('yokavpn_sub');
    if (saved) {
      try {
        setSubscription(JSON.parse(saved));
      } catch (e) {
        localStorage.removeItem('yokavpn_sub');
      }
    }
    setIsLoading(false);
  }, []);

  const login = (sub: Subscription) => {
    setSubscription(sub);
    localStorage.setItem('yokavpn_sub', JSON.stringify(sub));
  };

  const logout = () => {
    setSubscription(null);
    localStorage.removeItem('yokavpn_sub');
  };

  return (
    <AuthContext.Provider value={{ 
      subscription, 
      isAuthenticated: !!subscription, 
      login, 
      logout,
      isLoading 
    }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
