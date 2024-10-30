const getJWTToken = () => {
  const token = localStorage.getItem("token");
  return token;
};


export {getJWTToken}