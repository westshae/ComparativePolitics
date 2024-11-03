import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const Header = () => {
  const navigate = useNavigate();
  const [loggedIn, setLoggedIn] = useState<boolean>(false);

  useEffect(() => {
    setLoggedIn(localStorage.getItem("token") !== null);
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("name");
    localStorage.removeItem("email");
    setLoggedIn(false);
    navigate("/")
  };

  return (
    <div className="header-container">
      <button className="nav-button" onClick={() => navigate("/")}>
        Home
      </button>
      <div className="nav-group">
        {loggedIn ? (
          <>
            <button className="nav-button" onClick={() => navigate("/populate")}>
              Populate
            </button>
            <button className="nav-button" onClick={() => navigate("/answer")}>
              Answers
            </button>
            <button className="nav-button nav-button-secondary" onClick={() => handleLogout()}>
              Logout
            </button>
          </>
        ) : (
          <>
            <button className="nav-button nav-button-primary" onClick={() => navigate("/authenticate")}>
              Login
            </button>
            <button className="nav-button nav-button-secondary" onClick={() => navigate("/authenticate")}>
              Register
            </button>
          </>
        )}
      </div>
    </div>
  );
};

export default Header;
