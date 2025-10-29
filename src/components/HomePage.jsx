// Home page layout combining header, main content grid, and sidebar
import React from 'react'
import Header from './Header'
import MainContent from './MainContent'
import Sidebar from './Sidebar'
import './HomePage.css'

const HomePage = () => {
  // Two-column layout: main game mode grid + right sidebar
  return (
    <div className="home-page">
      <Header />
      <div className="main-layout">
        <MainContent />
        <Sidebar />
      </div>
    </div>
  )
}

export default HomePage
