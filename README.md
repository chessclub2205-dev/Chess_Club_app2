# Chess Club App

A modern chess gaming platform with monetization features and customizable game boards.

## Features

- **Dark Theme UI**: Sleek, modern interface with neon-green accents
- **Responsive Design**: Works on both mobile and desktop
- **Game Modes**: Battle Pass, Tournaments, VS, Story, and Practice
- **User Profile**: Profile picture and rank display
- **Sidebar Navigation**: Messages, Shop, Skin customization, and Settings
- **Notification System**: Badge notifications for messages

## Getting Started

### Prerequisites

- Node.js (version 16 or higher)
- npm or yarn

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

4. Open your browser and navigate to `http://localhost:3000`

## Project Structure

```
src/
├── components/
│   ├── HomePage.jsx          # Main page component
│   ├── Header.jsx            # Header with user profile
│   ├── MainContent.jsx       # Game modes grid
│   ├── Sidebar.jsx           # Right sidebar navigation
│   └── *.css                 # Component styles
├── App.jsx                   # Main app component
├── main.jsx                  # Entry point
└── index.css                 # Global styles
```

## Technologies Used

- React 18
- Vite (build tool)
- Lucide React (icons)
- CSS3 (styling)

## Responsive Design

The app is fully responsive and adapts to different screen sizes:
- **Desktop**: Full layout with sidebar
- **Tablet**: Optimized grid layout
- **Mobile**: Stacked layout with bottom navigation

## Future Features

- Chess game implementation
- Coin system (Gold, Silver, Bronze)
- Tournament management
- Battle Pass system
- Custom skins and themes
- Real money gaming features
