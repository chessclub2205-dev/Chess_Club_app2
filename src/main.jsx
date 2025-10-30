// Entry module that bootstraps the React application
import React from 'react'
import ReactDOM from 'react-dom/client'
// Root component of the app
import App from './App.jsx'
// Global styles applied to the entire document
import './index.css'

// Create the React root and render the application inside StrictMode
ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)

// Register a basic service worker for PWA capabilities
if ('serviceWorker' in navigator) {
  window.addEventListener('load', () => {
    navigator.serviceWorker.register('/sw.js').catch(console.error)
  })
}
