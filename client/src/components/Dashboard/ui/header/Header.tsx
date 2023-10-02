import { AiOutlineHeart, AiOutlineUser } from 'react-icons/ai'
import { HiMenuAlt1 } from 'react-icons/hi'
import OnlyHead from '../../../../assets/images/only_head.png'
import ButtonWithIcon from '../buttons/ButtonWithIcon'
import './header.css'
import {
  Sheet,
  SheetContent,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet'
export default function Header() {
  return (
    <div className="w-full bg-white shadow-xl sticky top-0 left-0 right-0">
      <div className="flex justify-around items-center h-16">
        <div className="p-8 flex absolute w-full justify-start lg:hidden">
          <Sheet>
            <SheetTrigger>
              <HiMenuAlt1 style={{ fontSize: '25px' }} />
            </SheetTrigger>
            <SheetContent>
              <ul className="mt-10 text-black font-bold text-base border-b">
                <li className="cursor-pointer hover:underline hover:text-teal-700 mb-4">
                  Nurseries
                </li>
                <li className="cursor-pointer hover:underline hover:text-teal-700 mb-4">
                  Primary
                </li>
                <li className="cursor-pointer hover:underline hover:text-teal-700 mb-5">
                  Secondary
                </li>
              </ul>

              <ul className="mt-5 text-slate-600 font-semibold text-base space-x-4">
                <SheetTitle className="mb-5">For School</SheetTitle>
                <li className="cursor-pointer hover:underline hover:text-teal-700 mb-4">
                  <a href="mailto:someone@example.com">Contact Us</a>
                </li>
                <li className="cursor-pointer hover:underline hover:text-teal-700 mb-4">
                  About Us
                </li>
                <li className="cursor-pointer hover:underline hover:text-teal-700 mb-5">
                  Work with Us
                </li>
              </ul>
            </SheetContent>
          </Sheet>
        </div>

        <div className="flex items-center">
          <div className="flex flex-row justify-center items-center cursor-pointer">
            <img src={OnlyHead} alt="Minnaks" className="w-14" />
            <h1 className="header text-teal-900 font-bold text-3xl ml-2">
              Minnaks
            </h1>
          </div>
          <ul className="hidden lg:flex text-black text-base space-x-10 ml-10">
            <li className="cursor-pointer hover:underline hover:text-teal-700">
              Nurseries
            </li>
            <li className="cursor-pointer hover:underline hover:text-teal-700">
              Primary
            </li>
            <li className="cursor-pointer hover:underline hover:text-teal-700">
              Secondary
            </li>
          </ul>
        </div>
        <div className="hidden lg:flex items-center">
          <ButtonWithIcon
            label="Saved"
            classes="flex flex-col justify-center items-center font-normal text-black text-base hover:underline"
            icon={
              <AiOutlineHeart style={{ color: 'black', fontSize: '25px' }} />
            }
          />
          <ButtonWithIcon
            label="Login"
            classes="flex flex-col justify-center items-center font-normal text-black hover:underline text-base"
            icon={
              <AiOutlineUser style={{ color: 'black', fontSize: '25px' }} />
            }
          />
        </div>
      </div>
    </div>
  )
}
