import axios from "axios";
import { useEffect, useState } from "react";
import Content from "../components/Content";

interface Side {
  statement: string;
  sideID: string;
}

const PopulatePage = () => {
  const [statement, setStatement] = useState<string>("");
  const [sides, setSides] = useState<Side[]>([]);

  useEffect(() => {
    loadSides();
  }, []);

  const loadSides = async () => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    if (backend_url) {
      axios.get(backend_url + "/sides")
        .then((response) => {
          setSides(response.data.sides);
        })
        .catch((error) => {
          console.log(error);
        });
    }
  }

  const handleSideSubmit = async () => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    const formData = new FormData();
    formData.append("statement", statement);

    if (backend_url) {
      axios.post(backend_url + "/createSide", formData)
        .then((response) => {
          console.log(response.data);
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  return (
    <Content>
      <div className="page-container">
        <h1 className="page-title">Add New Statement</h1>

        <div className="form-group">
          <label htmlFor="statement" className="input-label">Statement:</label>
          <input
            id="statement"
            placeholder="Enter statement here..."
            className="input-field"
            onChange={(event) => setStatement(event.target.value)}
          />
        </div>

        <button
          className="submit-button"
          disabled={statement === ""}
          onClick={handleSideSubmit}
        >
          Submit New Statement
        </button>

        <button
          className="submit-button"
          onClick={loadSides}
        >
          Refresh
        </button>


        <div className="table-container">
          <table className="table">
            <thead className="table-head">
              <tr>
                <th className="table-cell">SideID</th>
                <th className="table-cell">Statement</th>
              </tr>
            </thead>
            <tbody>
              {sides && sides.map((item, index) => (
                <tr key={index} className="table-row">
                  <td className="table-cell">{item.sideID}</td>
                  <td className="table-cell">{item.statement}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </Content>
  );
};

export default PopulatePage;
