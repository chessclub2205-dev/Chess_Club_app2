import React from 'react'
import BattlePassIcon from './icons/BattlePassIcon'
import TournamentIcon from './icons/TournamentIcon'
import VsIcon from './icons/VsIcon'
import StoryIcon from './icons/StoryIcon'
import PracticeIcon from './icons/PracticeIcon'
import './MainContent.css'

const MainContent = () => {
  const gameModes = [
    {
      id: 'battle-pass',
      icon: <BattlePassIcon className="chess-icon" size={40} />,
      label: 'Battle pass',
      position: 'top-left'
    },
    {
      id: 'tournaments',
      icon: <TournamentIcon className="chess-icon" size={40} />,
      label: 'Tournaments',
      position: 'top-right'
    },
    {
      id: 'vs',
      icon: <VsIcon className="chess-icon vs-icon" size={50} />,
      label: 'VS',
      position: 'center',
      isMain: true
    },
    {
      id: 'story',
      icon: <StoryIcon className="chess-icon" size={40} />,
      label: 'Story',
      position: 'bottom-left'
    },
    {
      id: 'practice',
      icon: <PracticeIcon className="chess-icon" size={40} />,
      label: 'Practice',
      position: 'bottom-right'
    }
  ]

  return (
    <div className="main-content">
      <div className="game-modes-grid">
        {gameModes.map((mode) => (
          <div 
            key={mode.id} 
            className={`game-mode ${mode.position} ${mode.isMain ? 'main-mode' : ''}`}
          >
            <div className="mode-icon-container">
              {mode.icon}
            </div>
            <div className="mode-label">{mode.label}</div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default MainContent
