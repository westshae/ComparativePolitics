import axios from "axios";
import { useState } from "react";

const RegisterPage = () => {
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
    console.log(email !== "" && username !== "" && password !== "" && password == passwordCheck)
  }

  const handleRegister = async () => {
    setShowEmailConfirmation(false)
    setShowError(false)

    const backend_url = import.meta.env.VITE_BACKEND_URL;
    const formData = new FormData();
    formData.append("email", email);
    formData.append("password", password);
    formData.append("name", username);

    if(backend_url){
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
    <>
      <h1>Register</h1>
      <label htmlFor="email">Email:</label>
      <input id="email" placeholder="Enter email here..." onChange={(event) => onEmailChange(event.target.value)}></input>

      <label htmlFor="username">Username:</label>
      <input id="username" placeholder="Enter username here..." onChange={(event) => onUsernameChange(event.target.value)}></input>

      <label htmlFor="password">Password:</label>
      <input id="password" placeholder="Enter password here..." onChange={(event) => onPasswordChange(event.target.value)}></input>

      <label htmlFor="passwordCheck">Confirm:</label>
      <input id="passwordCheck" placeholder="Confirm password here..." onChange={(event) => onPasswordCheckChange(event.target.value)}></input>

      <button disabled={email === "" || username === "" || password === "" || password !== passwordCheck} onClick={() => handleRegister()}>Register</button>

      <div className={`${showEmailConfirmation ? 'visible' : 'invisible'}`}>
        <h1>Registeration requires confirmation.</h1>
        <p>Check your inbox at {email} for a confirmation email.</p>
      </div>

      <div className={`${showError ? 'visible' : 'invisible'}`}>
        <h1>Oops, looks like you ran into an error when registering</h1>
        <p>{registerationError}</p>
      </div>

    </>
  );
}

export default RegisterPage;
