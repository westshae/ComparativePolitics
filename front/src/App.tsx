import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import PopulatePage from './pages/PopulatePage';
import AnswerPage from './pages/AnswerPage';
import Header from './components/Header';
import Footer from './components/Footer';
import Content from './components/Content';

const App = () => {
  return (

    <BrowserRouter>
      <Header />
      <Content>
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/populate" element={<PopulatePage />} />
          <Route path="/answer" element={<AnswerPage />} />
        </Routes>
      </Content>
      <Footer />

    </BrowserRouter>
  );
}

export default App;
