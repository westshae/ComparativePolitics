import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import PopulatePage from './pages/PopulatePage';
import AnswerPage from './pages/AnswerPage';
import AuthPage from './pages/AuthPage';

const App = () => {
  return (
    <BrowserRouter >
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/authenticate" element={<AuthPage />} />
        <Route path="/populate" element={<PopulatePage />} />
        <Route path="/answer" element={<AnswerPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
