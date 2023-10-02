import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import './login.css'

function SignUp() {
  return (
    <div>
      <Card className="w-[350px] rounded-lg shadow-2xl shadow-slate-700">
        <CardHeader className="shadow-md mb-4 flex flex-row justify-between">
          <CardTitle className="text-amber-500">Register</CardTitle>
          <CardTitle className="cursor-pointer text-xs text-slate-400 font-normal">
            <a href="mailto:yakcekoce@gmail.com">Need help ?</a>
          </CardTitle>
        </CardHeader>
        <CardContent>
          <form>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="name">Name & Surname</Label>
                <Input id="name" placeholder="Enter Name and Surname" />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="email">Email</Label>
                <Input id="email" placeholder="Email address" />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="password">Password</Label>
                <Input id="password" placeholder="Enter password" />
              </div>
            </div>
          </form>
        </CardContent>
        <CardFooter className="flex justify-between">
          <Button className="w-full bg-amber-600 hover:bg-amber-700">
            SIGN UP
          </Button>
        </CardFooter>
      </Card>
    </div>
  )
}

export default SignUp
