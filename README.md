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

## Building for Production

To create a production build:

```bash
npm run build
```

This will create a `dist` folder with optimized production files.

Preview the production build locally:
```bash
npm run preview
```

## Deployment

### Option 1: Vercel (Recommended)
1. Install Vercel CLI: `npm i -g vercel`
2. Run `vercel` in the project directory
3. Follow the prompts to deploy

### Option 2: Netlify
1. Install Netlify CLI: `npm i -g netlify-cli`
2. Run `netlify deploy --prod`
3. Follow the prompts

### Option 3: GitHub Pages
1. Update `vite.config.js` base path: `base: '/chess-club-app/'`
2. Run `npm run build`
3. Push the `dist` folder to GitHub Pages

### Option 4: Firebase Hosting
1. Install Firebase CLI: `npm i -g firebase-tools`
2. Run `firebase init hosting`
3. Run `firebase deploy`

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

## how to push
- git init
- git add -A
- git commit -m "Initial commit"
- git remote add origin https://github.com/chessclub2205-dev/Chess_Club_app2.git
- git push -u origin main