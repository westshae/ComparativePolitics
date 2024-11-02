import { FC, ReactNode } from "react"

const Content:FC<{children: ReactNode}> = ({children}) => {
  return (
    <div>
      {children}
    </div>
  )
}

export default Content