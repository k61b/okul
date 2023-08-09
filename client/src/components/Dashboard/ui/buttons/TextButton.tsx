import { Button } from '@/components/ui/button'

export default function TextButton(props: { label: string }) {
  return (
    <>
      <Button className='mt-7 bg-transparent text-orange-700 hover:bg-transparent text-base'>{props.label}</Button>
    </>
  )
}
