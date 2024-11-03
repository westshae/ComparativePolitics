import axios from "axios";
import { useEffect, useState } from "react";
import Content from "../components/Content";

interface Question {
  leftStatement: string;
  leftId: string;
  rightStatement: string;
  rightId: string;
}

const AnswerPage = () => {
  const [question, setQuestion] = useState<Question | null>(null);
  const [preferred, setPreferred] = useState<string>("");
  const [unpreferred, setUnpreferred] = useState<string>("");

  const getNewQuestion = async () => {
    setQuestion(null);
    setPreferred("");
    setUnpreferred("");

    const backend_url = import.meta.env.VITE_BACKEND_URL;
    if (backend_url) {
      axios.get(backend_url + "/question")
        .then((response) => {
          setQuestion(response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    }
  };

  useEffect(() => {
    getNewQuestion();
  }, []);

  const handleChoice = (preferredId: string, unpreferredId: string) => {
    setPreferred(preferredId);
    setUnpreferred(unpreferredId);
  };

  const handleAnswerSubmit = async () => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    const username = localStorage.getItem("name");
    if (!username) {
      return;
    }
    const formData = new FormData();
    formData.append("username", username);
    formData.append("preferred", preferred);
    formData.append("unpreferred", unpreferred);

    if (backend_url) {
      axios.post(backend_url + "/createAnswer", formData)
        .then((response) => {
          console.log(response.data);
          getNewQuestion();
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  return (
    <Content>
      <div className="page-container">
        <h1 className="page-title">Pick Your Preferred Statement</h1>
        {question && (
          <div className="choice-group">
            <label htmlFor="left" className="choice-label">Choice 1:</label>
            <button
              id="left"
              onClick={() => handleChoice(question.leftId, question.rightId)}
              className={`choice-button ${
                preferred === question.leftId ? "choice-button-selected" : ""
              }`}
            >
              {question.leftStatement}
            </button>

            <label htmlFor="right" className="choice-label">Choice 2:</label>
            <button
              id="right"
              onClick={() => handleChoice(question.rightId, question.leftId)}
              className={`choice-button ${
                preferred === question.rightId ? "choice-button-selected" : ""
              }`}
            >
              {question.rightStatement}
            </button>

            <button
              className="submit-button"
              disabled={preferred === "" || unpreferred === ""}
              onClick={handleAnswerSubmit}
            >
              Submit Answer
            </button>
          </div>
        )}
      </div>
    </Content>
  );
};

export default AnswerPage;
