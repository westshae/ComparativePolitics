import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

const Register = () => {
  const navigate = useNavigate();

  const [email, setEmail] = useState<string>("");
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [passwordCheck, setPasswordCheck] = useState<string>("");

  const [showEmailConfirmation, setShowEmailConfirmation] = useState<boolean>(false);
  const [showError, setShowError] = useState<boolean>(false);
  const [registerationError, setRegisterationError] = useState<string>("");

  const onEmailChange = (email: string) => {
    setEmail(email);
  }

  const onUsernameChange = (username: string) => {
    setUsername(username);
  }

  const onPasswordChange = (password: string) => {
    setPassword(password);
  }

  const onPasswordCheckChange = (passwordCheck: string) => {
    setPasswordCheck(passwordCheck);
  }

  const handleRegister = async () => {
    setShowEmailConfirmation(false)
    setShowError(false)

    const backend_url = import.meta.env.VITE_BACKEND_URL;
    const formData = new FormData();
    formData.append("email", email);
    formData.append("password", password);
    formData.append("name", username);

    if (backend_url) {
      axios.post(backend_url + "/register", formData).then((_response) => {
        setShowEmailConfirmation(true);
      })
        .catch((error) => {
          setShowError(true);
          setRegisterationError(error.response.data.error);
          console.log(error)
          console.log(error.response.data.error)
        })
    }
  }

  return (
    <div className="card w-full">
      <h1 className="header">Register</h1>
      <label className="label" htmlFor="email">Email:</label>
      <input className="textinput" id="email" placeholder="Enter email here..." onChange={(event) => onEmailChange(event.target.value)}></input>

      <label className="label" htmlFor="username">Username:</label>
      <input className="textinput" id="username" placeholder="Enter username here..." onChange={(event) => onUsernameChange(event.target.value)}></input>

      <label className="label" htmlFor="password">Password:</label>
      <input className="textinput" id="password" type="password" placeholder="Enter password here..." onChange={(event) => onPasswordChange(event.target.value)}></input>

      <label className="label" htmlFor="passwordCheck">Confirm:</label>
      <input className="textinput" id="passwordCheck" type="password" placeholder="Confirm password here..." onChange={(event) => onPasswordCheckChange(event.target.value)}></input>

      <button className="button" disabled={email === "" || username === "" || password === "" || password !== passwordCheck} onClick={() => handleRegister()}>Register</button>

      <div style={{ display: showEmailConfirmation ? 'block' : 'none' }}>
        <h1>Registeration requires confirmation.</h1>
        <p>Check your inbox at {email} for a confirmation email.</p>
        <button onClick={() => navigate("/login")}>Login</button>
      </div>

      <div style={{ display: showError ? 'block' : 'none' }}>
        <h1>Oops, looks like you ran into an error when registering</h1>
        <p>{registerationError}</p>
      </div>
    </div>
  );
}

export default Register;
