import { useEffect, useState } from 'react'
import './login.css'
import SignIn from './SignIn'
import SignUp from './SignUp'
import Images from './Images'
import ClimbKid from '../../../assets/images/climb_kid.png'

function Login() {
  const [radius, setRadius] = useState(0)
  const [showSignIn, setShowSignIn] = useState(true)
  const toggleView = () => {
    setShowSignIn(!showSignIn)

    const animationDuration = 500
    const targetRadius = 75

    const startTime = Date.now()

    const animationFrame = () => {
      const currentTime = Date.now()
      const progress = (currentTime - startTime) / animationDuration

      if (progress < 1) {
        setRadius(progress * targetRadius)
        requestAnimationFrame(animationFrame)
      } else {
        setRadius(targetRadius)
      }
    }

    requestAnimationFrame(animationFrame)
  }

  useEffect(() => {
    const animationDuration = 500
    const targetRadius = 75

    const startTime = Date.now()

    const animationFrame = () => {
      const currentTime = Date.now()
      const progress = (currentTime - startTime) / animationDuration

      if (progress < 1) {
        setRadius(progress * targetRadius)
        requestAnimationFrame(animationFrame)
      } else {
        setRadius(targetRadius)
      }
    }

    requestAnimationFrame(animationFrame)
  }, [])

  return (
    <div>
      <div
        className="bg w-full "
        style={{ borderBottomRightRadius: `${radius}%` }}
      >
        <div className="bg__h w-full h-full flex flex-col justify-center items-center">
          <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl text-white max-w-xs mb-12">
            Welcome to a fresh start!
          </h1>
          <div className="component__trs z-10">
            {showSignIn ? <SignIn /> : <SignUp />}
            {showSignIn ? (
              <div className="climb md:hidden flex flex-row justify-end w-full relative">
                <img
                  src={ClimbKid}
                  alt="Climb Kid"
                  className="w-48 h-auto absolute bottom-60 right-0"
                />
              </div>
            ) : (
              <div className="climb md:hidden flex flex-row justify-end w-full relative">
                <img
                  src={ClimbKid}
                  alt="Climb Kid"
                  className="w-48 h-auto absolute bottom-80 right-0"
                />
              </div>
            )}
          </div>

          <div
            className="flex flex-col justify-center items-center
        mt-12 z-10"
          >
            <span className="text-amber-700 scroll-m-20 text-xl font-semibold tracking-tight mb-2">
              {showSignIn
                ? "Don't have an account yet ?"
                : 'Do you already have an account?'}
            </span>
            <span
              className="text-amber-700 border-b border-amber-700 cursor-pointer scroll-m-20 text-xl font-semibold tracking-tight"
              onClick={toggleView}
            >
              {showSignIn ? 'Create an account' : 'Log in'}
            </span>
          </div>
          <Images />
        </div>
      </div>
    </div>
  )
}
export default Login
