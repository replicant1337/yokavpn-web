import React, { useState, useEffect } from 'react';
import Navbar from './components/Navbar';
import Hero from './components/Hero';
import LoginModal from './components/LoginModal';
import Dashboard from './components/Dashboard';
import { RemnaSubscription } from './services/api';

function App() {
  const [isLoginOpen, setIsLoginOpen] = useState(false);
  const [subscription, setSubscription] = useState<RemnaSubscription | null>(null);

  // Simple session persistence
  useEffect(() => {
    const saved = localStorage.getItem('yokavpn_sub');
    if (saved) {
      try {
        setSubscription(JSON.parse(saved));
      } catch (e) {
        localStorage.removeItem('yokavpn_sub');
      }
    }
  }, []);

  const handleLoginSuccess = (sub: RemnaSubscription) => {
    setSubscription(sub);
    setIsLoginOpen(false);
    localStorage.setItem('yokavpn_sub', JSON.stringify(sub));
  };

  const handleLogout = () => {
    setSubscription(null);
    localStorage.removeItem('yokavpn_sub');
  };

  return (
    <div className="app-container">
      <Navbar 
        onLoginClick={() => setIsLoginOpen(true)} 
        isLoggedIn={!!subscription}
        onLogout={handleLogout}
      />
      
      <main>
        {subscription ? (
          <Dashboard subscription={subscription} />
        ) : (
          <Hero />
        )}
      </main>

      {isLoginOpen && (
        <LoginModal 
          onClose={() => setIsLoginOpen(false)} 
          onLoginSuccess={handleLoginSuccess}
        />
      )}
    </div>
  );
}

export default App;
