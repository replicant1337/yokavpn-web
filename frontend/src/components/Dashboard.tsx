import React from 'react';
import styles from './Dashboard.module.css';
import { RemnaSubscription } from '../services/api';

interface DashboardProps {
  subscription: RemnaSubscription;
}

const Dashboard: React.FC<DashboardProps> = ({ subscription }) => {
  const formatBytes = (bytes: number) => {
    if (bytes === 0) return '0 b';
    const k = 1024;
    const sizes = ['b', 'kb', 'mb', 'gb', 'tb'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
  };

  const usedPercentage = (subscription.traffic.used / subscription.traffic.total) * 100;
  const expiresDate = new Date(subscription.expires_at).toLocaleDateString();

  return (
    <div className={styles.dashboard}>
      <h1 className={styles.title}>личный кабинет</h1>
      
      <div className={styles.grid}>
        <div className={styles.card}>
          <p className={styles.cardTitle}>трафик</p>
          <p className={styles.cardValue}>
            {formatBytes(subscription.traffic.used)} / {formatBytes(subscription.traffic.total)}
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
          {subscription.subscription_url}
        </div>
        <button 
          className={styles.copyBtn}
          onClick={() => {
            navigator.clipboard.writeText(subscription.subscription_url);
            alert('ссылка скопирована');
          }}
        >
          скопировать ссылку
        </button>
      </div>
    </div>
  );
};

export default Dashboard;
