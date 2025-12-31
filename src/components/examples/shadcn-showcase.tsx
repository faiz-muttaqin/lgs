import { useState } from 'react'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/calendar'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Checkbox } from '@/components/ui/checkbox'
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from '@/components/ui/collapsible'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from '@/components/ui/hover-card'
import { Input } from '@/components/ui/input'
import { InputOTP, InputOTPGroup, InputOTPSlot } from '@/components/ui/input-otp'
import { Label } from '@/components/ui/label'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group'
import { ScrollArea } from '@/components/ui/scroll-area'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Separator } from '@/components/ui/separator'
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@/components/ui/sheet'
import { Skeleton } from '@/components/ui/skeleton'
import { Slider } from '@/components/ui/slider'
import { Switch } from '@/components/ui/switch'
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { Textarea } from '@/components/ui/textarea'
import { Toggle } from '@/components/ui/toggle'
import { ToggleGroup, ToggleGroupItem } from '@/components/ui/toggle-group'
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip'
import {
  AlertCircle,
  Bell,
  Bold,
  ChevronDown,
  Italic,
  Plus,
  Settings,
  Terminal,
  Underline,
  User,
} from 'lucide-react'

export default function ShadcnShowcase() {
  const [date, setDate] = useState<Date | undefined>(new Date())
  const [sliderValue, setSliderValue] = useState([50])
  const [switchChecked, setSwitchChecked] = useState(false)

  return (
    <div className="container mx-auto space-y-8 p-6">
      <div className="space-y-2">
        <h1 className="text-4xl font-bold tracking-tight">Shadcn UI Components Showcase</h1>
        <p className="text-muted-foreground">
          A comprehensive collection of all available components
        </p>
      </div>

      <Separator />

      {/* Buttons Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Buttons</h2>
        <div className="flex flex-wrap gap-2">
          <Button>Default</Button>
          <Button variant="secondary">Secondary</Button>
          <Button variant="destructive">Destructive</Button>
          <Button variant="outline">Outline</Button>
          <Button variant="ghost">Ghost</Button>
          <Button variant="link">Link</Button>
          <Button size="sm">Small</Button>
          <Button size="lg">Large</Button>
          <Button size="icon">
            <Plus className="h-4 w-4" />
          </Button>
          <Button disabled>Disabled</Button>
        </div>
      </section>

      <Separator />

      {/* Alerts Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Alerts</h2>
        <div className="space-y-2">
          <Alert>
            <Terminal className="h-4 w-4" />
            <AlertTitle>Default Alert</AlertTitle>
            <AlertDescription>
              You can add components to your app using the CLI.
            </AlertDescription>
          </Alert>
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Error</AlertTitle>
            <AlertDescription>
              Your session has expired. Please log in again.
            </AlertDescription>
          </Alert>
        </div>
      </section>

      <Separator />

      {/* Badges Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Badges</h2>
        <div className="flex flex-wrap gap-2">
          <Badge>Default</Badge>
          <Badge variant="secondary">Secondary</Badge>
          <Badge variant="destructive">Destructive</Badge>
          <Badge variant="outline">Outline</Badge>
        </div>
      </section>

      <Separator />

      {/* Cards Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Cards</h2>
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <Card>
            <CardHeader>
              <CardTitle>Card Title</CardTitle>
              <CardDescription>Card Description</CardDescription>
            </CardHeader>
            <CardContent>
              <p>Card content goes here.</p>
            </CardContent>
            <CardFooter>
              <Button>Action</Button>
            </CardFooter>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Notifications</CardTitle>
              <CardDescription>You have 3 unread messages.</CardDescription>
            </CardHeader>
            <CardContent className="space-y-2">
              <div className="flex items-center gap-2">
                <Bell className="h-4 w-4" />
                <span className="text-sm">New message from John</span>
              </div>
            </CardContent>
          </Card>
        </div>
      </section>

      <Separator />

      {/* Accordion Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Accordion</h2>
        <Accordion type="single" collapsible className="w-full">
          <AccordionItem value="item-1">
            <AccordionTrigger>Is it accessible?</AccordionTrigger>
            <AccordionContent>
              Yes. It adheres to the WAI-ARIA design pattern.
            </AccordionContent>
          </AccordionItem>
          <AccordionItem value="item-2">
            <AccordionTrigger>Is it styled?</AccordionTrigger>
            <AccordionContent>
              Yes. It comes with default styles that matches the other components' aesthetic.
            </AccordionContent>
          </AccordionItem>
          <AccordionItem value="item-3">
            <AccordionTrigger>Is it animated?</AccordionTrigger>
            <AccordionContent>
              Yes. It's animated by default, but you can disable it if you prefer.
            </AccordionContent>
          </AccordionItem>
        </Accordion>
      </section>

      <Separator />

      {/* Avatar Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Avatars</h2>
        <div className="flex gap-2">
          <Avatar>
            <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
          <Avatar>
            <AvatarFallback>JD</AvatarFallback>
          </Avatar>
          <Avatar>
            <AvatarFallback>
              <User className="h-4 w-4" />
            </AvatarFallback>
          </Avatar>
        </div>
      </section>

      <Separator />

      {/* Form Inputs Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Form Inputs</h2>
        <div className="grid gap-4 md:grid-cols-2">
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" placeholder="Email" />
          </div>
          <div className="space-y-2">
            <Label htmlFor="password">Password</Label>
            <Input id="password" type="password" placeholder="Password" />
          </div>
          <div className="space-y-2">
            <Label htmlFor="message">Message</Label>
            <Textarea id="message" placeholder="Type your message here." />
          </div>
          <div className="space-y-2">
            <Label htmlFor="framework">Framework</Label>
            <Select>
              <SelectTrigger id="framework">
                <SelectValue placeholder="Select a framework" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="next">Next.js</SelectItem>
                <SelectItem value="react">React</SelectItem>
                <SelectItem value="vue">Vue</SelectItem>
                <SelectItem value="svelte">Svelte</SelectItem>
              </SelectContent>
            </Select>
          </div>
        </div>
      </section>

      <Separator />

      {/* Checkbox and Radio Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Checkbox & Radio</h2>
        <div className="grid gap-4 md:grid-cols-2">
          <div className="space-y-2">
            <Label>Checkbox</Label>
            <div className="flex items-center space-x-2">
              <Checkbox id="terms" />
              <label
                htmlFor="terms"
                className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              >
                Accept terms and conditions
              </label>
            </div>
          </div>
          <div className="space-y-2">
            <Label>Radio Group</Label>
            <RadioGroup defaultValue="option-one">
              <div className="flex items-center space-x-2">
                <RadioGroupItem value="option-one" id="option-one" />
                <Label htmlFor="option-one">Option One</Label>
              </div>
              <div className="flex items-center space-x-2">
                <RadioGroupItem value="option-two" id="option-two" />
                <Label htmlFor="option-two">Option Two</Label>
              </div>
            </RadioGroup>
          </div>
        </div>
      </section>

      <Separator />

      {/* Switch and Slider Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Switch & Slider</h2>
        <div className="grid gap-4 md:grid-cols-2">
          <div className="flex items-center space-x-2">
            <Switch
              id="airplane-mode"
              checked={switchChecked}
              onCheckedChange={setSwitchChecked}
            />
            <Label htmlFor="airplane-mode">Airplane Mode</Label>
          </div>
          <div className="space-y-2">
            <Label>Volume: {sliderValue[0]}%</Label>
            <Slider
              value={sliderValue}
              onValueChange={setSliderValue}
              max={100}
              step={1}
            />
          </div>
        </div>
      </section>

      <Separator />

      {/* Dialogs Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Dialogs & Modals</h2>
        <div className="flex flex-wrap gap-2">
          <Dialog>
            <DialogTrigger asChild>
              <Button variant="outline">Open Dialog</Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Are you absolutely sure?</DialogTitle>
                <DialogDescription>
                  This action cannot be undone. This will permanently delete your account
                  and remove your data from our servers.
                </DialogDescription>
              </DialogHeader>
              <DialogFooter>
                <Button>Confirm</Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>

          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button variant="outline">Open Alert Dialog</Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
              <AlertDialogHeader>
                <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
                <AlertDialogDescription>
                  This action cannot be undone. This will permanently delete your
                  account and remove your data from our servers.
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel>Cancel</AlertDialogCancel>
                <AlertDialogAction>Continue</AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>

          <Sheet>
            <SheetTrigger asChild>
              <Button variant="outline">Open Sheet</Button>
            </SheetTrigger>
            <SheetContent>
              <SheetHeader>
                <SheetTitle>Edit profile</SheetTitle>
                <SheetDescription>
                  Make changes to your profile here. Click save when you're done.
                </SheetDescription>
              </SheetHeader>
              <div className="grid gap-4 py-4">
                <div className="grid grid-cols-4 items-center gap-4">
                  <Label htmlFor="name" className="text-right">
                    Name
                  </Label>
                  <Input id="name" defaultValue="Pedro Duarte" className="col-span-3" />
                </div>
              </div>
            </SheetContent>
          </Sheet>
        </div>
      </section>

      <Separator />

      {/* Dropdown Menu Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Dropdown Menu</h2>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline">
              Open Menu <ChevronDown className="ml-2 h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Profile</DropdownMenuItem>
            <DropdownMenuItem>Billing</DropdownMenuItem>
            <DropdownMenuItem>Team</DropdownMenuItem>
            <DropdownMenuItem>Subscription</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </section>

      <Separator />

      {/* Popover & Hover Card Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Popover & Hover Card</h2>
        <div className="flex gap-4">
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline">Open Popover</Button>
            </PopoverTrigger>
            <PopoverContent>
              <div className="space-y-2">
                <h4 className="font-medium leading-none">Dimensions</h4>
                <p className="text-muted-foreground text-sm">
                  Set the dimensions for the layer.
                </p>
              </div>
            </PopoverContent>
          </Popover>

          <HoverCard>
            <HoverCardTrigger asChild>
              <Button variant="link">@nextjs</Button>
            </HoverCardTrigger>
            <HoverCardContent className="w-80">
              <div className="flex justify-between space-x-4">
                <Avatar>
                  <AvatarFallback>VC</AvatarFallback>
                </Avatar>
                <div className="space-y-1">
                  <h4 className="text-sm font-semibold">@nextjs</h4>
                  <p className="text-sm">
                    The React Framework – created and maintained by @vercel.
                  </p>
                </div>
              </div>
            </HoverCardContent>
          </HoverCard>
        </div>
      </section>

      <Separator />

      {/* Tabs Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Tabs</h2>
        <Tabs defaultValue="account" className="w-full">
          <TabsList>
            <TabsTrigger value="account">Account</TabsTrigger>
            <TabsTrigger value="password">Password</TabsTrigger>
            <TabsTrigger value="settings">Settings</TabsTrigger>
          </TabsList>
          <TabsContent value="account" className="space-y-2">
            <p className="text-sm">Make changes to your account here.</p>
            <Input placeholder="Username" />
          </TabsContent>
          <TabsContent value="password" className="space-y-2">
            <p className="text-sm">Change your password here.</p>
            <Input type="password" placeholder="New password" />
          </TabsContent>
          <TabsContent value="settings" className="space-y-2">
            <p className="text-sm">Manage your settings here.</p>
          </TabsContent>
        </Tabs>
      </section>

      <Separator />

      {/* Table Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Table</h2>
        <Table>
          <TableCaption>A list of your recent invoices.</TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[100px]">Invoice</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Method</TableHead>
              <TableHead className="text-right">Amount</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell className="font-medium">INV001</TableCell>
              <TableCell>Paid</TableCell>
              <TableCell>Credit Card</TableCell>
              <TableCell className="text-right">$250.00</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">INV002</TableCell>
              <TableCell>Pending</TableCell>
              <TableCell>PayPal</TableCell>
              <TableCell className="text-right">$150.00</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">INV003</TableCell>
              <TableCell>Unpaid</TableCell>
              <TableCell>Bank Transfer</TableCell>
              <TableCell className="text-right">$350.00</TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </section>

      <Separator />

      {/* Tooltip Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Tooltip</h2>
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button variant="outline">Hover me</Button>
            </TooltipTrigger>
            <TooltipContent>
              <p>Add to library</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </section>

      <Separator />

      {/* Toggle Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Toggle & Toggle Group</h2>
        <div className="space-y-2">
          <div className="flex gap-2">
            <Toggle aria-label="Toggle italic">
              <Italic className="h-4 w-4" />
            </Toggle>
            <Toggle aria-label="Toggle bold">
              <Bold className="h-4 w-4" />
            </Toggle>
            <Toggle aria-label="Toggle underline">
              <Underline className="h-4 w-4" />
            </Toggle>
          </div>
          <ToggleGroup type="single">
            <ToggleGroupItem value="bold" aria-label="Toggle bold">
              <Bold className="h-4 w-4" />
            </ToggleGroupItem>
            <ToggleGroupItem value="italic" aria-label="Toggle italic">
              <Italic className="h-4 w-4" />
            </ToggleGroupItem>
            <ToggleGroupItem value="underline" aria-label="Toggle underline">
              <Underline className="h-4 w-4" />
            </ToggleGroupItem>
          </ToggleGroup>
        </div>
      </section>

      <Separator />

      {/* Collapsible Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Collapsible</h2>
        <Collapsible>
          <CollapsibleTrigger asChild>
            <Button variant="ghost" className="flex w-full justify-between">
              <span>Can I use this in my project?</span>
              <ChevronDown className="h-4 w-4" />
            </Button>
          </CollapsibleTrigger>
          <CollapsibleContent className="space-y-2">
            <p className="text-sm">
              Yes. Free to use for personal and commercial projects. No attribution
              required.
            </p>
          </CollapsibleContent>
        </Collapsible>
      </section>

      <Separator />

      {/* Command Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Command</h2>
        <Command className="rounded-lg border shadow-md">
          <CommandInput placeholder="Type a command or search..." />
          <CommandList>
            <CommandEmpty>No results found.</CommandEmpty>
            <CommandGroup heading="Suggestions">
              <CommandItem>
                <Settings className="mr-2 h-4 w-4" />
                <span>Settings</span>
              </CommandItem>
              <CommandItem>
                <User className="mr-2 h-4 w-4" />
                <span>Profile</span>
              </CommandItem>
              <CommandItem>
                <Bell className="mr-2 h-4 w-4" />
                <span>Notifications</span>
              </CommandItem>
            </CommandGroup>
          </CommandList>
        </Command>
      </section>

      <Separator />

      {/* Calendar Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Calendar</h2>
        <div className="flex justify-center">
          <Calendar mode="single" selected={date} onSelect={setDate} className="rounded-md border" />
        </div>
      </section>

      <Separator />

      {/* Input OTP Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Input OTP</h2>
        <div className="space-y-2">
          <Label>One-Time Password</Label>
          <InputOTP maxLength={6}>
            <InputOTPGroup>
              <InputOTPSlot index={0} />
              <InputOTPSlot index={1} />
              <InputOTPSlot index={2} />
              <InputOTPSlot index={3} />
              <InputOTPSlot index={4} />
              <InputOTPSlot index={5} />
            </InputOTPGroup>
          </InputOTP>
        </div>
      </section>

      <Separator />

      {/* Skeleton Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Skeleton</h2>
        <div className="space-y-2">
          <Skeleton className="h-4 w-[250px]" />
          <Skeleton className="h-4 w-[200px]" />
          <Skeleton className="h-4 w-[150px]" />
        </div>
      </section>

      <Separator />

      {/* Scroll Area Section */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold">Scroll Area</h2>
        <ScrollArea className="h-[200px] w-full rounded-md border p-4">
          <div className="space-y-2">
            {Array.from({ length: 50 }).map((_, i) => (
              <div key={i} className="text-sm">
                Item {i + 1}
              </div>
            ))}
          </div>
        </ScrollArea>
      </section>
    </div>
  )
}
