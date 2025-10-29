// Top-level application component that composes the main page
import React from 'react'
import HomePage from './components/HomePage'

function App() {
  // Render the HomePage which contains header, content grid, and sidebar
  return (
    <div className="App">
      <HomePage />
    </div>
  )
}

export default App
