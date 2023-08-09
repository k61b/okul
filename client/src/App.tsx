import ButtonWithIcon from './components/Dashboard/ui/buttons/ButtonWithIcon '
import TextButton from './components/Dashboard/ui/buttons/TextButton'
import WrappedButton from './components/Dashboard/ui/buttons/WrappedButton'
import { MdOutlineFavoriteBorder } from 'react-icons/md'
import { BsInstagram, BsTwitter } from 'react-icons/bs'

function App() {
  return (
    <div className="flex flex-col justify-center items-center m-7">
      <WrappedButton label="Click Me" />

      <ButtonWithIcon
        label="Favorite"
        classes="mt-7 text-purple-800 hover:text-purple-900 hover:bg-transparent bg-transparent text-base"
        icon={
          <MdOutlineFavoriteBorder
            style={{ color: 'purple', fontSize: '22px', marginRight: '10px' }}
          />
        }
      />

      <ButtonWithIcon
        icon={<BsInstagram />}
        classes="mt-7 bg-transparent hover:bg-transparent text-slate-900 text-xl border	rounded-full hover:bg-slate-900 hover:text-white"
      />

      <ButtonWithIcon
        icon={<BsTwitter />}
        classes="mt-7 bg-transparent hover:bg-transparent text-slate-900 hover:bg-slate-900 hover:text-white text-xl border	rounded-full"
      />

      <TextButton label="Show On Map" />
    </div>
  )
}

export default App
