import { useNavigate } from "react-router-dom";
import Content from "../components/Content";

const LandingPage = () => {
  const navigate = useNavigate()

  return (
    <Content>
      <div className="page-container">
        <div className="hero-section">
          <h1 className="hero-title">Discover Your Political Landscape</h1>
          <p className="hero-subtitle">
            Answer thought-provoking questions and explore your political stances in a whole new way.
            Our platform will map out your preferences, helping you see how your views align with
            each other and, soon, with political parties and real-world policies.
          </p>
          <button className="call-to-action" onClick={() => navigate("/authenticate")}>Register or Login here</button>
          <button className="call-to-action" onClick={() => navigate("/answer")}>Already logged in? Click here</button>
        </div>

        <div className="features-section">
          <div className="feature-card">
            <h2 className="feature-title">Insightful Questions</h2>
            <p className="feature-description">
              Our carefully crafted questions cover a range of topics to help you uncover subtle nuances
              in your political beliefs.
            </p>
          </div>
          <div className="feature-card">
            <h2 className="feature-title">Personalized Mapping</h2>
            <p className="feature-description">
              See your preferences in a visual format, mapped based on your answers. Understand which issues
              matter most to you.
            </p>
          </div>
          <div className="feature-card">
            <h2 className="feature-title">Comparison to Real Data</h2>
            <p className="feature-description">
              Coming soon: Compare your views with political parties and representatives using voting data,
              policy records, and more.
            </p>
          </div>
          <div className="feature-card">
            <h2 className="feature-title">Dynamic Updates</h2>
            <p className="feature-description">
              Our platform evolves with current events, policy changes, and new insights, keeping your
              political map relevant.
            </p>
          </div>
        </div>
      </div>
    </Content>
  );
};

export default LandingPage;
