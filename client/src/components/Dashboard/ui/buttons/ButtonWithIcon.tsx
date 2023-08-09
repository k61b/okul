import { Button } from '@/components/ui/button'
import { ReactNode } from 'react'

export default function ButtonWithIcon(props: {
  icon: ReactNode
  classes?: string
  label?: string
}) {
  return (
    <>
      <Button className={props.classes}>
        {props.icon} {props.label}
      </Button>
    </>
  )
}
