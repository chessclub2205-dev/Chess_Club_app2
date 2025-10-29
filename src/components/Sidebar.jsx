// Right-side sidebar listing secondary actions with a pawn SVG
import React from 'react'
import './Sidebar.css'

// Decorative pawn icon used for each sidebar item
const PawnIcon = () => (
  <svg width="121" height="83" viewBox="0 0 121 83" fill="none" xmlns="http://www.w3.org/2000/svg">
    <g filter="url(#filter0_d_2192_130)">
      <path d="M7.51929 35.2143C8.73508 24.7679 16.0907 18.6428 24.0907 18.6428C32.0907 18.6428 33.7302 20.9651 37.8049 24.9286H38.9478L38.9478 20.9286C39.6614 18.979 40.0907 18.6428 42.3764 18.6428C44.6621 18.6428 45.4524 18.5229 45.8049 20.9286V22.6428H50.9478V24.9286H61.8049C76.9144 23.9109 83.5295 20.0987 91.5191 7.21427H95.5191C96.4046 3.53765 97.8048 2.07142 101.233 2.07142C104.662 2.07142 105.015 2.08567 106.376 3.7857C107.42 2.1574 107.519 1.49999 109.805 1.49999C112.091 1.49999 112.471 1.67324 112.662 3.7857L112.662 34.0714" stroke="#DAF1DE" stroke-width="3"/>
      <path d="M7.51929 35.7857C8.73508 46.2321 16.0907 52.3571 24.0907 52.3571C32.0907 52.3571 33.7302 50.0348 37.8049 46.0714H38.9478L38.9478 50.0714C39.6614 52.0209 40.0907 52.3571 42.3764 52.3571C44.6621 52.3571 45.4524 52.477 45.8049 50.0714V48.3571H50.9478V46.0714H61.8049C76.9144 47.0891 83.5295 50.9013 91.5191 63.7857H95.5191C96.4046 67.4623 97.8048 68.9286 101.233 68.9286C104.662 68.9286 105.015 68.9143 106.376 67.2143C107.42 68.8426 107.519 69.5 109.805 69.5C112.091 69.5 112.471 69.3267 112.662 67.2143L112.662 36.9286" stroke="#DAF1DE" stroke-width="3"/>
      <path d="M112.662 36.9285V34.0714" stroke="#DAF1DE" stroke-width="3"/>
      <path d="M7.51931 35.7857C7.49416 35.5619 7.49297 35.4368 7.51931 35.2143" stroke="#DAF1DE" stroke-width="3"/>
    </g>
    <defs>
      <filter id="filter0_d_2192_130" x="0" y="0" width="120.162" height="83" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
        <feFlood flood-opacity="0" result="BackgroundImageFix"/>
        <feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
        <feOffset dy="6"/>
        <feGaussianBlur stdDeviation="3"/>
        <feComposite in2="hardAlpha" operator="out"/>
        <feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0"/>
        <feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_2192_130"/>
        <feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_2192_130" result="shape"/>
      </filter>
    </defs>
  </svg>
)

const Sidebar = () => {
  // Sidebar menu configuration with optional badges
  const sidebarItems = [
    {
      id: 'messages',
      label: 'Messages',
      badge: 12
    },
    {
      id: 'shop',
      label: 'Shop'
    },
    {
      id: 'skin',
      label: 'Skin'
    },
    {
      id: 'settings',
      label: 'Settings'
    }
  ]

  return (
    <div className="sidebar">
      <div className="sidebar-content">
        {/* Render each sidebar row with an icon and label */}
        {sidebarItems.map((item) => (
          <div key={item.id} className="sidebar-item">
            <div className="sidebar-pawn-container">
              <PawnIcon />
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