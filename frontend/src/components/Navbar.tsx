import React from 'react';
import styles from './Navbar.module.css';

const Navbar: React.FC = () => {
  return (
    <nav className={styles.navbar}>
      <div className={styles.container}>
        <div className={styles.logo}>
          <span className={styles.logoText}>YokaVPN</span>
        </div>

        <div className={styles.navLinks}>
          <a href="#" className={styles.link}>КУПИТЬ VPN</a>
          <a href="#" className={styles.link}>ВОЗМОЖНОСТИ</a>
          <a href="#" className={styles.link}>ЛОКАЦИИ</a>
          <a href="#" className={styles.link}>F.A.Q.</a>
        </div>

        <div className={styles.actions}>
          <button className={styles.langBtn}>RU</button>
          <button className={styles.authBtn}>
            ВОЙТИ
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
