import { useNavigate } from "react-router-dom";

const Footer = () => {
  const navigate = useNavigate();


  return (
    <div>
      <button onClick={() => navigate("/")}>Home</button>
      <button onClick={() => navigate("/login")}>Login</button>
      <button onClick={() => navigate("/register")}>Register</button>
      <button onClick={() => navigate("/populate")}>Populate</button>
      <button onClick={() => navigate("/answer")}>Answers</button>
    </div>
  )
}

export default Footer