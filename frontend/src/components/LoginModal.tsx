import React, { useState } from 'react';
import styles from './LoginModal.module.css';
import { getSubscription, RemnaSubscription } from '../services/api';

interface LoginModalProps {
  onClose: () => void;
  onLoginSuccess: (sub: RemnaSubscription) => void;
}

const LoginModal: React.FC<LoginModalProps> = ({ onClose, onLoginSuccess }) => {
  const [shortId, setShortId] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleLogin = async () => {
    if (!shortId) return;
    
    setLoading(true);
    setError('');
    
    try {
      // Extract shortId if full URL was pasted
      const id = shortId.includes('/') ? shortId.split('/').pop() || '' : shortId;
      const sub = await getSubscription(id);
      onLoginSuccess(sub);
    } catch (err) {
      setError('subscription not found or server error');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className={styles.overlay} onClick={onClose}>
      <div className={styles.modal} onClick={(e) => e.stopPropagation()}>
        <button className={styles.closeBtn} onClick={onClose}>×</button>
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
          disabled={loading || !shortId}
        >
          {loading ? 'загрузка...' : 'войти'}
        </button>
        
        {error && <p className={styles.error}>{error}</p>}
      </div>
    </div>
  );
};

export default LoginModal;
