import { FC, ReactNode } from "react"
import Header from "./Header"
import Footer from "./Footer"

const Content: FC<{ children: ReactNode }> = ({ children }) => {
  return (
    <div className="content">
      <Header />
      <div className="flex-grow">
        {children}
      </div>
      <Footer />
    </div>
  )
}

export default Content