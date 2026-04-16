import React from 'react';
import styles from './Navbar.module.css';

interface NavbarProps {
  onLoginClick: () => void;
  isLoggedIn: boolean;
  onLogout: () => void;
}

const Navbar: React.FC<NavbarProps> = ({ onLoginClick, isLoggedIn, onLogout }) => {
  return (
    <nav className={styles.navbar}>
      <div className={styles.container}>
        <div className={styles.logo}>
          <span className={styles.logoText}>yokavpn</span>
        </div>

        <div className={styles.navLinks}>
          <a href="#" className={styles.link}>купить vpn</a>
          <a href="#" className={styles.link}>возможности</a>
          <a href="#" className={styles.link}>локации</a>
          <a href="#" className={styles.link}>f.a.q.</a>
        </div>

        <div className={styles.actions}>
          <button className={styles.langBtn}>ru</button>
          {isLoggedIn ? (
            <button className={styles.authBtn} onClick={onLogout}>
              выйти
            </button>
          ) : (
            <button className={styles.authBtn} onClick={onLoginClick}>
              войти
            </button>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
