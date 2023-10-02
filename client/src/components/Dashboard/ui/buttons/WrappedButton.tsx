import { Button } from '@/components/ui/button'

export default function WrappedButton(props: { label: string }) {
  return (
    <>
      <Button className="bg-slate-800 hover:bg-slate-900 px-8 rounded-lg mt-7">
        {props.label}
      </Button>
    </>
  )
}
