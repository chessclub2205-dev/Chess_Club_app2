import React from 'react'
import './Sidebar.css'

const Sidebar = () => {
  const sidebarItems = [
    {
      id: 'messages',
      icon: '💬',
      label: 'Messages',
      badge: 12
    },
    {
      id: 'shop',
      icon: '🛒',
      label: 'Shop'
    },
    {
      id: 'skin',
      icon: '🎨',
      label: 'Skin'
    },
    {
      id: 'settings',
      icon: '⚙️',
      label: 'Settings'
    }
  ]

  return (
    <div className="sidebar">
      <div className="sidebar-content">
        {sidebarItems.map((item) => (
          <div key={item.id} className="sidebar-item">
            <div className="sidebar-icon-container">
              <div className="sidebar-icon">{item.icon}</div>
              {item.badge && (
                <div className="notification-badge">
                  {item.badge}
                </div>
              )}
            </div>
            <div className="sidebar-label">{item.label}</div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default Sidebar
