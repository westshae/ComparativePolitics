import Content from "../components/Content";
import Register from "../components/Register";
import Login from "../components/Login";

const AuthPage = () => {
  return (
    <>
      <Content>
        <div className="w-1/2 gap-4 mx-auto flex flex-row">
          <Register />
          <Login />
        </div>
      </Content>
    </>
  );
}

export default AuthPage;