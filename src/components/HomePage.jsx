import React from 'react'
import Header from './Header'
import MainContent from './MainContent'
import Sidebar from './Sidebar'
import './HomePage.css'

const HomePage = () => {
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
