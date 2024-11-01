import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

interface Question {
  combiner: string;
  leftSideID: number;
  leftStatement: string;
  questionID: number;
  rightSideID: number;
  rightStatement: string;
}
interface Side {
  statement: string;
  sideID: string;
}

const PopulatePage = () => {
  const navigate = useNavigate();

  const [statement, setStatement] = useState<string>("");
  const [combiner, setCombiner] = useState<string>("");
  const [leftSideId, setLeftSideId] = useState<string>("");
  const [rightSideId, setRightSideId] = useState<string>("");

  const [questions, setQuestions] = useState<Question[]>([]);
  const [sides, setSides] = useState<Side[]>([]);

  useEffect(() => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    if (backend_url) {
      axios.get(backend_url + "/questions")
        .then((response) => {
          setQuestions(response.data.questions)
        })
        .catch((error) => {
          console.log(error)
        })
    }
  }, [])

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
      .then((response)=>{
        console.log(response.data)
      })
      .catch((err)=>{
        console.log(err)
      })
    }
  }

  const handleQuestionSubmit = async () => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    const formData = new FormData();
    formData.append("combiner", combiner);
    formData.append("leftSideId", leftSideId);
    formData.append("rightSideId", rightSideId);

    if (backend_url) {
      axios.post(backend_url + "/createQuestion", formData)
      .then((response)=>{
        console.log(response.data)
      })
      .catch((err)=>{
        console.log(err)
      })

    }
  }


  return (
    <>
      <h1>Add new statement</h1>
      <label htmlFor="statement">Statement:</label>
      <input id="statement" placeholder="Enter statement here..." onChange={(event) => setStatement(event.target.value)}></input>

      <button disabled={statement === ""} onClick={() => handleSideSubmit()}>Submit new statement</button>

      <h1>Add new question</h1>
      <label htmlFor="combiner">Combining Statement:</label>
      <input id="combiner" placeholder="Enter combining statement here..." onChange={(event) => setCombiner(event.target.value)}></input>
      <label htmlFor="leftSideId">LeftSideId:</label>
      <input id="leftSideId" placeholder="Enter left side id here..." onChange={(event) => setLeftSideId(event.target.value)}></input>
      <label htmlFor="rightSideId">RightSideID</label>
      <input id="rightSideId" placeholder="Enter right side id here..." onChange={(event) => setRightSideId(event.target.value)}></input>

      <button disabled={combiner === "" || leftSideId === "" || rightSideId === ""} onClick={() => handleQuestionSubmit()}>Submit new question</button>

      <table>
        <thead>
          <tr>
            <th>Combiner</th>
            <th>Left Side ID</th>
            <th>Left Statement</th>
            <th>Question ID</th>
            <th>Right Side ID</th>
            <th>Right Statement</th>
          </tr>
        </thead>
        <tbody>
          {questions.map((item, index) => (
            <tr key={index}>
              <td>{item.combiner}</td>
              <td>{item.leftSideID}</td>
              <td>{item.leftStatement}</td>
              <td>{item.questionID}</td>
              <td>{item.rightSideID}</td>
              <td>{item.rightStatement}</td>
            </tr>
          ))}
        </tbody>
      </table>

      <table>
        <thead>
          <tr>
            <th>SideID</th>
            <th>Statement</th>
          </tr>
        </thead>
        <tbody>
          {sides.map((item, index) => (
            <tr key={index}>
              <td>{item.sideID}</td>
              <td>{item.statement}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
}

export default PopulatePage;