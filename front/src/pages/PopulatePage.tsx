import axios from "axios";
import { useEffect, useState } from "react";
import Content from "../components/Content";
import Footer from "../components/Footer";
import Header from "../components/Header";

interface Side {
  statement: string;
  sideID: string;
}

const PopulatePage = () => {
  const [statement, setStatement] = useState<string>("");
  const [sides, setSides] = useState<Side[]>([]);

  useEffect(() => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    if (backend_url) {
      axios.get(backend_url + "/sides")
        .then((response) => {
          setSides(response.data.sides)
        })
        .catch((error) => {
          console.log(error)
        })
    }
  }, [])


  const handleSideSubmit = async () => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    const formData = new FormData();
    formData.append("statement", statement);

    if (backend_url) {
      axios.post(backend_url + "/createSide", formData)
        .then((response) => {
          console.log(response.data)
        })
        .catch((err) => {
          console.log(err)
        })
    }
  }


  return (
    <>
      <Header />
      <Content>

        <h1>Add new statement</h1>
        <label htmlFor="statement">Statement:</label>
        <input id="statement" placeholder="Enter statement here..." onChange={(event) => setStatement(event.target.value)}></input>

        <button disabled={statement === ""} onClick={() => handleSideSubmit()}>Submit new statement</button>

        <table>
          <thead>
            <tr>
              <th>SideID</th>
              <th>Statement</th>
            </tr>
          </thead>
          <tbody>
            {sides && sides.map((item, index) => (
              <tr key={index}>
                <td>{item.sideID}</td>
                <td>{item.statement}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </Content>
      <Footer />

    </>
  );
}

export default PopulatePage;