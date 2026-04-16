import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useMutation } from '@tanstack/react-query';
import styles from '@/shared/ui/LoginModal.module.css';
import { subscriptionApi } from '@/shared/api';
import { useAuth } from '@/features/auth/model/AuthContext';

const LoginPage: React.FC = () => {
  const [shortId, setShortId] = useState('');
  const [error, setError] = useState('');
  const { login } = useAuth();
  const navigate = useNavigate();

  const mutation = useMutation({
    mutationFn: (key: string) => subscriptionApi.getById(key),
    onSuccess: (data) => {
      login(data);
      navigate('/dashboard');
    },
    onError: () => {
      setError('subscription not found or server error');
    }
  });

  const handleLogin = async () => {
    if (!shortId) return;
    setError('');
    const id = shortId.includes('/') ? shortId.split('/').pop() || '' : shortId;
    mutation.mutate(id);
  };

  return (
    <div className={styles.pageContainer}>
      <div className={styles.modal}>
        <h2 className={styles.title}>вход в кабинет</h2>
        <p className={styles.subtitle}>введите ваш shortid или ссылку на подписку</p>
        
        <div className={styles.inputGroup}>
          <input 
            type="text" 
            className={styles.input} 
            placeholder="p-aphszhr53tpjtk..."
            value={shortId}
            onChange={(e) => setShortId(e.target.value)}
            onKeyDown={(e) => e.key === 'Enter' && handleLogin()}
          />
        </div>
        
        <button 
          className={styles.loginBtn} 
          onClick={handleLogin}
          disabled={mutation.isPending || !shortId}
        >
          {mutation.isPending ? 'загрузка...' : 'войти'}
        </button>
        
        {error && <p className={styles.error}>{error}</p>}
      </div>
    </div>
  );
};

export default LoginPage;
