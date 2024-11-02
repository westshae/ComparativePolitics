import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const Header = () => {
  const navigate = useNavigate();

  const [loggedIn, setLoggedIn] = useState<boolean>(false);

  useEffect(() => {
    setLoggedIn(localStorage.getItem("token") !== null);
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token")
    localStorage.removeItem("name")
    localStorage.removeItem("email")
    setLoggedIn(false)
  }

  return (
    <div>
      <button onClick={() => navigate("/")}>Home</button>
      {loggedIn &&
        <>
          <button onClick={() => navigate("/populate")}>Populate</button>
          <button onClick={() => navigate("/answer")}>Answers</button>
          <button onClick={() => handleLogout()}>Logout</button>
        </>
      }
      {!loggedIn &&
        <>
          <button onClick={() => navigate("/login")}>Login</button>
          <button onClick={() => navigate("/register")}>Register</button>
        </>
      }
    </div>
  )
}

export default Header