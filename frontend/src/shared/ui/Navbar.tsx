import React from 'react';
import { useNavigate, Link } from 'react-router-dom';
import styles from '@/shared/ui/Navbar.module.css';

interface NavbarProps {
  isLoggedIn: boolean;
  onLogout: () => void;
}

const Navbar: React.FC<NavbarProps> = ({ isLoggedIn, onLogout }) => {
  const navigate = useNavigate();

  return (
    <nav className={styles.navbar}>
      <div className={styles.container}>
        <div className={styles.logo}>
          <Link to="/" className={styles.logoLink}>
            <span className={styles.logoText}>yokavpn</span>
          </Link>
        </div>

        <div className={styles.navLinks}>
          <Link to="/" className={styles.link}>купить vpn</Link>
          <a href="#" className={styles.link}>возможности</a>
          <a href="#" className={styles.link}>локации</a>
          <a href="#" className={styles.link}>f.a.q.</a>
        </div>

        <div className={styles.actions}>
          <button className={styles.langBtn}>ru</button>
          {isLoggedIn ? (
            <>
              <button className={styles.navBtn} onClick={() => navigate('/dashboard')}>
                кабинет
              </button>
              <button className={styles.authBtn} onClick={onLogout}>
                выйти
              </button>
            </>
          ) : (
            <button className={styles.authBtn} onClick={() => navigate('/login')}>
              войти
            </button>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
