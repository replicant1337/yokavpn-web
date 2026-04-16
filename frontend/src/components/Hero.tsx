import React from 'react';
import styles from './Hero.module.css';

const Hero: React.FC = () => {
  return (
    <div className={styles.hero}>
      <div className={styles.spotlight}></div>

      <div className={styles.content}>
        <h1 className={styles.title}>
          БЫСТРЫЙ И <br /> БЕЗОПАСНЫЙ VPN
        </h1>

        <p className={styles.description}>
          VPN для iPhone, Android, Windows, Linux и macOS. Быстрая сеть, без лимитов, стабильная работа. Подключение за минуту.
        </p>

        <div className={styles.ctaWrapper}>
          <a
            href="#get-started"
            className={styles.ctaButton}
          >
            ПОПРОБОВАТЬ БЕСПЛАТНО
          </a>
        </div>
      </div>
    </div>
  );
};

export default Hero;
