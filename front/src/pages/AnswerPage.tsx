import axios from "axios";
import { useEffect, useState } from "react";

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
    setQuestion(null)
    setPreferred("")
    setUnpreferred("")
    
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    if (backend_url) {
      axios.get(backend_url + "/question")
        .then((response) => {
          setQuestion(response.data)
        })
        .catch((error) => {
          console.log(error)
        })
    }
  }

  useEffect(() => {
    getNewQuestion();
  }, [])

  const handleChoice = (preferredId: string, unpreferredId: string) => {
    setPreferred(preferredId)
    setUnpreferred(unpreferredId)
  }


  const handleAnswerSubmit = async () => {
    const backend_url = import.meta.env.VITE_BACKEND_URL;
    const formData = new FormData();
    formData.append("username", "McGhee");
    formData.append("preferred", preferred);
    formData.append("unpreferred", unpreferred);

    if (backend_url) {
      axios.post(backend_url + "/createAnswer", formData)
        .then((response) => {
          console.log(response.data)
          getNewQuestion()
        })
        .catch((err) => {
          console.log(err)
        })
    }
  }


  return (
    <>
      <h1>Pick your preferred statement</h1>
      {question &&
        <div>
          <label htmlFor="left">Choice 1:</label>
          <button id="left" onClick={() => handleChoice(question.leftId, question.rightId)}>{question?.leftStatement}</button>

          <label htmlFor="right">Choice 2:</label>
          <button id="right" onClick={() => handleChoice(question.rightId, question.leftId)}>{question?.rightStatement}</button>

          <button disabled={preferred === "" || unpreferred === ""} onClick={() => handleAnswerSubmit()}>Submit answer</button>
        </div>
      }
    </>
  );
}

export default AnswerPage;