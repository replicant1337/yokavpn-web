import React from 'react';
import styles from '@/shared/ui/Dashboard.module.css';
import { useAuth } from '@/features/auth/model/AuthContext';
import type { Subscription } from '@/shared/api.ts';

const DashboardPage: React.FC = () => {
  const { subscription } = useAuth();

  if (!subscription) return null;

  const expiresDate = new Date(subscription.expires_at).toLocaleDateString();

  return (
    <div className={styles.dashboard}>
      <h1 className={styles.title}>личный кабинет</h1>
      
      <div className={styles.grid}>
        <div className={styles.card}>
          <p className={styles.cardTitle}>статус</p>
          <p className={styles.cardValue}>{subscription.status}</p>
        </div>
        
        <div className={styles.card}>
          <p className={styles.cardTitle}>истекает</p>
          <p className={styles.cardValue}>{expiresDate}</p>
        </div>
      </div>
      
      <div className={styles.ctaSection}>
        <div className={styles.subLink}>
          {subscription.remna_sub_link}
        </div>
        <button 
          className={styles.copyBtn}
          onClick={() => {
            navigator.clipboard.writeText(subscription.remna_sub_link);
            alert('ссылка скопирована');
          }}
        >
          скопировать ссылку
        </button>
      </div>
    </div>
  );
};

export default DashboardPage;
