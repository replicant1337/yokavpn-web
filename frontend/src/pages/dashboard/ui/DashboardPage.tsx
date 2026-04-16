import React from 'react';
import styles from '@/shared/ui/Dashboard.module.css';
import { useAuth } from '@/features/auth/model/AuthContext';
import type { Subscription } from '@/shared/api.ts';

const DashboardPage: React.FC = () => {
  const { subscription } = useAuth();

  if (!subscription) return null;

  const formatBytes = (bytes: number) => {
    if (bytes === 0) return '0 b';
    const k = 1024;
    const sizes = ['b', 'kb', 'mb', 'gb', 'tb'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
  };

  const usedPercentage = (subscription.traffic_used / subscription.traffic_total) * 100;
  const expiresDate = new Date(subscription.expires_at).toLocaleDateString();

  return (
    <div className={styles.dashboard}>
      <h1 className={styles.title}>личный кабинет</h1>
      
      <div className={styles.grid}>
        <div className={styles.card}>
          <p className={styles.cardTitle}>трафик</p>
          <p className={styles.cardValue}>
            {formatBytes(subscription.traffic_used)} / {formatBytes(subscription.traffic_total)}
          </p>
          <div className={styles.progressBar}>
            <div 
              className={styles.progressFill} 
              style={{ width: `${Math.min(usedPercentage, 100)}%` }}
            ></div>
          </div>
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
