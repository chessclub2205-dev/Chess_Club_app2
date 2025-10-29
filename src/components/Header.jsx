// App header showing user profile and utility icons
import React from 'react'
import './Header.css'

const Header = () => {
  return (
    <header className="header">
      <div className="header-content">
        {/* Left: user avatar and basic info */}
        <div className="user-profile">
          <div className="profile-picture">
            <img 
              src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=100&h=100&fit=crop&crop=face" 
              alt="User Profile" 
            />
          </div>
          <div className="user-info">
            <div className="username">#username</div>
            <div className="rank">#rank/elo</div>
          </div>
        </div>
        
        {/* Right: placeholder utility icons */}
        <div className="utility-icons">
          <div className="icon-container">
            <div className="utility-icon">⚙</div>
          </div>
          <div className="icon-container">
            <div className="utility-icon">⚙</div>
          </div>
          <div className="icon-container">
            <div className="utility-icon">⚙</div>
          </div>
        </div>
      </div>
      <div className="header-divider"></div>
    </header>
  )
}

export default Header
