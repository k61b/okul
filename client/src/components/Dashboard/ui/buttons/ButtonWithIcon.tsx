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
        <div>{props.icon}</div>
        <div>{props.label}</div>
      </Button>
    </>
  )
}
