// Top-level application component that composes the main page
import React, { useEffect, useState } from 'react'
import HomePage from './components/HomePage'

function App() {
  // Logic to enforce landscape mode on mobile (<=768px) only
  const [showOverlay, setShowOverlay] = useState(false)
  useEffect(() => {
    const checkOrientation = () => {
      setShowOverlay(window.innerWidth <= 768 && window.innerHeight > window.innerWidth)
    }
    checkOrientation()
    window.addEventListener('resize', checkOrientation)
    window.addEventListener('orientationchange', checkOrientation)
    return () => {
      window.removeEventListener('resize', checkOrientation)
      window.removeEventListener('orientationchange', checkOrientation)
    }
  }, [])
  useEffect(() => {
    if (showOverlay) document.body.style.overflow = 'hidden'
    else document.body.style.overflow = ''
  }, [showOverlay])

  return (
    <>
      {showOverlay && (
        <div className="orientation-overlay">
          Please rotate your device to landscape mode to use this app.
        </div>
      )}
      <div className="App">
        <HomePage />
      </div>
    </>
  )
}

export default App
