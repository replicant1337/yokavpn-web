import React from 'react';
import Navbar from './components/Navbar';
import Hero from './components/Hero';

function App() {
  return (
    <div className="app-container">
      <Navbar />
      <main>
        <Hero />
      </main>
    </div>
  );
}

export default App;
