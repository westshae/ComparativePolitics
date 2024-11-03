import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const navigate = useNavigate();

  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const onEmailChange = (email: string) => {
    setEmail(email);
  }

  const onPasswordChange = (password: string) => {
    setPassword(password);
  }

  const handleLogin = async () => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    const formData = new FormData();
    formData.append("email", email);
    formData.append("password", password);

    if (backend_url) {
      axios.post(backend_url + "/login", formData)
        .then((response) => {
          localStorage.setItem("token", response.data.token)
          localStorage.setItem("email", response.data.email)
          localStorage.setItem("name", response.data.name)
          navigate("/")
        })
        .catch((error) => {
          console.log(error)
        })
    }
  }

  return (
    <div>
      <div className="card w-full">
        <h1 className="header">Login below</h1>
        <label className="label" htmlFor="email">Email:</label>
        <input className="textinput" id="email" placeholder="Enter email here..." onChange={(event) => onEmailChange(event.target.value)}></input>

        <label className="label" htmlFor="password">Password:</label>
        <input className="textinput" id="password" type="password" placeholder="Enter password here..." onChange={(event) => onPasswordChange(event.target.value)}></input>

        <button className="button" disabled={email === "" || password === ""} onClick={() => handleLogin()}>Login</button>
      </div>
    </div>
  );
}

export default Login;
